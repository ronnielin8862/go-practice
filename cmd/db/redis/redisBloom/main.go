package main

import (
	"fmt"
	redisBloom "github.com/RedisBloom/redisbloom-go"
	"github.com/gomodule/redigo/redis"
	//"github.com/gomodule/redigo/redis"
)

func init() {
	host := "localhost:6379"
	password := ""
	pool = &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", host, redis.DialPassword(password))
	}}

}

var pool *redis.Pool

func main() {
	// Connect to localhost with no password
	client := redisBloom.NewClientFromPool(pool, "bloom-client-1")
	// BF.ADD mytest item
	_, err := client.Add("mytest", "myItem")
	if err != nil {
		fmt.Println("Error:", err)
	}

	exists, err := client.Exists("mytest", "myItem")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("myItem exists in mytest: ", exists)
}
