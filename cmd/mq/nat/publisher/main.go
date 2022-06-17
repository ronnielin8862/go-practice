package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"strconv"
	"time"
)

type ChatHistory struct {
	Id         int64  `json:"id" gorm:"primaryKey;autoIncrement"` // bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	Uid        int    `json:"uid"`                                //int(11) DEFAULT NULL COMMENT '用户id',
	RoomId     int    `json:"room_id"`                            //varchar(100) DEFAULT NULL COMMENT '房间id',
	CreateTime int64  `json:"create_time"`                        //datetime DEFAULT NULL,
	Content    string `json:"content"`
}

type SendGiftReq struct {
	Anchorid int   `json:"anchorid" validate:"required"`
	Giftid   int   `json:"giftid" validate:"required"`
	Liveid   int64 `json:"liveid"`
	Count    int   `json:"count"`
	Uid      int
}

func main() {
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	jetStreamPubTestForDDU(nc)
}

func generalPublish(nc *nats.Conn) {

	// 发布-订阅 模式，向 test1 发布一个 `Hello World` 数据
	chat := ChatHistory{
		Id:         666,
		Uid:        555,
		RoomId:     444,
		CreateTime: 1655273431,
		Content:    "鬼塚英吉",
	}

	chatMarshal, err := json.Marshal(chat)
	if err != nil {
		return
	}
	_ = nc.Publish("chatRecordChannel", []byte(chatMarshal))

	gift := SendGiftReq{
		Anchorid: 666,
		Giftid:   555,
		Liveid:   444,
		Count:    666,
		Uid:      777,
	}
	giftMarshal, err := json.Marshal(gift)
	if err != nil {
		return
	}
	_ = nc.Publish("sendGiftChannel", []byte(giftMarshal))

	// 队列 模式，发布是一样的，只是订阅不同，向 test2 发布一个 `Hello zngw` 数据
	//_ = nc.Publish("test.2", []byte("鬼塚英吉"))

	// 请求-响应， 向 test3 发布一个 `help me` 请求数据，设置超时间3秒，如果有多个响应，只接收第一个收到的消息
	//msg, err := nc.Request("test3", []byte("help me"), 3*time.Second)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("help answer : %s\n", string(msg.Data))
	//}
	// 清空緩衝
	err = nc.Flush()
	if err != nil {
		log.Fatal("清空失敗")
	}

	time.Sleep(3 * time.Second)
}

func replyPublish(nc *nats.Conn) {

	reply := nats.NewInbox()
	i := 0

	for {
		// 自動建立一個唯一 subject
		// 發送訊息
		//nc.Publish("subject", []byte("Hello World!"))

		ms := "Hello world  " + strconv.Itoa(i)
		err := nc.PublishRequest("subject", reply, []byte(ms))
		if err != nil {
			log.Fatal("送不出去")
		}
		fmt.Println("send ", ms)

		i++

		time.Sleep(1 * time.Second)
	}
}

func jetStreamPublish(nc *nats.Conn) {
	js, _ := nc.JetStream()
	err := js.DeleteStream("Stream名稱_pub3")
	if err != nil {
		fmt.Println(err)
	}

	// 建立 Stream
	_, err = js.AddStream(&nats.StreamConfig{
		Name: "Stream名稱_pub3",
		Subjects: []string{
			"testTopic8.*", // 支援 wildcard
		},
		Storage:   nats.FileStorage,     // 儲存的方式 (預設 FileStorage)
		Retention: nats.WorkQueuePolicy, // 保留的策略
		Discard:   nats.DiscardOld,      // 丟棄的策略
		// ...
	})
	if err != nil {
		log.Fatalf("建立 Stream 失敗: %v", err)
	}

	i := 0

	for {

		ms := "Hello world  " + strconv.Itoa(i)
		_, err = js.Publish("testTopic8.12345", []byte(ms))
		if err != nil {
			log.Fatal("送不出去, err = ", err)
		}
		fmt.Println("send ", ms)

		i++

		time.Sleep(1 * time.Second)
	}
}

func jetStreamPubTestForDDU(nc *nats.Conn) {
	js, _ := nc.JetStream()

	_, err := js.AddStream(&nats.StreamConfig{
		Name: "webapiStreamChannel",
		Subjects: []string{
			"testTopic8.*", "chatRecordChannel", "sendGiftChannel", // 支援 wildcard
		},
		Storage:   nats.FileStorage,     // 儲存的方式 (預設 FileStorage)
		Retention: nats.WorkQueuePolicy, // 保留的策略
		Discard:   nats.DiscardOld,      // 丟棄的策略
		// ...
	})
	if err != nil {
		log.Fatalf("建立 Stream 失敗: %v", err)
	}

	i := 0
	for {
		// 发布-订阅 模式，向 test1 发布一个 `Hello World` 数据
		chat := ChatHistory{
			Id:         int64(i),
			Uid:        555,
			RoomId:     444,
			CreateTime: 1655273431,
			Content:    "鬼塚英吉",
		}

		chatMarshal, err := json.Marshal(chat)
		if err != nil {
			return
		}
		_, err = js.Publish("chatRecordChannel", []byte(chatMarshal))
		if err != nil {
			log.Fatal("送不出去, err = ", err)
		}
		fmt.Println("send chat ", i)

		gift := SendGiftReq{
			Anchorid: i,
			Giftid:   555,
			Liveid:   444,
			Count:    666,
			Uid:      777,
		}
		giftMarshal, err := json.Marshal(gift)
		if err != nil {
			return
		}
		_, err = js.Publish("sendGiftChannel", []byte(giftMarshal))
		if err != nil {
			log.Fatal("送不出去, err = ", err)
		}
		fmt.Println("send gift ", i)

		// 队列 模式，发布是一样的，只是订阅不同，向 test2 发布一个 `Hello zngw` 数据
		//_ = nc.Publish("test.2", []byte("鬼塚英吉"))

		// 请求-响应， 向 test3 发布一个 `help me` 请求数据，设置超时间3秒，如果有多个响应，只接收第一个收到的消息
		//msg, err := nc.Request("test3", []byte("help me"), 3*time.Second)
		//if err != nil {
		//	fmt.Println(err)
		//} else {
		//	fmt.Printf("help answer : %s\n", string(msg.Data))
		//}

		time.Sleep(2 * time.Second)
		i++
	}
}
