package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
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
	//key := "test996"
	//value := "可憐哪"
	//tryExpirationMessage(key, value)

	i, err := rs.HGetAll(context.Background(), "testInt64").Result()

	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("i = ", i)
	}
}

func tryExpirationMessage(key string, value string) {
	err := rs.LPush(context.Background(), key, value).Err()
	if err != nil {
		fmt.Println("err : ", err)
	}
	r, err := rs.Expire(context.Background(), key, 3*time.Second).Result()
	fmt.Println("r = ", r, ", err = ", err)

	r2, err := rs.Exists(context.Background(), key).Result()
	fmt.Println("r2 = ", r2, ", err = ", err)

	time.Sleep(3 * time.Second)

	r3, err := rs.Exists(context.Background(), key).Result()
	fmt.Println("r3 = ", r3, ", err = ", err)
}
