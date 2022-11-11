package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x  int64
	wg sync.WaitGroup

	mutex sync.Mutex // 互斥锁

	rwMutex sync.RWMutex // 读写互斥锁
)

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	mutex.Unlock()                    // 解互斥锁
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()                 // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()               // 释放互斥锁
	wg.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()                  // 释放写锁
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.RUnlock()            // 释放读锁
	wg.Done()
}

// do 模拟执行读写操作
// 4个参数，两个函数和两个数字
// wf:写操作的函数 wc:写操作的次数
// rf:读操作的函数 rc:读操作的次数
func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)

}

func main() {
	// 互斥锁模拟10次写操作和1000次读操作
	// do(writeWithLock, readWithLock, 10, 1000)

	// 使用读写互斥锁模拟10次写错做和1000次读操作
	do(writeWithRWLock, readWithRWLock, 10, 1000)
}
