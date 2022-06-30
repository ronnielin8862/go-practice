package main

import (
	"time"
)

func main() {
	var param MessageParam

	param.secretId = "69e4cf01d82e95c3c3140af2eaff6cfa"
	param.businessId = "053e9eecb83d438b9a9a2774f70f7514"
	param.version = "v2"
	param.timestamp = time.Now().Unix()
	param.nonce = "5487"
	param.mobile = "13164925430"
	param.templateId = "15564"
	param.params = "牛總  收到訊息  請告知 Ronnie Tks"
	param.paramType = "json"

	//neCaptchaVerifier, _ := sdk.New(param.businessId, param.secretId, "b2d39d3890587fed463c304b920cb23e")
	//neCaptchaVerifier.Verify()
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
