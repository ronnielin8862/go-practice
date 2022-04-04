package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	SendRequest("POST", "http://localhost:5795/upload", bytes.NewReader([]byte("test")))
}

func SendRequest(method, url string, body io.Reader) error {
	httpReq, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	if method == "POST" {
		httpReq.Header.Add("Content-Type", "application/json")
	}
	httpReq.Close = true
	respData, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	respBytes, err := ioutil.ReadAll(respData.Body)
	if err != nil {
		return err
	}
	defer respData.Body.Close()
	fmt.Println(respBytes)
	return nil
}
