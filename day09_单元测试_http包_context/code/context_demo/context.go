package main

import (
	"context"
	"fmt"
	"time"
)

// 多层goroutine控制

// 1->2->3
// main函数里面通知3 goroutine退出

func c1(ctx context.Context) {
	ctx2, cancel := context.WithTimeout(ctx, time.Hour) // 串成一串
	defer cancel()
	go c2(ctx2)
	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println("c1")
		case <-ctx.Done():
			fmt.Println("c1收到main函数的退出信号，但是我不退出")
			time.Sleep(time.Second)
			return
		}
	}
}

func c2(ctx context.Context) {
	go c3(ctx)
	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println("c2")
		case <-ctx.Done():
			fmt.Println("c2收到main函数的退出信号，但是我不退出")
			time.Sleep(time.Second)
		}
	}
}

func c3(ctx context.Context) {
	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println("c3")
		case <-ctx.Done():
			fmt.Println("c3要退出啦，它好开心啊")
			return
		}
	}
}
