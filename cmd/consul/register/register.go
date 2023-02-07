package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/google/uuid"
	"github.com/ronnielin8862/go-practice/pkg/consul"
)

func main() {
	consulAddr := flag.String("consul.addr", "0.0.0.0", "consul address")     //consul addr
	consulPort := flag.Int("consul.port", 8500, "consul port")                //consul端口
	servicePort := flag.Int("service.port", 12310, "service port")            //服务端口
	serviceName := flag.String("service.name", "go-server-1", "service name") //服务名称
	serviceAddr := flag.String("service.addr", "127.0.0.1", "service addr")   //服务地址

	flag.Parse()
	ctx, _ := context.WithCancel(context.Background())

	instanceId := *serviceName + "-" + strings.Replace(uuid.New().String(), "-", "", -1)

	client := consul.NewDiscoveryClient(*consulAddr, *consulPort) //获取consul客户端，用于服务注册和发现

	//将服务注册到consul中
	err := client.Register(ctx, *serviceName, instanceId, "/health", *serviceAddr, *servicePort, nil, nil)

	if err != nil {
		log.Fatal(err)
	}

	//开启http服务，并注册handle
	go func() {
		http.HandleFunc("/health", checkHealth)
		http.ListenAndServe(fmt.Sprintf("%s:%d", *serviceAddr, *servicePort), nil)
	}()

	// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()
	log.Printf("exit %s", <-c)
	client.Deregister(ctx, instanceId)
	log.Printf("Deregister service %s", instanceId)

}

func checkHealth(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "SUCCESS")
}
