package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/ronnielin8862/go-practice/globle"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

const mapping = `
{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
      }
    }
  }
}`

const (
	AttentAnchorChannel = "attentAnchorChannel"
	ChangePrpos         = "changePropChannel"
	FootballTextLive    = "football_text_live"
)

func connectEs() (*elastic.Client, error) {
	return elastic.NewClient(
		// 设置Elastic服务地址
		//elastic.SetURL("https://es-digpfxq8.public.tencentelasticsearch.com:9200"),
		elastic.SetURL("http://127.0.0.1:9200"),
		// 連線帳密
		elastic.SetBasicAuth("elastic", "Ddu@2022!"),
		// 是否转换请求地址，默认为true,当等于true时 请求http://ip:port/_nodes/http，将其返回的url作为请求路径
		elastic.SetSniff(false),
		// 心跳检查,间隔时间
		elastic.SetHealthcheckInterval(time.Second*5),
		// 设置错误日志
		elastic.SetErrorLog(log.New(os.Stderr, "ES-ERROR ", log.LstdFlags)),
		// 设置info日志
		elastic.SetInfoLog(log.New(os.Stdout, "ES-INFO ", log.LstdFlags)),
	)
}

type UserInfo struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Birth string `json:"birth"`
}

func main() {
	client, err := connectEs()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 健康检查
	do, _ := client.ClusterHealth().Index().Do(context.TODO())
	fmt.Println("健康检查:", do)
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	//add(client)
	//findMany(client)
	//count(client)
	//filteAndDelete(client)
	//findByQuery(client)
	//deleteByQuery(client)
	CleanChatHistory(client)
}

func deleteByQuery(client *elastic.Client) {
	bq := elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("age").Lt(19))
	res, err := client.DeleteByQuery("go-test").Query(bq).Do(context.Background())
	if err != nil {
		fmt.Println("删除失败:%s", err)
		return
	}
	fmt.Println("删除成功", res.Total)
}

func CleanChatHistory(client *elastic.Client) {
	//t := time.Now().Unix() - 60*60*24
	q := elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("create_time").Lt(1))
	do, err := client.DeleteByQuery("es_chat_history").Query(q).Do(context.Background())
	if err != nil {
		fmt.Println("删除失败:%s", err)
	} else {
		fmt.Println("删除成功", do.Total)
	}
}

func findByQuery(client *elastic.Client) {
	bq := elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("create_time").Gt(1657959002))
	do, err := client.Search("match_live_text_3666726").Query(bq).Size(100).Do(context.Background())
	//  todo 測試刪除

	if err != nil {
		fmt.Println("查询失败: ", err)
	} else {
		//fmt.Println("查询成功: ")
		//for _, hit := range do.Hits.Hits {
		//	fmt.Println(hit.Id, hit.Index, string(hit.Source))
		//}
		fmt.Println("len do = ", len(do.Hits.Hits))
	}
}

func filteAndDelete(client *elastic.Client) {
	//do, err := client.Get().Index("match_live_text").Id("").Do(context.Background())
	indices, err := client.CatIndices().Index("match_live_text_*").Do(context.Background())
	if err != nil {
		fmt.Println("查询失败: ", err)
	} else {
		fmt.Println("查询成功", indices)
	}
	for _, v := range indices {
		result, err := client.Search(v.Index).Do(context.Background())
		if err != nil {
			fmt.Println("查询失败:", err)
		} else {
			//fmt.Printf("查询成功 , id : %v , content : %v ", v.Index, result)
		}
		now := time.Now().Unix()
		var text globle.TextLiveStruct

		var maxTime int64
		for _, j := range result.Each(reflect.TypeOf(text)) {
			rowTime := j.(globle.TextLiveStruct).CreateTime
			if err != nil {
				fmt.Println("转换失败:", err)
			}
			if rowTime > maxTime {
				maxTime = rowTime
			}
		}

		if now-maxTime > 60*60*24 {
			do, err := client.DeleteIndex(v.Index).Do(context.Background())
			if err != nil {
				fmt.Println("删除失败: ", err)
			} else {
				fmt.Println("删除成功", do)
			}

			//fmt.Println("now = ", now, "   ,maxtime : ", maxTime, "   ,now-maxTime : ", now-maxTime)
		}

		//break
	}
}

func count(client *elastic.Client) {
	bq := elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("create_time").Gt(1657959102))
	res, err := client.Count("match_live_text_3666726").Query(bq).Do(context.Background())
	if err != nil {
		fmt.Println("查询失败:%s", err)
		return
	}
	fmt.Println("查询成功", res)
}

func findMany(client *elastic.Client) {

	// 创建查询语句
	//query := elastic.NewMatchAllQuery()
	// 查询®
	//sourceContext := elastic.FetchSourceContext{}

	// 分頁搜索
	sort := elastic.NewFieldSort("create_time").Order(false)
	res, err := client.Search("match_live_text_3666726").Size(3).SortBy(sort).Do(context.Background())

	if err != nil {
		fmt.Println("查询失败:%s", err)
		return
	}
	// 打印查询结果
	if res.TotalHits() > 0 {
		var b1 RoomTextLive
		items := res.Each(reflect.TypeOf(b1))

		fmt.Println("len = ", len(items))

		for _, item := range res.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			if t, ok := item.(RoomTextLive); ok {
				fmt.Println("data : ", t.Data)
			} else {
				fmt.Println("err : ", ok)
			}
		}
	}
}

type RoomTextLive struct {
	Id        int64  `json:"match_id"`   // 赛事id
	Time      string `json:"time"`       // 事件时间
	EventType int    `json:"event_type"` // 事件类型
	Data      string `json:"data"`       // 事件文本
	Position  int8   `json:"position"`   // 事件發生方， 0-中立 1-主队 2-客队
	Main      int8   `json:"main"`       // 是否重要事件 0-否 1-是
}

type RoomTextLiveSlice struct {
	SentMessageStruct
	TextLive []*RoomTextLive `json:"text_live"`
}

type RoomTextLiveMessage struct {
	SentMessageStruct
	Id         int64  `json:"match_id"`    // 赛事id
	Time       string `json:"time"`        // 事件时间
	Type       int8   `json:"type"`        // 事件类型
	Data       string `json:"data"`        // 事件文本
	Position   int8   `json:"position"`    // 事件發生方， 0-中立 1-主队 2-客队
	Main       int8   `json:"main"`        // 是否重要事件 0-否 1-是
	CreateTime string `json:"create_time"` // 創建時間
}
type SentMessageStruct struct {
	Type    string `json:"type"`
	Message string `default:"" json:"message,omitempty"`
}

func add(client *elastic.Client) {

	//client, _ := connectEs()
	ctx := context.Background()
	for i := 0; i <= 70; i++ {
		// 创建userInfo
		userInfo := UserInfo{
			Name:  "张三！",
			Age:   18,
			Birth: "1991-03-04",
		}
		_, err := client.Index().Index("go-test").BodyJson(userInfo).Do(ctx)
		if err != nil {

			if err != nil {
				fmt.Println("添加失败:%s", err)
			} else {
				fmt.Println("添加成功", i)
			}
		}
	}
}

func addMany(client *elastic.Client) {
	ctx := context.Background()
	// 创建用户
	userNames := map[string]string{
		"李四": "1992-04-25",
		"张亮": "1994-07-15",
		"小明": "1991-12-03",
	}
	rand.Seed(time.Now().Unix())
	// 创建bulk
	userBulk := client.Bulk().Index("go-test")
	id := 4
	for n, b := range userNames {
		userTmp := UserInfo{Name: n, Age: rand.Intn(50), Birth: b}
		// 批量添加到bulk
		doc := elastic.NewBulkIndexRequest().Id(strconv.Itoa(id)).Doc(userTmp)
		userBulk.Add(doc)
		id++
	}
	// 检查被添加数据是否为空
	if userBulk.NumberOfActions() < 1 {
		fmt.Println("被添加的数据不能为空！")
		return
	}
	// 保存
	res, err := userBulk.Do(ctx)
	if err != nil {
		fmt.Println("保存失败:%s", err)
		return
	}
	fmt.Println("保存成功: ", res)
}
