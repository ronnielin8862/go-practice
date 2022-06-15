package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Create server connection
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	defer nc.Close()

	mcbAny := func(msg *nats.Msg) {
		log.Println("Any:", string(msg.Data))

	}
	mcbIthome := func(msg *nats.Msg) {
		log.Println("Ithome:", string(msg.Data))

	}
	var Sub1Cb *nats.Subscription
	var Sub2Cb *nats.Subscription
	var err error
	go func() {
		Sub1Cb, err = nc.Subscribe("testTopic.*", mcbAny)
		Sub1Cb, err = nc.Subscribe("test1", mcbAny)
		if err != nil {
			log.Println("queue subscribe testTopic.*:", err)
		}
	}()

	go func() {
		Sub2Cb, err = nc.Subscribe("*", mcbIthome)
		if err != nil {
			log.Println("queue subscribe testTopic.ithome:", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	Sub1Cb.Unsubscribe()
	Sub2Cb.Unsubscribe()
}
