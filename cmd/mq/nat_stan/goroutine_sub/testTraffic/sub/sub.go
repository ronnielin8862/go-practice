package sub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/spf13/cast"
	"time"
)

var stanS stan.Conn
var rs *redis.Client

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

	rs = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	err = rs.Ping().Err()
	if err != nil {
		fmt.Println("connect to redis failed, ", err)
	}
}

func main() {
	subscribe2, err := stanS.Subscribe("test_2", printSub2, stan.DurableName("test_2"), stan.AckWait(111*time.Second)) //, stan.SetManualAckMode()
	if err != nil {
		fmt.Println("error by subscribe: ", err)
	}
	defer subscribe2.Close()
}

func printSub2(msg *stan.Msg) {
	go func() {
		fmt.Println("Sub2 Received a message: ", string(msg.Data))
		//err := msg.Ack()
		//if err != nil {
		//	fmt.Println("error by ack: ", err)
		//}
		time.Sleep(1 * time.Millisecond)
		//time.Sleep(3 * time.Second)
	}()
}

func HandlerFootballLineup(msg *stan.Msg) {
	var raw Lineup
	err := json.Unmarshal(msg.Data, &raw)
	if err != nil {
		fmt.Errorf("Unmarshal Lineup err [%v], raw data [%v]", err, string(msg.Data))
		return
	}

	// 資料比對
	rdsKey := fmt.Sprint(MatchLiveConst, ":", LineupConst, ":", fmt.Sprintf("%v", raw.MatchId))
	newData, isNew := footballLineupGetNewData(raw, rdsKey)
	if !isNew {
		return
	}

	match_handler.CreateRoom(raw.MatchId)

	rawMarshal, _ := json.Marshal(newData)
	fmt.Println("receive : New footballLineup , room_id : %v, content: %v", raw.MatchId, string(rawMarshal))

	newData.MatchId = raw.MatchId
	newData.Confirmed = raw.Confirmed
	newData.HomeFormation = raw.HomeFormation
	newData.AwayFormation = raw.AwayFormation

	newMsg := RoomLineupMessage{
		SentMessageStruct: SentMessageStruct{Type: RoomLineupMessageType},
		Lineup:            newData,
	}
	// 存入redis
	rawMsg := RoomLineupMessage{
		SentMessageStruct: SentMessageStruct{Type: RoomLineupMessageType},
		Lineup:            raw,
	}
	n, _ := json.Marshal(rawMsg)
	err = server.GetRedis().DB.Set(context.Background(), rdsKey, string(n), 24*time.Hour).Err()
	if err != nil {
		fmt.Errorf("footballLineup Save rds err : %v, room_id : %v", err, raw.MatchId)
	}
	// 廣播 部分更新的資料
	m, _ := json.Marshal(newMsg)
	imsdk.GetSdk().Broadcast(cast.ToString(raw.MatchId), []string{string(m)})

}

func footballLineupGetNewData(raw Lineup, rdsKey string) (newData Lineup, isNew bool) {
	result, err := server.GetRedis().DB.Get(context.Background(), rdsKey).Result()
	// 如果rds沒有資料，則新增
	if err != nil {
		fmt.Println("Get redis data err : %v , room_id : %v", err, raw.MatchId)
		return raw, true
	}

	// 無法解析rds舊資料時，用新資料覆蓋
	var rdsMsg RoomLineupMessage
	err = json.Unmarshal([]byte(result), &rdsMsg)
	if err != nil {
		fmt.Println("Unmarshal redis data err [%v] , room_id : %v", err, raw.MatchId)
		return raw, true
	}

	rdsData := rdsMsg.Lineup

	// match header 有變更時，重新處理整筆資料
	if raw.Confirmed != rdsData.Confirmed || raw.HomeFormation != rdsData.HomeFormation || raw.AwayFormation != rdsData.AwayFormation {
		return raw, true
	}

	h := make(map[int]int)
	for k, v := range rdsData.Home {
		h[v.LineupId] = k
	}

	for _, rawItem := range raw.Home {
		// 比對lineItem 不同的資料加入newData處理
		if k, ok := h[rawItem.LineupId]; !ok {
			fmt.Println("LineupId not found in redis, new lineupId : ", rawItem.LineupId)
			return raw, true
		} else {
			rdsItem := rdsData.Home[k]
			if rawItem.TeamId != rdsItem.TeamId || rawItem.First != rdsItem.First || rawItem.Captain != rdsItem.Captain ||
				rawItem.Name != rdsItem.Name || rawItem.Logo != rdsItem.Logo || rawItem.NationalLogo != rdsItem.NationalLogo ||
				rawItem.ShirtNumber != rdsItem.ShirtNumber || rawItem.Position != rdsItem.Position || rawItem.X != rdsItem.X ||
				rawItem.Y != rdsItem.Y || rawItem.Rating != rdsItem.Rating {
				newData.Home = append(newData.Home, rawItem)
				isNew = true
				continue
			}
			// 比對incidents, 有新增事件則加入 ,往下處理
			if len(rdsItem.Incidents) != len(rawItem.Incidents) {
				newData.Home = append(newData.Home, rawItem)
				isNew = true
				continue
			}
		}
	}

	for k, v := range rdsData.Away {
		h[v.LineupId] = k
	}

	for _, rawItem := range raw.Away {
		// 比對lineItem 不同的資料加入 newData
		if k, ok := h[rawItem.LineupId]; !ok {
			fmt.Println("LineupId not found in redis, new lineupId : ", rawItem.LineupId)
			return raw, true
		} else {
			rdsItem := rdsData.Away[k]
			if rawItem.TeamId != rdsItem.TeamId || rawItem.First != rdsItem.First || rawItem.Captain != rdsItem.Captain ||
				rawItem.Name != rdsItem.Name || rawItem.Logo != rdsItem.Logo || rawItem.NationalLogo != rdsItem.NationalLogo ||
				rawItem.ShirtNumber != rdsItem.ShirtNumber || rawItem.Position != rdsItem.Position || rawItem.X != rdsItem.X ||
				rawItem.Y != rdsItem.Y || rawItem.Rating != rdsItem.Rating {
				newData.Away = append(newData.Away, rawItem)
				isNew = true
				continue
			}
			if len(rdsItem.Incidents) != len(rawItem.Incidents) {
				newData.Away = append(newData.Away, rawItem)
				isNew = true
				continue
			}
		}
	}
	return newData, isNew
}
