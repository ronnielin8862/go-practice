package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/spf13/cast"
	"sync"
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
		stan.PubAckWait(90*time.Second),
		stan.MaxPubAcksInflight(200000))
	if err != nil {
		fmt.Println("error by nats connect: ", err)
	}
}

func main() {
	handler := func(ackUid string, err error) {
		if err != nil {
			fmt.Println("error by publish: ", err)
		}
	}
	// 測試單訂閱 多線程的效果
	//chanOne(handler)

	// 測試多訂閱 多線程的效果

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {

		chanTwo(handler, &wg)
	}()

	go func() {
		chan3(handler, &wg)
	}()

	wg.Wait()
}

func chanOne(handler stan.AckHandler) {
	for i := 0; i < 3; i++ {
		async, err := stanP.PublishAsync("test_1", []byte(cast.ToString(i)), handler)
		if err != nil {
			fmt.Printf("publish1 err : i = %v, err = %s ", i, err)
		} else {
			fmt.Println("publish1 success: ", async)
		}
		time.Sleep(time.Millisecond * 1)
	}
}

func chanTwo(handler stan.AckHandler, wg *sync.WaitGroup) {
	for i := 0; i < 20000; i++ {
		_, err := stanP.PublishAsync("test_2", []byte(cast.ToString(i)), handler)
		if err != nil {
			fmt.Printf("publish2 err : i = %v, err = %s ", i, err)
		} else {
			fmt.Println("publish2 success: ", i)
		}
		time.Sleep(time.Millisecond * 1)
	}
	wg.Done()
}

func chan3(handler stan.AckHandler, wg *sync.WaitGroup) {
	for i := 0; i < 20000; i++ {
		_, err := stanP.PublishAsync("test_3", []byte(cast.ToString(i)), handler)
		if err != nil {
			fmt.Printf("publish3 err : i = %v, err = %s ", i, err)
		} else {
			fmt.Println("publish3 success: ", i)
		}
		time.Sleep(time.Millisecond * 1)
	}
	wg.Done()
}
