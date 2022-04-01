package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("hello")
}

func hello2(i int) {
	defer wg.Done()
	fmt.Println("hello ", i)
}

func f() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func f2() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // 执行到这一句的时候再去访问i的值，此时i是多少不确定
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	// wg.Add(1)
	// go hello() // 创建goroutine是需要资源的

	// for i := 0; i < 100000; i++ {
	// 	wg.Add(1)
	// 	// 启动一个goroutine就登记+1
	// 	go hello2(i)
	// }

	// fmt.Println("你好")
	// // time.Sleep(time.Second)
	// wg.Wait()

	// f()
	f2()
}
