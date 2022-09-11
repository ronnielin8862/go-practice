package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	user1PB "github.com/ronnielin8862/go-practice/cmd/etcd/step1"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

// ServiceDiscovery 服務發現
type ServiceDiscovery struct {
	cli        *clientv3.Client  //etcd client
	serverList map[string]string //服務列表
	lock       sync.Mutex
}

// NewServiceDiscovery  新建發現服務
func NewServiceDiscovery(endpoints []string) *ServiceDiscovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceDiscovery{
		cli:        cli,
		serverList: make(map[string]string),
	}
}

// WatchService 初始化服務列表和監視
func (s *ServiceDiscovery) WatchService(prefix string) error {
	//根據字首獲取現有的key
	resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	for _, ev := range resp.Kvs {
		s.SetServiceList(string(ev.Key), string(ev.Value))
	}
	conn, err = grpc.Dial(s.serverList["/serviceA"], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	client = user1PB.NewUserServiceClient(conn)

	//監視字首，修改變更的server
	go s.watcher(prefix)
	return nil
}

var conn *grpc.ClientConn
var client user1PB.UserServiceClient

// watcher 監聽字首
func (s *ServiceDiscovery) watcher(prefix string) {

	rch := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Printf("watching prefix:%s now...", prefix)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: //修改或者新增
				if s.serverList[string(ev.Kv.Key)] != string(ev.Kv.Value) {
					s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
					conn, err := grpc.Dial(s.serverList["/serviceA"], grpc.WithInsecure())
					if err != nil {
						log.Fatalf("failed to dial: %v", err)
					}
					client = user1PB.NewUserServiceClient(conn)
				}
			case mvccpb.DELETE: //刪除
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
	defer conn.Close()
}

// SetServiceList 新增服務地址
func (s *ServiceDiscovery) SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.serverList[key] = string(val)
	log.Println("put key :", key, "val:", val)
}

// DelServiceList 刪除服務地址
func (s *ServiceDiscovery) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	log.Println("del key:", key)
}

// GetServices 獲取服務地址
func (s *ServiceDiscovery) GetServices() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]string, 0)

	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

// Close 關閉服務
func (s *ServiceDiscovery) Close() error {
	return s.cli.Close()
}

//go:generate  protoc *.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
func main() {
	var endpoints = []string{"localhost:2379"}
	ser := NewServiceDiscovery(endpoints)
	defer ser.Close()
	ser.WatchService("/serviceA")

	//ser.WatchService("/gRPC/")
	for {
		select {
		case <-time.Tick(10 * time.Second):
			log.Println(ser.GetServices())
			user, err := client.GetUser(context.Background(), &user1PB.User{Id: 1})
			if err != nil {
				fmt.Println("err:", err)
			}
			fmt.Println("user:", user)
		}
	}
}
