package main

import (
	"context"
	pb "github.com/ronnielin8862/go-practice/api/grpc/clientStram"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("0.0.0.0:50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	client := pb.NewStreamServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.FetchResponse(ctx)

	for i := 10; i < 17; i++ {
		if err := stream.Send(&pb.Request{Id: int32(i)}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, i, err)
		}
	}
	// 告知 server 傳送完畢，並準備接收 server 的回應
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("reply = : %v", reply)

}
