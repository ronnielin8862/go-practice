package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/ronnielin8862/go-practice/pkg/utils"
	"github.com/spf13/cast"
	"log"
	"time"
)

var Nats stan.Conn

func init() {
	url := fmt.Sprintf("nats://127.0.0.1:4223")
	nc, _ := nats.Connect(
		url,
		nats.UserInfo("nats%3admin##1", "oscars3higehaohaizi"),
		nats.Timeout(time.Second*10),
		nats.PingInterval(time.Second*4),
	)
	ns, err := stan.Connect("test-cluster", "natsClicent107", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: ", reason)
		}))
	if err != nil {
		log.Fatal("nats connect fail")
	}
	Nats = ns
}

func main() {
	tryPush(Nats)
}

func tryPush(nats stan.Conn) {
	utils.TimeCost()()
	n := 3000000
	a := make(chan string, n)
	numChan := make(chan int, 10)
	sub := "tryLossMsg"
	for i := 0; i < n; i++ {
		a <- cast.ToString(i)
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case g := <-a:

					_, err := nats.PublishAsync(sub, []byte(g), nil)
					if err != nil {
						log.Println("push err", err)
					}
					time.Sleep(time.Millisecond * 1)
					fmt.Println("push ", g)
				default:
					numChan <- 1
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		<-numChan
		fmt.Printf("結束第%v個推送者\n", i)
	}

}
