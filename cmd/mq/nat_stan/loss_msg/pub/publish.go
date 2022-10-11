package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/spf13/cast"
	"time"
)

var stanP stan.Conn

func init() {
	var err error
	url := fmt.Sprintf("nats://127.0.0.1:4223")
	nc, _ := nats.Connect(
		url,
		nats.UserInfo("nats%3admin##1", "oscars3higehaohaizi"),
		nats.Timeout(time.Second*10),
		nats.PingInterval(time.Second*4),
	)
	stanP, err = stan.Connect("test-cluster", "natsClicent108", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: ", reason)
		}),
		stan.PubAckWait(30*time.Second),
		stan.MaxPubAcksInflight(100))

	if err != nil {
		fmt.Println("error by nats connect: ", err)
	}
}

// 這裡實驗結果呈現，如果執行緒結束前沒有等待一毫秒左右，異步發送的消息會丟失。
func main() {
	//var handler stan.AckHandler
	handler := func(ackUid string, err error) {
		if err != nil {
			fmt.Println("error by publish: ", err)
		}
	}

	//chanOne(handler)

	chanTwo(handler)

}

func chanOne(handler stan.AckHandler) {
	for i := 0; i < 3; i++ {
		//async, err := stanP.Publish("test_1", []byte(cast.ToString(i)), handler)
		//if err != nil {
		//	fmt.Printf("publish1 err : i = %v, err = %s ", i, err)
		//} else {
		//	fmt.Println("publish1 success: ", async)
		//}

		err := stanP.Publish("test_1", []byte(cast.ToString(i)))
		if err != nil {
			fmt.Printf("publish1 err : i = %v, err = %s ", i, err)
		} else {
			fmt.Println("publish1 success: ", i)
		}
	}
}

func chanTwo(handler stan.AckHandler) {
	//for i := 0; i < 3; i++ {
	//	async, err := stanP.PublishAsync("test_2", []byte(cast.ToString(i)), handler)
	//	if err != nil {
	//		fmt.Printf("publish2 err : i = %v, err = %s ", i, err)
	//	} else {
	//		fmt.Println("publish2 success: ", async)
	//	}
	//	//time.Sleep(time.Millisecond * 1)
	//}
	defer fmt.Println("end")
	for i := 0; i < 10; i++ {
		_, err := stanP.PublishAsync("test_2", []byte(cast.ToString(i)), nil)
		if err != nil {
			fmt.Printf("publish2 err : i = %v, err = %s ", i, err)
		}
	}
	//time.Sleep(1 * time.Millisecond)
}
