package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.Keccak256()
	fmt.Println("hash = ", hash)
	// 	hash.append(transferFnSignature)
	// 	methodID := hash.Sum(nil)[:4]
	// 	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
}
