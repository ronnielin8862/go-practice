package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	url := "https://api.sportradar.com/soccer-extended/trial/v4/stream/events/subscribe?api_key="

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Headers:", resp.Header)
	reader := bufio.NewReader(resp.Body)
	line, isPrefix, err := reader.ReadLine()
	for err == nil {
		fmt.Print(string(line))
		if !isPrefix {
			fmt.Println()
		}
		line, isPrefix, err = reader.ReadLine()
	}
	/**
	  dyq227dyvfct38emkur65tez
	*/
}
