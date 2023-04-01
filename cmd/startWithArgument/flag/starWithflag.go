package main

import "flag"

var (
	name, password string
)

// 除了cobra以外的另一種啟動帶入參數的方式
func main() {
	flag.StringVar(&name, "name", "default name", "log in user name")
	flag.StringVar(&password, "password", "default password", "log in user password")

	flag.Parse()

	println(name, password)

}
