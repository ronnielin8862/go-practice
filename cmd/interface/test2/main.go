package main

import "fmt"

type Database interface {
	Read(string) string
	Write(string)
}

// MySQL 會實作 Database 介面。
type MySQL struct{}

func (m MySQL) Read(s string) string {
	fmt.Println("Mysql的 ", s)
	return s
}
func (m MySQL) Write(string) {}

// MongoDB 會實作 Database 介面。
type MongoDB struct{}

func (m MongoDB) Read(s string) string {
	fmt.Println("Mongo的 ", s)
	return s
}
func (m MongoDB) Write(string) {}

// NewDb 會將接收到的物件以 Database 實作，並且呼叫相關函式對資料庫進行操作。
func NewDb(db Database, s string) {
	// 讀取資料庫。
	db.Read(s)
	// 寫入資料庫。
	db.Write(s)
}

func main() {

	s := "我叫你吹"
	// 將建構體傳入 NewDb 就會被實作成 Database。
	NewDb(MySQL{}, s)
	NewDb(MongoDB{}, s)

}
