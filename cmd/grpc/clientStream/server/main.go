package main

import (
	"fmt"
	pb "github.com/ronnielin8862/go-practice/api/grpc/clientStram"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct{}

func (*Server) FetchResponse(request pb.StreamService_FetchResponseServer) error {

	//聲明指定容量的slice
	//slice := make([]int32 ,0 , 5)
	//聲明不指定容量的slice
	var slice []int32

	for {
		id, err := request.Recv()

		if err == io.EOF {
			//字串處理：可以正常列印，但無法存儲
			fmt.Printf("收完了。開始列印。 收到的id = %v\n", slice)
			//無法正常處理[]int32 效果: 1 完成囉～～～～，收到你以下id
			response := "完成囉～～～～，收到你以下id " + string(slice)
			fmt.Println("1", response)
			//最佳處理結果： 2 完成囉～～～～，收到你以下id [10 11 12 13 14 15 16]
			temp := fmt.Sprintf("2 完成囉～～～～，收到你以下id %v", slice)
			fmt.Println(temp)
			return request.SendAndClose(&pb.Response{Result: temp})
		}

		slice = append(slice, id.GetId())

		fmt.Println("收到 id = ", id)
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
