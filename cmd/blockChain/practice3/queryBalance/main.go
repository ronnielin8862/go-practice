package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ronnielin8862/go-api/pkg/blockChain"
)

func main() {
	// TODO: 用本地節點永遠是０餘額，但是ＣＭＤ卻可以正常查詢到。  原因待確認
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/b904b0ebc4cb4a6692b47ef9147cafd8")
	// client, err := ethclient.Dial("https://localhost:30303")
	if err != nil {
		fmt.Println(err)
	}

	account := common.HexToAddress("0x2910543af39aba0cd09dbb2d50200b3e800a63d2")

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("從wei轉成eth單位", blockChain.WeiToEther(balance))
}
