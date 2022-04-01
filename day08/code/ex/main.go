package main

import (
	"fmt"
	"time"
)

func f1(ch1 chan int) {
	// time.Sleep(10 * time.Second)
	ch1 <- 10 // 没有缓冲区，没有人接收，代码执行不下去，一直阻塞在这里
	ch1 <- 100
	ch1 <- 200
	fmt.Println("王文建")
}

func main() {
	ch1 := make(chan int) // 无缓冲的channel
	go f1(ch1)
	v := <-ch1 // 在main goroutine有人接收值
	fmt.Println(v)
	v2 := <-ch1
	fmt.Println(v2)


	t := time.Tick(10*time.Second)
	<- t  // 10秒钟之后就能取到值了
}
