package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/ronnielin8862/go-practice/pkg/utils"
	"time"
)

var stanS stan.Conn

func init() {
	var err error
	url := fmt.Sprintf("nats://127.0.0.1:4223")
	nc, _ := nats.Connect(
		url,
		nats.UserInfo("nats%3admin##1", "oscars3higehaohaizi"),
		nats.Timeout(time.Second*10),
		nats.PingInterval(time.Second*4),
	)
	stanS, err = stan.Connect("test-cluster", "natsClicent107", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: ", reason)
		}))
	if err != nil {
		fmt.Println("error by nats connect: ", err)
	}
}

func main() {

	//subscribe1, err := stanS.Subscribe("test_1", printSub1, stan.DurableName("test_1"), stan.AckWait(3*time.Second), stan.SetManualAckMode())
	//if err != nil {
	//	fmt.Println("error by subscribe: ", err)
	//}
	//defer subscribe1.Close()

	// 以下測試出 兩個不同頻道的訂閱，是同線程在處理。 約2ms 處理完 各1000 mq, 當採用兩
	defer utils.TimeCost()()
	//go func() {
	subscribe2, err := stanS.Subscribe("test_2", printSub2, stan.DurableName("test_2"), stan.AckWait(111*time.Second)) //, stan.SetManualAckMode()
	if err != nil {
		fmt.Println("error by subscribe: ", err)
	}
	defer subscribe2.Close()
	//	select {}
	//}()
	//
	//go func() {
	subscribe3, err := stanS.Subscribe("test_3", printSub3, stan.DurableName("test_3"), stan.AckWait(111*time.Second)) //, stan.SetManualAckMode()
	if err != nil {
		fmt.Println("error by subscribe: ", err)
	}

	defer subscribe3.Close()
	//	select {}
	//}()
	select {}
}

// 單訂閱為單線程，可透過goroutine來達到多線程
func printSub1(msg *stan.Msg) {
	go func() {
		fmt.Println("Sub1 Received a message: ", string(msg.Data))
		time.Sleep(3 * time.Second)
		err := msg.Ack()
		if err != nil {
			fmt.Println("error by ack: ", err)
		}
	}()
}

// 2, 3 兩個驗證了不同訂閱是相同的線程，單單增加一個線程處理速度就快了非常多
func printSub2(msg *stan.Msg) {
	go func() {
		fmt.Println("Sub2 Received a message: ", string(msg.Data))
		//err := msg.Ack()
		//if err != nil {
		//	fmt.Println("error by ack: ", err)
		//}
		time.Sleep(1 * time.Millisecond)
		//time.Sleep(3 * time.Second)
	}()
}

func printSub3(msg *stan.Msg) {
	go func() {
		fmt.Println("Sub3 Received a message: ", string(msg.Data))
		//err := msg.Ack()
		//if err != nil {
		//	fmt.Println("error by ack: ", err)
		//}
		time.Sleep(1 * time.Millisecond)
		//time.Sleep(3 * time.Second)
	}()
}
