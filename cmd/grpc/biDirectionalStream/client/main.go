package main

import (
	"context"
	"fmt"
	pb "github.com/ronnielin8862/go-api/api/grpc/biDirectionalStreaming"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	// dial server
	conn, err := grpc.Dial("0.0.0.0:50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	client := pb.NewStreamServiceClient(conn)

	stream , err := client.FetchResponse(context.Background())

	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	var id int32 = 11
	var name = "ED 是個..."

	stream.Send(&pb.Request{Id: id, Name: name})
	go func() {
		for i := 0 ; i < 5; i++{
			resp, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("進入EOF")
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			fmt.Println("收到訊息： ", resp)

			id = id + 1
			name = name + "小天才?"

			stream.Send(&pb.Request{Id: id , Name: name})
		}
		fmt.Println("該結束囉")
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second *2 )
		done <- true
	}()
	<-done //we will wait until all response is received
	log.Printf("finished")
}