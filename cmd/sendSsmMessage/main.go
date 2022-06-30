package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func structToString(p interface{}) string {
	tpe := reflect.TypeOf(p)
	val := reflect.ValueOf(p)
	s := make([]string, 0)
	for i := 0; i < tpe.NumField(); i++ {
		ft := tpe.Field(i)
		ss := ft.Tag.Get("json") + "="
		n := val.FieldByName(ft.Name).Type().Name()
		fmt.Println(n)
		switch {
		case n == "int":
			ss += strconv.FormatInt(val.FieldByName(ft.Name).Int(), 10)
			s = append(s, ss)
		case n == "string":
			ss = val.FieldByName(ft.Name).String()
			s = append(s, ss)
		}
	}
	return strings.Join(s, "&")
}

type MessageParam struct {
	secretId   string
	businessId string
	version    string
	timestamp  int64
	nonce      string
	signature  string
	mobile     string
	templateId string
	params     string
	paramType  string
}

func main() {
	strings := []string{"secretId", "businessId", "version", "timestamp", "nonce", "mobile", "templateId", "params", "paramType"}
	sort.Strings(strings)
	fmt.Println(strings)

	var param MessageParam

	verifyCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	param.secretId = "69e4cf01d82e95c3c3140af2eaff6cfa"
	param.businessId = "ae5cc7a207b7494b9b3c537db1aed8d4"
	param.version = "v2"
	param.timestamp = time.Now().Unix() * 1000
	param.nonce = "d34c049e4f434d20b244cd70092a5a9a"
	param.mobile = "13164925430"
	param.templateId = "15564"
	param.params = "{\"code\":" + verifyCode + "}"
	param.paramType = "json"

	fmt.Println("param == ", param.timestamp)

	newStrings := "businessId" + param.businessId + "mobile" + param.mobile + "nonce" + param.nonce + "paramType" + param.paramType +
		"params" + param.params + "secretId" + param.secretId + "templateId" + param.templateId + "timestamp" +
		strconv.FormatInt(param.timestamp, 10) + "version" + param.version + "b2d39d3890587fed463c304b920cb23e"
	fmt.Println("newStrings = ", newStrings)

	has := md5.Sum([]byte(newStrings))
	afterMd5 := hex.EncodeToString(has[:])
	fmt.Println("EncodeToString", afterMd5)
	param.signature = afterMd5

	//s := structToString(param)

	params := url.Values{}
	params.Set("secretId", param.secretId)
	params.Set("businessId", param.businessId)
	params.Set("version", param.version)
	t := strconv.FormatInt(param.timestamp, 10)
	params.Set("timestamp", t)
	params.Set("nonce", param.nonce)
	params.Set("mobile", param.mobile)
	params.Set("templateId", param.templateId)
	params.Set("params", param.params)
	params.Set("paramType", param.paramType)
	params.Set("signature", param.signature)

	send(params.Encode())
}

func send(s string) {
	url := "https://sms.dun.163.com/v2/sendsms"
	method := "POST"

	fmt.Println("ggggggg = ", s)

	payload := strings.NewReader(s)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
