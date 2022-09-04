package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/spf13/cast"
	"log"
	"sync/atomic"
	"time"
)

var natConn *nats.Conn

func init() {
	url := fmt.Sprintf("nats://127.0.0.1:4223")
	nc, _ := nats.Connect(
		url,
		//nats.UserInfo("nats%3admin##1", "oscars3higehaohaizi"),
		nats.Timeout(time.Second*10),
		nats.PingInterval(time.Second*4),
	)
	natConn = nc
}

func main() {
	trySub(natConn)
}

var count int = 0
var at atomic.Int32

var subs []*stan.Conn

func trySub(nats *nats.Conn) {
	for i := 0; i < 10; i++ {
		fmt.Println("sub  i = ", i)
		stanConn, err := stan.Connect("test-cluster", "natsClicent1"+cast.ToString(i), stan.NatsConn(nats),
			stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
				fmt.Println("Connection lost, reason: ", reason)
			}))
		if err != nil {
			log.Fatal("nats connect fail:", err)
		}
		subs = append(subs, &stanConn)
		Counting(stanConn)
	}
	select {}
}

func Counting(stanConn stan.Conn) {
	handler := func(msg *stan.Msg) {
		count++
		at.Add(1)
		fmt.Println("count= ", count, ", at.Load()= ", at.Load())
	}

	_, err := stanConn.QueueSubscribe("tryLossMsg", "queue ", handler, stan.DurableName("tryLossMsg"))
	if err != nil {
		fmt.Println("sub err", err)
	}
}
