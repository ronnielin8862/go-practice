package main

import (
	"fmt"
	pb "github.com/ronnielin8862/go-api/api/grpc/biDirectionalStreaming"
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

	for  {
		receive, err := request.Recv()
		if err == io.EOF {
			finishedMsg := fmt.Sprintf("完成囉～～～～，收到 id = %v , name = %v", ids, names)
			fmt.Println(finishedMsg)
			return request.Send(&pb.Response{Result: finishedMsg})
		}

		ids = append(ids, receive.GetId())
		names = append(names, receive.GetName())

		tempMsg := fmt.Sprintf("收到 id = %v , name = %v", receive.GetId(), receive.GetName())
		fmt.Println(tempMsg)
		request.Send(&pb.Response{Result: tempMsg})

		time.Sleep(time.Second *1)
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