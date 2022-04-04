package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {

	phone := "965166374"
	area := "+855"

	allMobile := area + phone

	code := "12345"

	content := "【牛逼2】Your SMS verification code is " + code

	url := "http://api.wftqm.com/api/sms/mtsend"

	data := make(map[string]string)

	data["appkey"] = "tJjr6nTf"
	data["secretkey"] = "FGsavoh0"
	data["phone"] = allMobile
	data["content"] = content

	client := resty.New()
	resp, err := client.R().EnableTrace().
		SetBody(data).
		Post(url)
	if err != nil {
		fmt.Println("ERROR = ", err)
	}
	fmt.Println(resp.Status(), resp.Body(), resp.RawBody())
}
