package main

import (
	"bytes"
	"net/http"
)

func main() {
	CallbackAfterNewMemberJoin()
}

func CallbackAfterNewMemberJoin() {
	url := "http://127.0.0.1:1988/webapi/tencentIm/callback?SdkAppid=1400722522&CallbackCommand=Group.CallbackAfterNewMemberJoin&contenttype=json&ClientIP=1.2.3.4&OptPlatform=IOS"
	body := "{\n    \"CallbackCommand\": \"Group.CallbackAfterNewMemberJoin\", \n    \"GroupId\" : \"@TGS#2J4SZEAEL\",\n    \"Type\": \"Public\",\n    \"JoinType\": \"Apply\",\n    \"Operator_Account\": \"leckie\", \n    \"NewMemberList\": [ \n        {\n            \"Member_Account\": \"jared\"\n        },\n        {\n            \"Member_Account\": \"tommy\"\n        }\n    ]\n}"

	var jsonStr = []byte(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
