package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/ronnielin8862/go-practice/globle"
	"io/ioutil"
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
	NatsDB, err := stan.Connect("test-cluster", "natsClicent107", stan.NatsConn(nc),
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
	//natsStreamingForDDUBasketballText(NatsDB)
	//natsStreamingForDDUBasketballScore(NatsDB)
	//natsStreamingForDDUBasketballstats(NatsDB)
	natsStreamingForDDUBasketballRecord(NatsDB)
}

func natsStreamingForDDUBasketballRecord(NatsDB stan.Conn) {
	subject := fmt.Sprint(globle.BasketballRecordLive)
	fmt.Println("subject = ", subject)

	file, err := ioutil.ReadFile("/Users/ronnie/Downloads/Telegram Desktop/篮球阵容数据 (1).json")
	data := file

	lineupJson, err := json.Marshal(data)
	//err = NatsDB.Publish(subject, lineupJson)
	err = NatsDB.Publish(subject, file)
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	}

	fmt.Println("publish success : ", string(lineupJson))
	time.Sleep(3 * time.Second)
}

func natsStreamingForDDUBasketballstats(NatsDB stan.Conn) {
	subject := fmt.Sprint(globle.BasketballStatsLive)
	fmt.Println("subject = ", subject)

	data := basketBallStatsLiveMock()

	lineupJson, err := json.Marshal(data)
	//err = NatsDB.Publish(subject, lineupJson)
	err = NatsDB.Publish(subject, lineupJson)
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	}

	fmt.Println("publish success : ", string(lineupJson))
	time.Sleep(3 * time.Second)
}

func basketBallStatsLiveMock() (msgs []globle.BasketStatsLiveMessage) {
	for i := 1; i <= 9; i++ {
		stats := globle.BasketStatsLiveMessage{
			Id:   6000098,
			Type: i,
			Home: float64(i),
			Away: float64(3 + wantRun),
		}
		msgs = append(msgs, stats)
	}
	return msgs
}

var (
	esCount    = 0
	statsCount = 1
	wantRun    = 1
)

func natsStreamingForDDUBasketballScore(NatsDB stan.Conn) {
	subject := fmt.Sprint(globle.BasketballScoreLive)
	fmt.Println("subject = ", subject)

	data := generateBasketballScore(wantRun)

	lineupJson, err := json.Marshal(data)
	//err = NatsDB.Publish(subject, lineupJson)
	err = NatsDB.Publish(subject, lineupJson)
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	}

	fmt.Println("publish success : ", string(lineupJson))
	time.Sleep(3 * time.Second)
}

func generateBasketballScore(run int) (BsScores []globle.BasketballScore) {

	for i := 1; i <= 1; i++ {
		bs := globle.BasketballScore{
			MatchId:     6000098,
			MatchStatus: 2,
			TimeLeft:    15 - i,
			AwayScore:   []int{6, 2, 3, 4, 5},
			HomeScore:   []int{1, 2, 3, 4, 5},
		}

		BsScores = append(BsScores, bs)
	}
	return BsScores
}

func natsStreamingForDDUBasketballText(NatsDB stan.Conn) {

	subject := fmt.Sprint(globle.BasketballTextLive)
	fmt.Println("subject = ", subject)

	data := generateBasketballData(5)

	lineupJson, err := json.Marshal(data)
	//err = NatsDB.Publish(subject, lineupJson)
	err = NatsDB.Publish(subject, lineupJson)
	if err != nil {
		fmt.Println("送不出去, err = ", err)
	}

	fmt.Println("publish success : ", string(lineupJson))
	time.Sleep(3 * time.Second)
}

func generateBasketballData(wantRun int) (msgs []globle.TextLiveStruct) {

	for i := wantRun - 1; i >= 0; i-- {
		var m globle.TextLiveStruct
		m.Id = 3666736
		m.Time = strconv.Itoa(i)
		m.Position = int8(i)
		m.AwayScore = i
		m.HomeScore = i
		m.Data = "23456"
		msgs = append(msgs, m)
	}
	return msgs

	//for i := 1; i <= wantRun; i++ {
	//	var m globle.TextLiveStruct
	//	m.Id = 3666736
	//	m.Time = "10"
	//	m.Position = int8(i)
	//	m.AwayScore = wantRun
	//	m.HomeScore = i
	//	m.Data = "23456"
	//	msgs = append(msgs, m)
	//}
	//return msgs
}

func natsStreamingForDDULineup(NatsDB stan.Conn) {

	subject := fmt.Sprint(globle.FootballLineupLive)
	fmt.Println("subject = ", subject)
	file, err := ioutil.ReadFile("/Users/ronnie/Downloads/Telegram Desktop/test1.json")
	if err != nil {
		fmt.Println("err ")
	}
	var bs []globle.Lineup
	err = json.Unmarshal(file, &bs)
	if err != nil {
		fmt.Println("unmarshal err")
	}

	for i := 1; i <= 100; i++ {
		//lineup := lineupLiveMock(1 + i)
		bs[0].Home[0].Name = fmt.Sprintf("%s%d", "金城武4", i)
		marshal, err := json.Marshal(bs)
		if err != nil {
			return
		}
		//_, err := json.Marshal(lineup)
		//err = NatsDB.Publish(subject, lineupJson)
		err = NatsDB.Publish(subject, marshal)
		if err != nil {
			fmt.Println("送不出去, err = ", err)
		}

		fmt.Println("publish success : ", string(marshal))
		time.Sleep(60 * time.Second)
	}
}

func lineupLiveMock(i int) []globle.Lineup {
	lineup := globle.Lineup{
		MatchId: 3680783, Confirmed: 0, HomeFormation: "4-4-2", AwayFormation: "4-3-3",
	}

	var incidentsA *globle.Incidents
	var incidentsB globle.Incidents
	var incidentsC globle.Incidents
	var incidentsD globle.Incidents

	incidentsA = &globle.Incidents{
		Type:       1,
		Time:       "A",
		Belong:     2,
		HomeScore:  i,
		AwayScore:  i,
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

func lineupLiveData2(i string) []byte {
	return []byte("[\n    {\n    \"match_id\": 3680783,\n    \"confirmed\": 1,\n    \"home_formation\": \"\",\n    \"away_formation\": \"\",\n    \"home\": [\n        {\n            \"id\": 1510517,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"高蒂·阿诺·乌尔法松\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 2,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1510518,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 1,\n            \"name\": \"安德鲁·皮尤\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 44,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1554985,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"ruben lozano\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 19,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"74\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1146124,\n                        \"name\": \"詹姆斯·戴尔\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1554985,\n                        \"name\": \"ruben lozano\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1263285,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"迈克尔·凯德曼\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 14,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"64\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1264413,\n                        \"name\": \"乔恩·英加森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1263285,\n                        \"name\": \"迈克尔·凯德曼\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1510523,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"斯特凡·拉法尔·丹尼尔森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 12,\n            \"position\": \"G\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1510520,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"拉格纳·冈纳森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 6,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1510522,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"尼古拉·朱里克\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 22,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 3,\n                    \"time\": \"33\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 2,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 1510522,\n                        \"name\": \"尼古拉·朱里克\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                },\n                {\n                    \"type\": 9,\n                    \"time\": \"88\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1306339,\n                        \"name\": \"巴勃罗·加列戈\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1510522,\n                        \"name\": \"尼古拉·朱里克\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1160079,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"豪库尔·莱福尔·埃里克森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/f2d114a9ac6ad1f4cfa3678a5f6da31d.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 69,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 988178,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"达古尔·古德森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/d19b74e873e7db294548b6261d5705d2.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 27,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 971553,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"乌纳尔·阿里·汉森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/6d22d1117ad356d9a3ab854fa850db26.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 16,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 3,\n                    \"time\": \"63\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 971553,\n                        \"name\": \"乌纳尔·阿里·汉森\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1518353,\n            \"team_id\": 24328,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"亚历山大·赫尔加森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 10,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 3,\n                    \"time\": \"26\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 2,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 1518353,\n                        \"name\": \"亚历山大·赫尔加森\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1145980,\n            \"team_id\": 24328,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"什克尔岑·瓦西里\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 11,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1553739,\n            \"team_id\": 24328,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"haukur palsson\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 15,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1538818,\n            \"team_id\": 24328,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"艾萨克\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 1,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1264413,\n            \"team_id\": 24328,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"乔恩·英加森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 23,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"64\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1264413,\n                        \"name\": \"乔恩·英加森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1263285,\n                        \"name\": \"迈克尔·凯德曼\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 88343,\n            \"team_id\": 24328,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"阿克塞尔·弗赖尔·哈达森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/2b7237b5d17676e6b1f267a139d8b701.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 5,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1306339,\n            \"team_id\": 24328,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"巴勃罗·加列戈\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/57d0b6b18c876535c6bec45024213bdd.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 9,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"88\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1306339,\n                        \"name\": \"巴勃罗·加列戈\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1510522,\n                        \"name\": \"尼古拉·朱里克\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1146124,\n            \"team_id\": 24328,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"詹姆斯·戴尔\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/3bad34d6db39da48c94665c7989d4f7c.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 4,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"74\",\n                    \"belong\": 1,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1146124,\n                        \"name\": \"詹姆斯·戴尔\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1554985,\n                        \"name\": \"ruben lozano\"\n                    }\n                }\n            ]\n        }\n    ],\n    \"away\": [\n        {\n            \"id\": 1224636,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"瓦尔·尼古拉斯·冈纳森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/c3c74d380d81ae0c7071c7b9fa623125.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 18,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 3,\n                    \"time\": \"73\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 1224636,\n                        \"name\": \"瓦尔·尼古拉斯·冈纳森\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                },\n                {\n                    \"type\": 9,\n                    \"time\": \"86\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1502044,\n                        \"name\": \"奥恩·阿伦·索瓦达森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1224636,\n                        \"name\": \"瓦尔·尼古拉斯·冈纳森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1091324,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 1,\n            \"name\": \"阿吉尔·埃索森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/da4aa1c0a62c29951b9d3d1a77e0d74b.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 2,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1395402,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"奥拉福尔·赫尔加森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 1,\n            \"position\": \"G\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 22536,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"亚斯基尔·博库尔·斯盖尔森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/50fb65e2305b73ae3192cbeba37be6c9.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 10,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 993006,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"阿诺·布雷基·阿瑟森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/473a7a05106b4cf4c5d83224a6f475a9.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 27,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1096922,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"比尔基·埃索森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/9f360e1bb044a0e3e829ecb8cf575995.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 17,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"73\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1395410,\n                        \"name\": \"哈尔·索尔斯泰森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1096922,\n                        \"name\": \"比尔基·埃索森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1510139,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"达里乌斯·贝内迪克特·加德森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 28,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 1,\n                    \"time\": \"11\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 2,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 1180944,\n                        \"name\": \"索杜尔·贡纳尔·哈夫森\"\n                    },\n                    \"assist1\": {\n                        \"id\": 1510139,\n                        \"name\": \"达里乌斯·贝内迪克特·加德森\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1224634,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"阿尔纳·琼森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/5c6b2e23f7c1a1b5f706babe55e5120f.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 4,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 1,\n                    \"time\": \"45\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 1224634,\n                        \"name\": \"阿尔纳·琼森\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1538686,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"马蒂亚斯·劳尔森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 9,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 1,\n                    \"time\": \"5\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 1,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 1538686,\n                        \"name\": \"马蒂亚斯·劳尔森\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                },\n                {\n                    \"type\": 3,\n                    \"time\": \"64\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 1538686,\n                        \"name\": \"马蒂亚斯·劳尔森\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                },\n                {\n                    \"type\": 9,\n                    \"time\": \"70\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1395403,\n                        \"name\": \"奥斯卡·伯格奥尔森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1538686,\n                        \"name\": \"马蒂亚斯·劳尔森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 62687,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"奥里·斯文·史蒂芬森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/0eafe41f07e044e8452834108a2ce6ba.jpg\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 5,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"86\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1263185,\n                        \"name\": \"阿克塞尔·曼尼·古比约恩森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 62687,\n                        \"name\": \"奥里·斯文·史蒂芬森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1180944,\n            \"team_id\": 10352,\n            \"first\": 1,\n            \"captain\": 0,\n            \"name\": \"索杜尔·贡纳尔·哈夫森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/c8da0fa681a4cfe7b0a17a77e62dc27e.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 11,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 1,\n                    \"time\": \"11\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 2,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 1180944,\n                        \"name\": \"索杜尔·贡纳尔·哈夫森\"\n                    },\n                    \"assist1\": {\n                        \"id\": 1510139,\n                        \"name\": \"达里乌斯·贝内迪克特·加德森\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"out_player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    }\n                },\n                {\n                    \"type\": 9,\n                    \"time\": \"70\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1466009,\n                        \"name\": \"比约恩·奥马尔·斯蒂芬森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1180944,\n                        \"name\": \"索杜尔·贡纳尔·哈夫森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1502044,\n            \"team_id\": 10352,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"奥恩·阿伦·索瓦达森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 19,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"86\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1502044,\n                        \"name\": \"奥恩·阿伦·索瓦达森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1224636,\n                        \"name\": \"瓦尔·尼古拉斯·冈纳森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1395410,\n            \"team_id\": 10352,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"哈尔·索尔斯泰森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 20,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"73\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1395410,\n                        \"name\": \"哈尔·索尔斯泰森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1096922,\n                        \"name\": \"比尔基·埃索森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1466009,\n            \"team_id\": 10352,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"比约恩·奥马尔·斯蒂芬森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 22,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"70\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1466009,\n                        \"name\": \"比约恩·奥马尔·斯蒂芬森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1180944,\n                        \"name\": \"索杜尔·贡纳尔·哈夫森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1395403,\n            \"team_id\": 10352,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"奥斯卡·伯格奥尔森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/f348233b50d676d3e1affbbef11092fb.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 77,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"70\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1395403,\n                        \"name\": \"奥斯卡·伯格奥尔森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 1538686,\n                        \"name\": \"马蒂亚斯·劳尔森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 1502045,\n            \"team_id\": 10352,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"拉芬·英加森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 31,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        },\n        {\n            \"id\": 1263185,\n            \"team_id\": 10352,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"阿克塞尔·曼尼·古比约恩森\",\n            \"logo\": \"\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 15,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\",\n            \"Incidents\": [\n                {\n                    \"type\": 9,\n                    \"time\": \"86\",\n                    \"belong\": 2,\n                    \"home_score\": 0,\n                    \"away_score\": 3,\n                    \"reason_type\": 0,\n                    \"player\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist1\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"assist2\": {\n                        \"id\": 0,\n                        \"name\": \"\"\n                    },\n                    \"in_player\": {\n                        \"id\": 1263185,\n                        \"name\": \"阿克塞尔·曼尼·古比约恩森\"\n                    },\n                    \"out_player\": {\n                        \"id\": 62687,\n                        \"name\": \"奥里·斯文·史蒂芬森\"\n                    }\n                }\n            ]\n        },\n        {\n            \"id\": 96650,\n            \"team_id\": 10352,\n            \"first\": 0,\n            \"captain\": 0,\n            \"name\": \"弗罗斯蒂·布林乔尔夫森\",\n            \"logo\": \"https://cdn.sportnanoapi.com/football/player/e5b0d22b1c03d1dae599807ee1112e62.png\",\n            \"national_logo\": \"\",\n            \"shirt_number\": 6,\n            \"position\": \"\",\n            \"x\": 0,\n            \"y\": 0,\n            \"rating\": \"0.0\"\n        }\n    ]\n}\n]")
	//return []byte("[\n    {\n        \"match_id\": 3680783,\n        \"confirmed\": 1,\n        \"home_formation\": \"4-3-9\",\n        \"away_formation\": \"4-4-2\",\n        \"home\": [\n            {\n                \"id\": 1414339,\n                \"team_id\": 14702,\n                \"first\": 1,\n                \"captain\": 0,\n                \"name\": " + i + ",\n                \"logo\": \"\",\n                \"national_logo\": \"\",\n                \"shirt_number\": 16,\n                \"position\": \"\",\n                \"x\": 0,\n                \"y\": 0,\n                \"rating\": \"0.0\",\n                \"incidents\": [\n                    {\n                        \"type\": 1,\n                        \"time\": \"9\",\n                        \"belong\": 1,\n                        \"home_score\": 2,\n                        \"away_score\": 0,\n                        \"reason_type\": 0,\n                        \"player\": {\n                            \"id\": 1414339,\n                            \"name\": \"居琼·比雅尼·布林乔尔夫森\"\n                        },\n                        \"assist1\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist2\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"in_player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"out_player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        }\n                    },\n                    {\n                        \"id\": 1414340,\n                        \"team_id\": 14702,\n                        \"first\": 1,\n                        \"captain\": 0,\n                        \"name\": \"埃尔瓦尔·鲍德温森\",\n                        \"logo\": \"\",\n                        \"national_logo\": \"\",\n                        \"shirt_number\": 18,\n                        \"position\": \"\",\n                        \"x\": 0,\n                        \"y\": 0,\n                        \"rating\": \"0.0\"\n                    },\n                    {\n                        \"type\": 1,\n                        \"time\": \"85\",\n                        \"belong\": 1,\n                        \"home_score\": 5,\n                        \"away_score\": 0,\n                        \"reason_type\": 0,\n                        \"player\": {\n                            \"id\": 1414339,\n                            \"name\": \"居琼·比雅尼·布林乔尔夫森\"\n                        },\n                        \"assist1\": {\n                            \"id\": 1513567,\n                            \"name\": \"西格弗斯·冈纳森·范纳尔\"\n                        },\n                        \"assist2\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"in_player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"out_player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        }\n                    }\n                ]\n            },\n            {\n                \"id\": 1433617,\n                \"team_id\": 14702,\n                \"first\": 1,\n                \"captain\": 0,\n                \"name\": \"克里斯托弗·克里斯蒂安松\",\n                \"logo\": \"\",\n                \"national_logo\": \"\",\n                \"shirt_number\": 15,\n                \"position\": \"\",\n                \"x\": 0,\n                \"y\": 0,\n                \"rating\": \"0.0\"\n            }\n        ],\n        \"away\": [\n            {\n                \"id\": 1146124,\n                \"team_id\": 24328,\n                \"first\": 1,\n                \"captain\": 0,\n                \"name\": \"詹姆斯·戴尔3\",\n                \"logo\": \"https://cdn.sportnanoapi.com/football/player/3bad34d6db39da48c94665c7989d4f7c.png\",\n                \"national_logo\": \"\",\n                \"shirt_number\": 4,\n                \"position\": \"\",\n                \"x\": 0,\n                \"y\": 0,\n                \"rating\": \"0.3\",\n                \"incidents\": [\n                    {\n                        \"type\": 9,\n                        \"time\": \"46\",\n                        \"belong\": 2,\n                        \"home_score\": 3,\n                        \"away_score\": 0,\n                        \"reason_type\": 0,\n                        \"player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist1\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist2\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"in_player\": {\n                            \"id\": 1553739,\n                            \"name\": \"haukur palsson\"\n                        },\n                        \"out_player\": {\n                            \"id\": 1146124,\n                            \"name\": \"詹姆斯·戴尔\"\n                        }\n                    }\n,{\n                        \"type\": 9,\n                        \"time\": \"46\",\n                        \"belong\": 2,\n                        \"home_score\": 3,\n                        \"away_score\": 0,\n                        \"reason_type\": 0,\n                        \"player\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist1\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"assist2\": {\n                            \"id\": 0,\n                            \"name\": \"\"\n                        },\n                        \"in_player\": {\n                            \"id\": 1553739,\n                            \"name\": \"haukur palsson\"\n                        },\n                        \"out_player\": {\n                            \"id\": 1146124,\n                            \"name\": \"詹姆斯·戴尔\"\n                        }\n                    }\n                ]\n            },\n            {\n                \"id\": 1145980,\n                \"team_id\": 24328,\n                \"first\": 1,\n                \"captain\": 0,\n                \"name\": \"什克尔岑·瓦西里1\",\n                \"logo\": \"\",\n                \"national_logo\": \"\",\n                \"shirt_number\": 11,\n                \"position\": \"\",\n                \"x\": 0,\n                \"y\": 0,\n                \"rating\": \"0.5\"\n            }\n        ]\n    }\n]")
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

func statsLiveMock(statsCount int) (msgs []globle.FootballStatsLiveMessage) {
	for i := 1; i <= statsCount; i++ {
		stats := globle.FootballStatsLiveMessage{
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
