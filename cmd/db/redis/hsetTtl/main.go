package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rds *redis.Client

func InitRedis() *redis.Client {
	Host, Port, Password := "localhost", "6379", ""
	rds = redis.NewClient(&redis.Options{
		Addr:     Host + ":" + Port,
		Password: Password,
		DB:       0,
	})
	return rds
}

func getTodaySurplusSecond1() (int64, error) {
	layout := "2006-01-02"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(layout, time.Now().Format(layout), loc)

	if err != nil {
		return 0, err
	}
	return 86400 - (time.Now().Unix() - t.Unix()), nil
}

func getTodaySurplusSecond() time.Time {
	layout := "2006-01-02"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(layout, time.Now().Format(layout), loc)
	t = t.AddDate(0, 0, 1)
	return t
}

func main() {
	InitRedis()

	getTodaySurplusSecond()

	fmt.Println("  ccc = ", rds.ClientID())

	result, err := rds.LRange("cc", 0, -1).Result()
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println("result =", result)
	}

	result2, err := rds.LPush("cc", "b").Result()
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println("2", result2)
	}
	rds.Close()
	//result, err = rds.LPush("c", "d").Result()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(result)
	//}

	//rds.Set("AAA", "BBB", time.LoadLocation("Asia/Taipei"))
}

//func hSetTtl(key, field, value string) {
//	result, err := rds.HSet(key, field, value).Result()
//	rds.TTL()
//	if err != nil {
//		fmt.Println("err = ", err)
//	} else {
//		fmt.Println("re = ", result)
//	}
//}
