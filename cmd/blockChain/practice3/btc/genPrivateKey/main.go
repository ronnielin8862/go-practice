package main

import (
	"github.com/ronnielin8862/go-practice/pkg/blockChain/practice3/btc/genPrivateKey"
)

// Deprecated: 採用 genPrivateKey2
func main() {
	wallet := genPrivateKey.NewWallet()

	wallet.GetAddress()
}
