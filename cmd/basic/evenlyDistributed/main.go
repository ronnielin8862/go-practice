package main

import (
	"fmt"
	"math/rand"
)

/**
亂數平均分佈
*/

func main() {
	//傳入值
	value := 100

	//亂數分佈區間
	rangePercent := 0.2

	//需混幣次數
	count := 3

	mixCoin(float64(value), rangePercent, count)

}

func mixCoin(value float64, rangePercent float64, count int) {
	var list []float64
	var total float64
	avg := value / float64(count)
	min := avg * (1 - rangePercent)
	max := avg * (1 + rangePercent)

	for i := 0; i < count; i++ {
		if i == count-1 {
			list = append(list, value-total)
		} else {
			num := float64(rand.Intn(int(max)-int(min)) + int(min))
			list = append(list, num)
			total += num
		}
	}

	for _, v := range list {
		fmt.Println(v)
	}
}
