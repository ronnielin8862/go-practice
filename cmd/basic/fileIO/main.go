package main

import (
	"encoding/json"
	"fmt"
	"github.com/ronnielin8862/go-practice/globle"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("/Users/ronnie/Downloads/Telegram Desktop/篮球阵容数据 (1).json")
	if err != nil {
		fmt.Println("io err")
	}
	var sb []globle.BasketballPlayers
	err = json.Unmarshal(file, &sb)
	if err != nil {
		fmt.Println("json err")
	} else {
		fmt.Println("json : ", sb)
	}

}
