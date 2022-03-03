package main

import (
	"errors"
	"fmt"
	pb "github.com/ronnielin8862/go-practice/api/grpc/biDirectionalStreaming"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type Server struct{}

func (s *Server) FetchResponse(request pb.StreamService_FetchResponseServer) error {

	var ids []int32
	var names []string
	var nilCount int
	for {
		receive, err := request.Recv()
		if err == io.EOF {
			finishedMsg := fmt.Sprintf("完成囉～～～～，收到 id = %v , name = %v", ids, names)
			fmt.Println(finishedMsg)
			return request.Send(&pb.Response{Result: finishedMsg})
		}

		fmt.Println("receive = ", receive)
		//針對客戶端斷訊狀況檢測
		if receive == nil {
			nilCount++
			fmt.Println("nilCount + 1 = ", nilCount)
			if nilCount == 3 {
				return errors.New("收到nil request 多次，疑似客戶端斷訊，系統結束監聽，請重新連接")
			}
			time.Sleep(time.Second * 1)
			continue
		} else {
			nilCount = 0
		}

		ids = append(ids, receive.GetId())
		names = append(names, receive.GetName())

		tempMsg := fmt.Sprintf("收到 id = %v , name = %v", receive.GetId(), receive.GetName())
		request.Send(&pb.Response{Result: tempMsg})

		time.Sleep(time.Second * 1)

		//測試server端監測到某種不正常狀態，主動中斷對話  (測試成功，可以讓客戶端也同時終止連線)
		//return errors.New("測試終止對話")
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &Server{})

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
