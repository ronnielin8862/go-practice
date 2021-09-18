package main

import (
	"bufio"
	"fmt"
	"os"

	bc "github.com/ronnielin8862/go-api/pkg/blockChain/practice2"
)

func Help() {
	fmt.Println("There are 3 operations:")
	fmt.Println("Type 1 for adding a new Block")
	fmt.Println("Type 2 for printing the Blockchain")
	fmt.Println("Type 3 for exiting")
}

func main() {
	fmt.Println("Welcome to our Blockchain project.")
	fmt.Println("Enter h for help")

	var operations string

	NewBlockchain := bc.CreateBlockchain() // 新增 第一個 初始區塊鏈

	for {
		fmt.Scanln(&operations)

		if operations == "h" { // 顯示使用方法
			fmt.Println("Printing the help")
			Help()
		} else if operations == "1" { //開始創建 後續區塊鏈
			fmt.Println("Entering your data:")
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine() // 讀一整行 input
			NewBlockchain.AddBlock(data)    // 利用 input 作為 data 來創建區塊鏈
		} else if operations == "2" {
			for _, block := range NewBlockchain.Blocks {
				fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Println()
			} // 查詢資料
		} else if operations == "3" {
			break
		} else {
			fmt.Println("Please Enter h, 1, 2, 3")
		}
	}

}
