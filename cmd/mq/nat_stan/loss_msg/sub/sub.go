package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
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
	_, err := stanS.Subscribe("test_1", printSub1)
	if err != nil {
		fmt.Println("error by subscribe: ", err)
	}
	//defer subscribe1.Close()

	_, err = stanS.Subscribe("test_2", printSub2, stan.DurableName("test_2"), stan.AckWait(3*time.Second), stan.SetManualAckMode())
	if err != nil {
		fmt.Println("error by subscribe: ", err)
	}
	//defer subscribe12.Close()

	select {}
}

func printSub1(msg *stan.Msg) {
	fmt.Println("Sub1 Received a message: ", string(msg.Data))
	//msg.Ack()
}

func printSub2(msg *stan.Msg) {
	fmt.Println("Sub2 Received a message: ", string(msg.Data))
	err := msg.Ack()
	if err != nil {
		fmt.Println("error by ack: ", err)
	}
	//msg.Ack()
}
