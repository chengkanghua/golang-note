package main

import (
	"context"
	"fmt"
	"sync"

	"time"
)

type mystring string

type CtxKey int8

const (
	CtxName CtxKey = iota
	CtxAge  CtxKey = iota
	// ...
)

func f1() {
	ctx := context.WithValue(context.Background(), CtxName, "杨俊")

	// 根据key取值
	value := ctx.Value(CtxName).(string) // 对取出的值做类型断言
	fmt.Println(value)
}

func f2() {
	// 一个请求来请求用户个人中心
	// 总体请求时间是300ms超时
	// 分两个goroutine去查询 用户个人信息和用户订单信息
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	// 查用户个人信息
	go func(ctx context.Context) {
		// 100ms？
		// 基于父 ctx造一个子ctx
		ctx2, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()
	}(ctx)

	// 查用户订单信息
	go func(ctx context.Context) {
		ctx3, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()
		// ...
	}(ctx)
}

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		// ch := ctx.Done()
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
