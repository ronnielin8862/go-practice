package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Create server connection
	nc, _ := nats.Connect("nats://127.0.0.1:4222")
	defer nc.Close()
	jetStream(nc)
}

func generalSubscript(nc *nats.Conn) {
	// Create server connection
	nc, _ = nats.Connect("nats://127.0.0.1:4222")
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

	// only for here, it is no need in general project
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	Sub1Cb.Unsubscribe()
	Sub2Cb.Unsubscribe()
}

func replySubscript(nc *nats.Conn) {

	var Sub1Cb *nats.Subscription

	// 接收訊息
	Sub1Cb, err := nc.Subscribe("subject", func(msg *nats.Msg) {
		fmt.Println("收到了", string(msg.Data))

		msg.Respond([]byte("i got it")) // 生產者會監聽 reply，來確認消費者有沒有收到
	})
	if err != nil {
		log.Fatal("訂閱失敗")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	Sub1Cb.Unsubscribe()
}

func jetStream(nc *nats.Conn) {

	// Create server connection
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("取得 JetStream 的 Context 失敗: %v", err)
	}

	// 建立 Stream
	_, err = js.AddStream(&nats.StreamConfig{
		Name: "Stream名稱_pub1",
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

	mcbAny := func(msg *nats.Msg) {
		log.Println("receive:", string(msg.Data))

	}

	_, err = js.Subscribe("testTopic8.*", mcbAny)
	if err != nil {
		log.Println("queue subscribe testTopic.*:", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func jetStreaForDDU(nc *nats.Conn) {
	js, err := nc.JetStream()

	// build stream
	_, err = js.AddStream(&nats.StreamConfig{
		Name: "webapiStreamChannel",
		Subjects: []string{
			"chatRecordChannel", "sendGiftChannel", // 支援 wildcard
		},
		Storage:   nats.FileStorage,     // 儲存的方式 (預設 FileStorage)
		Retention: nats.WorkQueuePolicy, // 保留的策略
		Discard:   nats.DiscardOld,      // 丟棄的策略
	})
	chatHistoryProc := func(msg *nats.Msg) {
		var chatHistory ChatHistory
		err = json.Unmarshal(msg.Data, &chatHistory)
		if err != nil {
			fmt.Println("SubMessage chatRecord InsertChatHistory err [%v]", err)
		}
		fmt.Println("chatHistory: ", chatHistory)
		//repository.InsertChatHistory(&chatHistory)
	}

	sendGiftProc := func(msg *nats.Msg) {
		var sendGiftReq SendGiftReq
		json.Unmarshal(msg.Data, &sendGiftReq)
		fmt.Println("sendGiftReq: ", sendGiftReq)
		//_, err = services.SendGift(&sendGiftReq, sendGiftReq.Uid)
		if err != nil {
			fmt.Println("发送礼物失败%v,礼物数据:%+v", err, msg.Data)
		}
	}

	_, err = js.Subscribe("chatRecordChannel", chatHistoryProc)
	if err != nil {
		fmt.Println("SubMessage ChatRecord err [%v]", err)
	}
	_, err = js.Subscribe("sendGiftChannel", sendGiftProc)
	if err != nil {
		fmt.Println("SubMessage SendGift err [%v]", err)
	}
}

func natsStreamForDDU(nc *nats.Conn) {
	_, err := GetNats().Subscribe(ChatRecordChannel, ChatRecordHandler, stan.DurableName(ChatRecordChannel))
	if err != nil {
		fmt.Println("订阅top%s失败,err:%v", ChatRecordChannel, err)
	}
	_, err = GetNats().Subscribe(SendGiftChannel, SendGiftHandler, stan.DurableName(SendGiftChannel))
	if err != nil {
		fmt.Println("订阅top%s失败,err:%v", ChatRecordChannel, err)
	}
}

//聊天记录处理
func ChatRecordHandler(msg *stan.Msg) {
	var chatHistory ChatHistory
	fmt.Println("chatHistory: ", chatHistory)
	json.Unmarshal(msg.Data, &chatHistory)
}

//礼物处理
func SendGiftHandler(msg *stan.Msg) {
	var sendGiftReq SendGiftReq
	fmt.Println("sendGiftReq: ", sendGiftReq)
	json.Unmarshal(msg.Data, &sendGiftReq)

}

const (
	WebapiStreamChannel = "webapiStreamChannel"
	ChatRecordChannel   = "JetStream.ChatRecordSubject"
	SendGiftChannel     = "JetStream.SendGiftSubject"
	AttentAnchorChannel = "attentAnchorChannel"
	ChangePropChannel   = "changePropChannel"
)

type SendGiftReq struct {
	Anchorid int   `json:"anchorid" validate:"required"`
	Giftid   int   `json:"giftid" validate:"required"`
	Liveid   int64 `json:"liveid"`
	Count    int   `json:"count"`
	Uid      int
}

type ChatHistory struct {
	Id         int64  `json:"id" gorm:"primaryKey;autoIncrement"` // bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	Uid        int    `json:"uid"`                                //int(11) DEFAULT NULL COMMENT '用户id',
	RoomId     int    `json:"room_id"`                            //varchar(100) DEFAULT NULL COMMENT '房间id',
	CreateTime int64  `json:"create_time"`                        //datetime DEFAULT NULL,
	Content    string `json:"content"`
}

var NatsDB stan.Conn

func InitNats() error {
	fName := "InitNats"
	err := GetNatsConn(
		"0.0.0.0",
		"4222",
		"nats%3admin##1",
		"oscars3higehaohaizi",
		"test-cluster",
		"natsClicent01")
	if err != nil {
		return fmt.Errorf("%s Init fail %s", fName, err.Error())
	}
	return nil
}

func GetNatsConn(host, port, user, passwd, stanClusterID, clientID string) error {
	url := fmt.Sprintf("nats://%s:%s", host, port)
	nc, err := nats.Connect(
		url,
		nats.UserInfo(user, passwd),
		nats.Timeout(time.Second*10),
		nats.PingInterval(time.Second*4),
	)
	if err != nil {
		fmt.Println("error by nats connect: %v", err)
	}
	NatsDB, err = stan.Connect(stanClusterID, clientID, stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: %v\n", reason)
		}))
	if err != nil {
		fmt.Println("error by nats connect: %v", err)
	}

	return nil
}

func GetNats() stan.Conn {
	return NatsDB
}
