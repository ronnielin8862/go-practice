package main

import (
	"github.com/ronnielin8862/go-api/pkg/blockChain/practice3/btc/genPrivateKey"
)

func main() {
	wallet := genPrivateKey.NewWallet()

	//fmt.Println(genPrivateKey.ByteString(wallet.PrivateKey))
	//fmt.Println("raw public key", genPrivateKey.ByteString(wallet.PublicKey))
	wallet.GetAddress()
}