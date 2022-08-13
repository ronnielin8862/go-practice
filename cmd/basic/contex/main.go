package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "keyA", "valueA")
	//ctx2, _ := context.WithDeadline(ctx, time.Now().Add(time.Second*2)) // 设置超时时间點
	ctx2, _ := context.WithTimeout(ctx, time.Second*5) // 设置超時倒數時間
	go layer1(ctx2)
	select {
	case <-ctx2.Done():
		fmt.Println("main done")
	}
	time.Sleep(time.Second * 3)
}

func layer1(ctx context.Context) {
	fmt.Println("layer1")
	layer2(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("layer1 done")
	}
}

func layer2(ctx context.Context) {
	fmt.Println("layer2")
	layer3(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("layer2 done")
	}
}

func layer3(ctx context.Context) {
	fmt.Println("layer3")
	fmt.Println("value = ", ctx.Value("keyA"))
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Println("deadline = ", deadline)
	}
	layer4(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("layer3 done")
	}
}

func layer4(ctx context.Context) {
	time.Sleep(time.Second * 1)
	err := ctx.Err()
	if err != nil {
		fmt.Println("layer4 err = ", err)
	} else {
		fmt.Println("layer4 err = nil")
	}
}
