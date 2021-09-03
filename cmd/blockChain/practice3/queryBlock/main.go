package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/b904b0ebc4cb4a6692b47ef9147cafd8")
	if err != nil {
		fmt.Println(err)
	}

	//獲取最先區塊 header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v \n", header.Number.String())
	fmt.Printf("%+v \n", header.Hash().String())
	fmt.Printf("%+v \n", header.TxHash.String())

	//也是獲得最新區塊號
	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v \n", blockNum)

	//取得最新 －５的區塊相關資訊
	block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNum-5))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block.Number", block.Number())
	fmt.Println("block.Body", block.Body())
	fmt.Println("block.Hash", block.Hash())
	fmt.Println("block.Transactions().Len()", block.Transactions().Len())
	fmt.Println("block.Transactions", block.Transactions()) //效果同 block.Body

	tran, err := client.TransactionInBlock(context.Background(), block.Hash(), 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tran.Data", tran.Data())
	fmt.Println("tran.Hash", tran.Hash())

}
