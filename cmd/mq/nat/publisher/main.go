package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/ronnielin8862/go-practice/globle"
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
	url := fmt.Sprintf("nats://127.0.0.1:4222")
	//url := fmt.Sprintf("nats://%s:%s", "52.221.194.38", "4344")
	nc, _ := nats.Connect(
		url,
		nats.UserInfo("nats%3admin##1", "oscars3higehaohaizi"),
		nats.Timeout(time.Second*10),
		nats.PingInterval(time.Second*4),
	)
	//jetStreamPubTestForDDU(nc)
	//natsStreamingForDDUMatchTextLive(nc)
	//natsStreamingForDDUScoreLive(nc)
	natsStreamingForDDUStatsLive(nc)
}

func natsStreamingForDDUScoreLive(nc *nats.Conn) {
	NatsDB, err := stan.Connect("test-cluster", "natsClicent07", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: ", reason)
		}))
	if err != nil {
		fmt.Println("error by nats connect: %v", err)
	}
	subject := fmt.Sprint(globle.FootballScoreLive)

	for i := 1; i <= wantRun; i++ {
		homeScore := Score{
			Score:        i,
			HalfScore:    2,
			RedCard:      1,
			YellowCard:   3,
			CornerKick:   99,
			OTScore:      1,
			PenaltyScore: 3,
		}
		awayScore := Score{
			Score:        4,
			HalfScore:    2,
			RedCard:      1,
			YellowCard:   3,
			CornerKick:   99,
			OTScore:      1,
			PenaltyScore: 3,
		}
		roomScore := RoomScoreLive{
			Id:          6000098,
			Status:      i,
			HomeScore:   homeScore,
			AwayScore:   awayScore,
			KickOutTime: 1653249600,
		}

		mjs := []RoomScoreLive{roomScore}
		mj, _ := json.Marshal(mjs)
		err = NatsDB.Publish(subject, mj)
		if err != nil {
			fmt.Println("送不出去, err = ", err)
		}

		fmt.Println("publish success : ", roomScore)
		time.Sleep(3 * time.Second)
	}
}

var (
	esCount = 0
	wantRun = 1
)

func natsStreamingForDDUStatsLive(nc *nats.Conn) {
	NatsDB, err := stan.Connect("test-cluster", "natsClicent07", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: ", reason)
		}))
	if err != nil {
		fmt.Println("error by nats connect: ", err)
	}
	subject := fmt.Sprint(globle.FootballStatsLive)
	fmt.Println("subject = ", subject)
	// 假資料組成
	for i := 1; i <= wantRun; i++ {
		msgs := statsLiveMock(i)

		mj, _ := json.Marshal(msgs)
		err = NatsDB.Publish(subject, mj)
		if err != nil {
			fmt.Println("送不出去, err = ", err)
		}

		time.Sleep(3 * time.Second)
	}
}

func statsLiveMock(wantRun int) (msgs []RoomStatsLiveMessage) {
	for i := 1; i <= esCount+wantRun; i++ {
		msg := RoomStatsLiveMessage{
			SentMessageStruct: SentMessageStruct{
				Type: globle.StatusLive,
			},
			Id:   6000098,
			Type: 3,
			Home: 2,
			Away: 3,
		}
		msgs = append(msgs, msg)
	}

	fmt.Println("last msgs: ", msgs[esCount+wantRun-1])
	return msgs
}

func natsStreamingForDDUMatchTextLive(nc *nats.Conn) {
	NatsDB, err := stan.Connect("test-cluster", "natsClicent07", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: %v\n", reason)
		}))
	if err != nil {
		fmt.Println("error by nats connect: %v", err)
	}
	subject := fmt.Sprint(globle.FootballTextLive)
	fmt.Println("subject = ", subject)
	// 假資料組成
	for i := 1; i <= wantRun; i++ {
		msgs := textLiveMock(i)

		mj, _ := json.Marshal(msgs)
		err = NatsDB.Publish(subject, mj)
		if err != nil {
			fmt.Println("送不出去, err = ", err)
		}

		time.Sleep(3 * time.Second)
	}
}

func textLiveMock(wantRun int) (msgs []RoomTextLiveMessage) {
	for i := 1; i <= esCount+wantRun; i++ {
		msg := RoomTextLiveMessage{
			SentMessageStruct: SentMessageStruct{
				Type:    globle.TextLive,
				Message: "test",
			},
			Id:       6000098,
			Time:     strconv.Itoa(i),
			Type:     3,
			Data:     "14' - 第1张黄牌，裁判出示了本场比赛的第一张黄牌，给了(夏洛特独立)",
			Position: 2,
			Main:     1,
		}
		msgs = append(msgs, msg)
	}

	fmt.Println("last msgs: ", msgs[esCount+wantRun-1])
	return msgs
}

type RoomScoreLiveMessage struct {
	SentMessageStruct
	ScoreLive RoomScoreLive `json:"score_live"`
}

type RoomScoreLive struct {
	Id          int   `json:"match_id"`
	Status      int   `json:"match_status"`
	HomeScore   Score `json:"home_score"`
	AwayScore   Score `json:"away_score"`
	KickOutTime int64 `json:"kick_out_time"`
}

type Score struct {
	Score        int `json:"score"`
	HalfScore    int `json:"half_score"`
	RedCard      int `json:"red_card"`
	YellowCard   int `json:"yellow_card"`
	CornerKick   int `json:"corner_kick"`
	OTScore      int `json:"ot_score"`
	PenaltyScore int `json:"penalty_score"`
}

type RoomStatsLiveMessage struct {
	SentMessageStruct
	Id   int64 `json:"match_id"`
	Type int   `json:"type"`
	Home int   `json:"home"`
	Away int   `json:"away"`
}

type RoomTextLiveMessage struct {
	SentMessageStruct
	Id         int64  `json:"match_id"`    // 赛事id
	Time       string `json:"time"`        // 事件时间
	Type       int8   `json:"type"`        // 事件类型
	Data       string `json:"data"`        // 事件文本
	Position   int8   `json:"position"`    // 事件發生方， 0-中立 1-主队 2-客队
	Main       int8   `json:"main"`        // 是否重要事件 0-否 1-是
	CreateTime int64  `json:"create_time"` // 創建時間
}

type SentMessageStruct struct {
	Type    string `json:"type"`
	Message string `default:"" json:"message,omitempty"`
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

func natsStreamingForDDU(nc *nats.Conn) {

	NatsDB, err := stan.Connect("test-cluster", "natsClicent01", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: %v\n", reason)
		}))
	if err != nil {
		fmt.Println("error by nats connect: %v", err)
	}

	fmt.Println("JetStream.ChatRecordSubject = ", "DDDFDFDFDFDFDFDF")
	err = NatsDB.Publish("JetStream.ChatRecordSubject", []byte("DDDFDFDFDFDFDFDF"))
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	}

	err = NatsDB.Publish("JetStream.SendGiftSubject", []byte("dfdfdffdfeefefeffe"))
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	}

	time.Sleep(time.Second * 2)

	NatsDB.Subscribe("JetStream.ChatRecordSubject", ChatRecordHandler, stan.DurableName("JetStream.ChatRecordSubject"))
	if err != nil {
		fmt.Println("订阅top%s失败,err:%v", "JetStream.ChatRecordSubject", err)
	}
	time.Sleep(time.Second * 2)

	// Unsubscribe
	//sub.Unsubscribe()

	// Close connection
	NatsDB.Close()

}

//聊天记录处理
func ChatRecordHandler(msg *stan.Msg) {
	fmt.Println("收到聊天记订阅消息:%v", string(msg.Data))

	var chatHistory ChatHistory
	err := json.Unmarshal(msg.Data, &chatHistory)
	if err != nil {
		fmt.Println("SubMessage chatRecord InsertChatHistory err [%v]", err)
	}
	fmt.Println("chatHistory: ", chatHistory)
}

//
//func main() {
//	publisher, err := nats.NewStreamingPublisher(
//		nats.StreamingPublisherConfig{
//			ClusterID: "test-cluster",
//			ClientID:  "natsClicent01",
//			StanOptions: []stan.Option{
//				stan.NatsURL("nats://nats-streaming:4222"),
//			},
//			Marshaler: nats.GobMarshaler{},
//		},
//		watermill.NewStdLogger(false, false),
//	)
//	if err != nil {
//		panic(err)
//	}
//}
