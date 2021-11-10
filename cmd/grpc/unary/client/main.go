//asshole,fuck
package main

import (
	"context"
	"fmt"
	calculatorPB "github.com/ronnielin8862/go-api/api/grpc/helloWorld"
	"google.golang.org/grpc"
	"log"
)

//測試 godoc，會不會出現呢
//要兩行嗎
func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := calculatorPB.NewCalculatorServiceClient(conn)

	doUnary(client)
}

//測試囉，godoc的狗屁用
//要兩行嗎
func doUnary(client calculatorPB.CalculatorServiceClient) {
	fmt.Println("Staring to do a Unary RPC")
	req := &calculatorPB.CalculatorRequest{
		A: 3,
		B: 19,
	}

	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling CalculatorService: %v \n", err)
	}

	log.Printf("Response from CalculatorService: %v", res.Result)
}
