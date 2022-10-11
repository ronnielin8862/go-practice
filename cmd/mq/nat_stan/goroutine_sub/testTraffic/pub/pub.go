package pub

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"io/ioutil"
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
		for i := 0; i < 20000; i++ {
			chanTwo(handler, i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 20000; i++ {
			chan3(handler, i)
		}
		wg.Done()
	}()

	wg.Wait()
}

func chanTwo(handler stan.AckHandler, i int) {
	lineup, err := ioutil.ReadFile("/Users/ronnie/Documents/work/DDU/ddu-document/LiveStream/賽事/足球陣容數據-test_file.json")
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}
	_, err = stanP.PublishAsync("test_2", lineup, handler)
	if err != nil {
		fmt.Printf("publish2 err : i = %v, err = %s ", i, err)
	} else {
		fmt.Println("publish2 success: ", i)
	}
	time.Sleep(time.Millisecond * 1)
}

func chan3(handler stan.AckHandler, i int) {

	lineup, err := ioutil.ReadFile("/Users/ronnie/Documents/work/DDU/ddu-document/LiveStream/賽事/足球陣容數據-test_file.json")
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}
	_, err = stanP.PublishAsync("test_3", lineup, handler)
	if err != nil {
		fmt.Printf("publish3 err : i = %v, err = %s ", i, err)
	} else {
		fmt.Println("publish3 success: ", i)
	}
	time.Sleep(time.Millisecond * 1)

}
