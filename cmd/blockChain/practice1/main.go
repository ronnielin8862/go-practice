// Deprecated: 測試失敗，沒屌用的學習文章
package main

import (
	"context"
	"fmt"

	createConn "github.com/ronnielin8862/go-practice/pkg/blockChain/practice1"
)

// Deprecated: 測試失敗，沒屌用的學習文章
func main() {
	client, err := createConn.Connect("http://localhost:8545")
	if err != nil {
		fmt.Println(err.Error())
	}

	blockNumber, err := client.GetBlockNumber(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("blockNumber = ", blockNumber)
}
