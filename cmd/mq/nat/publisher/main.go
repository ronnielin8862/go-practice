package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/ronnielin8862/go-practice/globle"
	"log"
	"strconv"
	"time"
)

func main() {
	url := fmt.Sprintf("nats://127.0.0.1:4222")
	//url := fmt.Sprintf("nats://%s:%s", "52.221.194.38", "4344")
	nc, _ := nats.Connect(
		url,
		nats.UserInfo("nats%3admin##1", "oscars3higehaohaizi"),
		nats.Timeout(time.Second*10),
		nats.PingInterval(time.Second*4),
	)
	NatsDB, err := stan.Connect("test-cluster", "natsClicent07", stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Println("Connection lost, reason: ", reason)
		}))
	if err != nil {
		fmt.Println("error by nats connect: ", err)
	}
	//jetStreamPubTestForDDU(nc,NatsDB)
	//natsStreamingForDDUMatchTextLive(NatsDB)
	//natsStreamingForDDUScoreLive(NatsDB)
	//natsStreamingForDDUStatsLive(NatsDB)
	//natsStreamingForDDULineup(NatsDB)
	natsStreamingForDDUBasketballText(NatsDB)
}

func natsStreamingForDDUBasketballText(NatsDB stan.Conn) {

	subject := fmt.Sprint(globle.BasketballTextLive)
	fmt.Println("subject = ", subject)

	data := generateBasketballData(wantRun)

	lineupJson, err := json.Marshal(data)
	//err = NatsDB.Publish(subject, lineupJson)
	err = NatsDB.Publish(subject, lineupJson)
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	}

	fmt.Println("publish success : ", string(lineupJson))
	time.Sleep(3 * time.Second)
}

func generateBasketballData(wantRun int) (msgs []globle.BasketballText) {

	for i := 1; i <= wantRun; i++ {
		var m globle.BasketballText
		m.MatchId = 6000098
		m.Time = "10"
		m.EventTeam = int8(i)
		m.AwayScore = 444
		m.HomeScore = i
		m.Text = "23456"
		msgs = append(msgs, m)
	}
	return msgs
}

func natsStreamingForDDULineup(NatsDB stan.Conn) {

	subject := fmt.Sprint(globle.FootballLineupLive)
	fmt.Println("subject = ", subject)
	lineup := lineupLiveMock()

	lineupJson, err := json.Marshal(lineup)
	//err = NatsDB.Publish(subject, lineupJson)
	err = NatsDB.Publish(subject, lineupLiveData2())
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	}

	fmt.Println("publish success : ", string(lineupJson))
	time.Sleep(3 * time.Second)
}

func lineupLiveMock() []globle.Lineup {
	lineup := globle.Lineup{
		MatchId: 1, Confirmed: 0, HomeFormation: "4-4-2", AwayFormation: "4-3-3",
	}

	var incidentsA *globle.Incidents
	var incidentsB globle.Incidents
	var incidentsC globle.Incidents
	var incidentsD globle.Incidents

	incidentsA = &globle.Incidents{
		Type:       1,
		Time:       "A",
		Belong:     2,
		HomeScore:  3,
		AwayScore:  4,
		ReasonType: 5,
		Player:     globle.Player{PlayerId: 9, Name: "AAA"},
		Assist1:    globle.Player{PlayerId: 10, Name: "BBB"},
		InPlayer:   globle.Player{PlayerId: 12, Name: "DDD"},
		OutPlayer:  globle.Player{PlayerId: 13, Name: "EEE"},
	}

	copier.Copy(&incidentsB, &incidentsA)
	copier.Copy(&incidentsC, &incidentsA)
	copier.Copy(&incidentsD, &incidentsA)

	var lineupItemA *globle.LineupItem
	var lineupItemB globle.LineupItem
	var lineupItemC globle.LineupItem
	var lineupItemD globle.LineupItem
	lineupItemA = &globle.LineupItem{
		LineupId:     1,
		TeamId:       2,
		First:        3,
		Captain:      4,
		Name:         "A",
		Logo:         "B",
		NationalLogo: "C",
		ShirtNumber:  5,
		Position:     "D",
		X:            6,
		Y:            7,
		Rating:       "E",
		Incidents:    nil,
	}

	lineupItemA.Incidents = append(lineupItemA.Incidents, *incidentsA)
	lineupItemB.Incidents = append(lineupItemB.Incidents, incidentsB)
	lineupItemB.Incidents = append(lineupItemB.Incidents, incidentsC)
	lineupItemB.Incidents = append(lineupItemB.Incidents, incidentsD)

	copier.Copy(&lineupItemB, &lineupItemA)
	copier.Copy(&lineupItemC, &lineupItemA)
	copier.Copy(&lineupItemD, &lineupItemA)
	lineupItemB.LineupId = 11
	lineupItemC.LineupId = 12
	lineupItemD.LineupId = 13

	lineup.Home = append(lineup.Home, *lineupItemA)
	lineup.Away = append(lineup.Away, lineupItemB)
	//lineup.Away = append(lineup.Away, lineupItemC)
	//lineup.Away = append(lineup.Away, lineupItemD)

	var lp []globle.Lineup
	lineup2 := lineup
	lp = append(lp, lineup2)
	lineup2.MatchId = 2

	return lp
}

func lineupLiveData2() []byte {
	return []byte("[\n    {\n        \"match_id\": 6000098,\n        \"confirmed\": 1,\n        \"home_formation\": \"4-3-9\",\n        \"away_formation\": \"4-4-2\",\n        \"home\": [\n            {\n                \"id\": 1414339,\n                \"team_id\": 14702,\n                \"first\": 1,\n                \"captain\": 0,\n                \"name\": \"居琼·比雅尼·布林乔尔夫森3\",\n                \"logo\": \"\",\n                \"national_logo\": \"\",\n                \"shirt_number\": 16,\n                \"position\": \"\",\n                \"x\": 0,\n                \"y\": 0,\n                \"rating\": \"0.0\",\n                \"incidents\": [\n                    {\n                        \"type\": 1,\n                        \"time\": \"9\",\n                        \"belong\": 1,\n                        \"home_score\": 2,\n                        \"away_score\": 0,\n                        \"reason_type\": 0,\n                        \"player\": {\n                            \"id\": 1414339,\n                            \"name\": \"居琼·比雅尼·布林乔尔夫森\"\n                        },\n                        \"assist1\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist2\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"in_player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"out_player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        }\n                    },\n                    {\n                        \"id\": 1414340,\n                        \"team_id\": 14702,\n                        \"first\": 1,\n                        \"captain\": 0,\n                        \"name\": \"埃尔瓦尔·鲍德温森\",\n                        \"logo\": \"\",\n                        \"national_logo\": \"\",\n                        \"shirt_number\": 18,\n                        \"position\": \"\",\n                        \"x\": 0,\n                        \"y\": 0,\n                        \"rating\": \"0.0\"\n                    },\n                    {\n                        \"type\": 1,\n                        \"time\": \"85\",\n                        \"belong\": 1,\n                        \"home_score\": 5,\n                        \"away_score\": 0,\n                        \"reason_type\": 0,\n                        \"player\": {\n                            \"id\": 1414339,\n                            \"name\": \"居琼·比雅尼·布林乔尔夫森\"\n                        },\n                        \"assist1\": {\n                            \"id\": 1513567,\n                            \"name\": \"西格弗斯·冈纳森·范纳尔\"\n                        },\n                        \"assist2\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"in_player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"out_player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        }\n                    }\n                ]\n            },\n            {\n                \"id\": 1433617,\n                \"team_id\": 14702,\n                \"first\": 1,\n                \"captain\": 0,\n                \"name\": \"克里斯托弗·克里斯蒂安松\",\n                \"logo\": \"\",\n                \"national_logo\": \"\",\n                \"shirt_number\": 15,\n                \"position\": \"\",\n                \"x\": 0,\n                \"y\": 0,\n                \"rating\": \"0.0\"\n            }\n        ],\n        \"away\": [\n            {\n                \"id\": 1146124,\n                \"team_id\": 24328,\n                \"first\": 1,\n                \"captain\": 0,\n                \"name\": \"詹姆斯·戴尔3\",\n                \"logo\": \"https://cdn.sportnanoapi.com/football/player/3bad34d6db39da48c94665c7989d4f7c.png\",\n                \"national_logo\": \"\",\n                \"shirt_number\": 4,\n                \"position\": \"\",\n                \"x\": 0,\n                \"y\": 0,\n                \"rating\": \"0.3\",\n                \"incidents\": [\n                    {\n                        \"type\": 9,\n                        \"time\": \"46\",\n                        \"belong\": 2,\n                        \"home_score\": 3,\n                        \"away_score\": 0,\n                        \"reason_type\": 0,\n                        \"player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist1\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist2\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"in_player\": {\n                            \"id\": 1553739,\n                            \"name\": \"haukur palsson\"\n                        },\n                        \"out_player\": {\n                            \"id\": 1146124,\n                            \"name\": \"詹姆斯·戴尔\"\n                        }\n                    }\n,{\n                        \"type\": 9,\n                        \"time\": \"46\",\n                        \"belong\": 2,\n                        \"home_score\": 3,\n                        \"away_score\": 0,\n                        \"reason_type\": 0,\n                        \"player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist1\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist2\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"in_player\": {\n                            \"id\": 1553739,\n                            \"name\": \"haukur palsson\"\n                        },\n                        \"out_player\": {\n                            \"id\": 1146124,\n                            \"name\": \"詹姆斯·戴尔\"\n                        }\n                    }\n                ]\n            },\n            {\n                \"id\": 1145980,\n                \"team_id\": 24328,\n                \"first\": 1,\n                \"captain\": 0,\n                \"name\": \"什克尔岑·瓦西里1\",\n                \"logo\": \"\",\n                \"national_logo\": \"\",\n                \"shirt_number\": 11,\n                \"position\": \"\",\n                \"x\": 0,\n                \"y\": 0,\n                \"rating\": \"0.5\"\n            }\n        ]\n    }\n]")
}

func natsStreamingForDDUScoreLive(NatsDB stan.Conn) {

	subject := fmt.Sprint(globle.FootballScoreLive)

	for i := 1; i <= wantRun; i++ {
		homeScore := globle.Score{
			Score:        i,
			HalfScore:    2,
			RedCard:      1,
			YellowCard:   3,
			CornerKick:   99,
			OTScore:      1,
			PenaltyScore: 3,
		}
		awayScore := globle.Score{
			Score:        4,
			HalfScore:    2,
			RedCard:      1,
			YellowCard:   3,
			CornerKick:   99,
			OTScore:      1,
			PenaltyScore: 3,
		}
		roomScore := globle.RoomScoreLive{
			Id:          6000098,
			Status:      i,
			HomeScore:   homeScore,
			AwayScore:   awayScore,
			KickOutTime: 1653249600,
		}

		mjs := []globle.RoomScoreLive{roomScore}
		mj, _ := json.Marshal(mjs)
		err := NatsDB.Publish(subject, mj)
		if err != nil {
			fmt.Println("送不出去, err = ", err)
		}

		fmt.Println("publish success : ", roomScore)
		time.Sleep(3 * time.Second)
	}
}

var (
	esCount    = 0
	statsCount = 1
	wantRun    = 8
)

func natsStreamingForDDUStatsLive(NatsDB stan.Conn) {
	var (
		statsCount = 9
	)

	subject := fmt.Sprint(globle.FootballStatsLive)
	fmt.Println("subject = ", subject)
	// 假資料組成

	msgs := statsLiveMock(statsCount)

	mj, _ := json.Marshal(msgs)
	err := NatsDB.Publish(subject, mj)
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	} else {
		fmt.Println("publish success : ", string(mj))
	}

	time.Sleep(3 * time.Second)

}

func statsLiveMock(statsCount int) (msgs []globle.StatsLiveMessage) {
	for i := 1; i <= statsCount; i++ {
		stats := globle.StatsLiveMessage{
			Id:   6000098,
			Type: wantRun,
			Home: 4 + wantRun,
			Away: 3 + wantRun,
		}
		msgs = append(msgs, stats)
	}
	return msgs
}

func natsStreamingForDDUMatchTextLive(NatsDB stan.Conn) {

	subject := fmt.Sprint(globle.FootballTextLive)
	fmt.Println("subject = ", subject)
	// 假資料組成

	msgs := textLiveMock(wantRun)

	mj, _ := json.Marshal(msgs)
	err := NatsDB.Publish(subject, mj)
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	} else {
		fmt.Println("publish success : ", string(mj))
	}

	time.Sleep(3 * time.Second)

}

func textLiveMock(wantRun int) (msgs []globle.RoomTextLiveMessage) {
	for i := 1; i <= esCount+wantRun; i++ {
		msg := globle.RoomTextLiveMessage{
			SentMessageStruct: globle.SentMessageStruct{
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

	return msgs
}

func generalPublish(nc *nats.Conn) {

	// 发布-订阅 模式，向 test1 发布一个 `Hello World` 数据
	chat := globle.ChatHistory{
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

	gift := globle.SendGiftReq{
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
		chat := globle.ChatHistory{
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

		gift := globle.SendGiftReq{
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

func natsStreamingForDDU(NatsDB stan.Conn) {

	fmt.Println("JetStream.ChatRecordSubject = ", "DDDFDFDFDFDFDFDF")
	err := NatsDB.Publish("JetStream.ChatRecordSubject", []byte("DDDFDFDFDFDFDFDF"))
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

	var chatHistory globle.ChatHistory
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
