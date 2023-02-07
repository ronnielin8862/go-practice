package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ronnielin8862/go-practice/pkg/consul"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	consulAddr := flag.String("consul.addr", "0.0.0.0", "consul address")     //consul addr
	consulPort := flag.Int("consul.port", 8500, "consul port")                //consul端口
	serviceName := flag.String("service.name", "go-server-1", "service name") //服务名称

	flag.Parse()
	ctx, _ := context.WithCancel(context.Background())

	client := consul.NewDiscoveryClient(*consulAddr, *consulPort) //获取consul客户端，用于服务注册和发现

	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	//定时刷新示例
	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			instance, err := client.DiscoverServices(ctx, *serviceName)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("get instance num", len(instance))
			for _, v := range instance {
				fmt.Printf("instanceID:%s,address:%s,port:%d\n", v.ID, v.Address, v.Port)
			}

		case <-c:
			log.Println("discovery service exit!")
		}
	}

}
