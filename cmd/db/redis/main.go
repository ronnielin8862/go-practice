package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rs *redis.Client

func init() {
	rs = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	err := rs.Ping(context.Background()).Err()
	if err != nil {
		fmt.Println("connect to redis failed, ", err)
	}
}

func main() {
	cx := context.Background()
	result, err := rs.HGet(cx, "test", "111").Result()
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Println("re : ", result)
}
