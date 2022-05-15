package main

import (
	"fmt"
	"time"
)

func recv(c chan int) {
	// ret := <-c
	// fmt.Println("接收成功", ret)
	/* 	for {
		v, ok := <-c
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("接收到值 v:%#v ok: %#v \n", v, ok)
	} */
	for v := range c {
		fmt.Printf("接收到值 v:%#v", v) // 这里为什么接收不到20，上面for循环却可以
	}
}

func main() {
	var ch1 chan int
	// var ch2 chan bool
	// var ch3 chan []int
	fmt.Println(ch1) //<nil>

	ch2 := make(chan int)
	// ch3 := make(chan bool, 1) // // 声明一个缓冲区大小为1的通道
	go recv(ch2)
	ch2 <- 10 //把10发送到ch中 形成死锁，等待接收方才能发送成
	ch2 <- 20
	fmt.Println("发送成功")
	time.Sleep(time.Second)
	
}
