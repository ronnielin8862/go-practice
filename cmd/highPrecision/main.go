package main

import (
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
)

func main() {

}

func testFloat(s string) string {
	f, _ := strconv.ParseFloat(s, 64)
	s2 := strconv.FormatFloat(f, 'E', -1, 64) //float64
	return s2
}

func testNewRat(r string) string {
	rat := new(big.Rat)
	rat.SetString(r)

	return rat.RatString()
}

func testShopSpringDecimal(s string) string {
	v, _ := decimal.NewFromString(s)
	return v.String()
}
