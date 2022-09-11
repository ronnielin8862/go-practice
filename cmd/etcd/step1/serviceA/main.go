package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	user1PB "github.com/ronnielin8862/go-practice/cmd/etcd/step1"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

// ServiceRegister 建立租約註冊服務
type ServiceRegister struct {
	cli     *clientv3.Client //etcd client
	leaseID clientv3.LeaseID //租約ID
	//租約keepalieve相應chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string //key
	val           string //value
}

// NewServiceRegister 新建註冊服務
func NewServiceRegister(endpoints []string, key, val string, lease int64) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	ser := &ServiceRegister{
		cli: cli,
		key: key,
		val: val,
	}

	//申請租約設定時間keepalive
	if err = ser.putKeyWithLease(lease); err != nil {
		return nil, err
	}

	return ser, nil
}

// 設定租約
func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	//設定租約時間
	resp, err := s.cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	//註冊服務並繫結租約
	_, err = s.cli.Put(context.Background(), s.key, s.val, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	//設定續租 定期傳送需求請求
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), resp.ID)

	if err != nil {
		return err
	}
	s.leaseID = resp.ID
	log.Println(s.leaseID)
	s.keepAliveChan = leaseRespChan
	log.Printf("Put key:%s  val:%s  success!", s.key, s.val)
	return nil
}

// ListenLeaseRespChan 監聽 續租情況
func (s *ServiceRegister) ListenLeaseRespChan() {
	for leaseKeepResp := range s.keepAliveChan {
		log.Println("續約成功", leaseKeepResp)
	}
	log.Println("關閉續租")
}

// Close 登出服務
func (s *ServiceRegister) Close() error {
	//撤銷租約
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	log.Println("撤銷租約")
	return s.cli.Close()
}

type GetUserStruct struct{}

func main() {
	host := "localhost:50051"
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}
	grpcServer := grpc.NewServer()
	user1PB.RegisterUserServiceServer(grpcServer, &GetUserStruct{})

	var endpoints = []string{"localhost:2379"}
	ser, err := NewServiceRegister(endpoints, "/serviceA", host, 60)
	if err != nil {
		log.Fatalln(err)
	}
	//監聽續租相應chan
	go ser.ListenLeaseRespChan()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

func (u *GetUserStruct) GetUser(ctx context.Context, req *user1PB.User) (*user1PB.User, error) {
	fmt.Printf("接到user request: %+v \n", req)
	req.Name = "ronnie"
	req.Age = 18
	return req, nil
}
