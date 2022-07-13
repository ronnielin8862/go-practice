package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/robfig/cron/v3"
	"github.com/ronnielin8862/go-practice/globle"
	"log"
	"os"
	"reflect"
	"time"
)

//CRON_TZ=Asia/Shanghai 0 0 9 * * *
func main() {
	c := cron.New()
	c.AddFunc("CRON_TZ=Asia/Shanghai 40 * * * *", func() {
		fmt.Println("SYNC Expert start...")
	})

	c.AddFunc("CRON_TZ=Asia/Shanghai 25 16 * * *", filteAndDelete)

	c.Start()
	select {}
}

func filteAndDelete() {
	client, err := connectEs()
	if err != nil {
		fmt.Println(err)
		return
	}
	//do, err := client.Get().Index("match_live_text").Id("").Do(context.Background())
	indices, err := client.CatIndices().Index("match_live_text_3687022").Do(context.Background())
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
		//now := time.Now().Unix()
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

		//if now-maxTime > 60*60*24 {
		do, err := client.DeleteIndex(v.Index).Do(context.Background())
		if err != nil {
			fmt.Println("删除失败: ", err)
		} else {
			fmt.Println("删除成功", do)
		}

		//}

	}
}

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
