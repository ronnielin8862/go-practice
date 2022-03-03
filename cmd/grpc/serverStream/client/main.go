package main

import (
	"context"
	"fmt"
	pb "github.com/ronnielin8862/go-practice/api/grpc/serverStream"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	// dial server
	conn, err := grpc.Dial("0.0.0.0:50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	client := pb.NewStreamServiceClient(conn)
	in := &pb.Request{Id: 55}

	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("進入EOF")
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Result)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}
