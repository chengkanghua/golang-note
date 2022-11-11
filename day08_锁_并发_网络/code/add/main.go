package main

import (
	"fmt"
	"sync"
)

var (
	x int64 // 因数据竞争导致并发不安全

	wg sync.WaitGroup // 等待组

	mutex sync.Mutex
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		mutex.Lock() // 加锁
		x = x + 1
		mutex.Unlock() // 解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x) // 10000
}
