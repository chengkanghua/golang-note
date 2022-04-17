package main

import (
	"context"
	"fmt"
	"time"
)

// context demo

func f1() {
	// context.Context  // 接口类型
	// 得到context的方法
	// 表示一个请求的context的顶层contex（祖先context、带头大哥）
	context.Background() // emptyCtx != nil
	context.TODO()

	// TODO: 等下有接口

	// // WithCancel 得到一个字ctx和取消函数
	// ctx, cancel := context.WithCancel(context.Background())
	// cancel() // 取消

	// WithDeadline
	// now := time.Now()
	// ctx, cancel := context.WithDeadline(context.Background(), now.Add(5*time.Second))
	// defer cancel()
	// 	<-ctx.Done()

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// context.WithValue
	type MyKey string

	// ctx := context.WithValue(context.Background(), MyKey("name"), "杨俊")

	// ctx.Value("name")
	// v, ok := ctx.Value(MyKey("name")).(string)
	// fmt.Println(v, ok)

	// // 其他人存name
	// context.WithValue(context.Background(), "name", "gin")
	// v, ok = ctx.Value("name").(string)
	// fmt.Println(v, ok)

}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				// 外面通知我退出啦，我好开心啊
				return
			case dst <- n:
				n++
				// default:
			}
		}
	}()
	return dst
}

func do() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func main() {
	// do()
	ctx, cancel := context.WithCancel(context.Background())
	go c1(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
