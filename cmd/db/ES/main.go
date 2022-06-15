package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"math/rand"
	"os"
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

func connectEs() (*elastic.Client, error) {
	return elastic.NewClient(
		// 设置Elastic服务地址
		elastic.SetURL("http://52.221.194.38:9200"),
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
	addOne(client)
}

func addOne(client *elastic.Client) {

	//client, _ := connectEs()
	ctx := context.Background()
	// 创建userInfo
	userInfo := UserInfo{
		Name:  "张三！",
		Age:   18,
		Birth: "1991-03-04",
	}
	res, err := client.Index().Index("go-test").Id("1").BodyJson(userInfo).Do(ctx)
	if err != nil {
		fmt.Println("添加失败:%s", err)
	}
	fmt.Println("添加成功", res)
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
