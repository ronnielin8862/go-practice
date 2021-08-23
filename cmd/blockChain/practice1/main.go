package main

import (
	"context"
	"fmt"

	"github.com/ronnielin8862/go-api/pkg/blockChain/practice1"
)

func main() {
	client, err := practice1.Connect("http://localhost:8545")
	if err != nil {
		fmt.Println(err.Error())
	}

	blockNumber, err := client.GetBlockNumber(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("blockNumber = ", blockNumber)
}
