package main

import (
	"fmt"
	"time"
)

// 群里王文建提供的示例

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("等待接收")
		v, ok := <-ch
		fmt.Println(ok)
		fmt.Println("接收到值了", v)
	}()

	time.Sleep(time.Second * 3)
	ch <- 10
	fmt.Println("发送成功")
}
