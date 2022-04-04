package sdk

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/galaxy-book/captcha-golang-demo/http"
	"github.com/galaxy-book/captcha-golang-demo/json"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const (
	verifyApi = "http://c.dun.163yun.com/api/v2/verify"
	version   = "v2"
)

var (
	CaptchaIdIsEmpty    = errors.New("captchaId is empty")
	SecretIdIsEmpty     = errors.New("secretId is empty")
	SecretKeyIsEmpty    = errors.New("secretKey is empty")
	ValidateDataIsEmpty = errors.New("validate data is empty")
)

type neCaptchaVerifier struct {
	captchaId  string        `json:"captchaId"`
	secretPair *neSecretPair `json:"secretPair"`
}

type neSecretPair struct {
	secretId  string `json:"secretId"`
	secretKey string `json:"secretKey"`
}

type VerifyResult struct {
	Err       int    `json:"error"`
	Msg       string `json:"msg"`
	Result    bool   `json:"result"`
	Phone     string `json:"phone"`
	ExtraData string `json:"extraData"`
}

func New(captchaId, secretId, secretKey string) (*neCaptchaVerifier, error) {
	if captchaId == "" {
		return nil, CaptchaIdIsEmpty
	}
	if secretId == "" {
		return nil, SecretIdIsEmpty
	}
	if secretKey == "" {
		return nil, SecretKeyIsEmpty
	}
	return &neCaptchaVerifier{
		captchaId: captchaId,
		secretPair: &neSecretPair{
			secretId:  secretId,
			secretKey: secretKey,
		},
	}, nil
}

func (n *neCaptchaVerifier) Verify(validate, user string) (*VerifyResult, error) {
	if validate == "" || user == "" {
		return nil, ValidateDataIsEmpty
	}
	params := map[string]string{}
	params["captchaId"] = n.captchaId
	params["validate"] = validate
	params["user"] = user
	params["secretId"] = n.secretPair.secretId
	params["version"] = version
	params["timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	params["nonce"] = random(20)
	params["signature"] = genSignature(n.secretPair.secretKey, params)

	data, err := http.PostForm(verifyApi, params)
	if err != nil {
		return nil, err
	}
	fmt.Println(verifyApi, "request data:", params)
	fmt.Println(verifyApi, "response data:", data)
	verifyResult := &VerifyResult{}
	json.FromJsonIgnoreError(data, verifyResult)
	return verifyResult, nil
}

func genSignature(secretKey string, params map[string]string) string {
	var keys []string
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	buf := bytes.NewBufferString("")
	for _, key := range keys {
		buf.WriteString(key + params[key])
	}
	buf.WriteString(secretKey)
	has := md5.Sum(buf.Bytes())
	return fmt.Sprintf("%x", has)
}

func random(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
