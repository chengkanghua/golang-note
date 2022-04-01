package main

import (
	"fmt"
	"sync"
	"time"
)

// 七米老师上周被同事问到的问题示例

var wg sync.WaitGroup

type job struct {
	id int64
}

func task() {
	jobs := []job{
		{1},
		{2},
		{3},
		{4},
	}

	for _, item := range jobs {
		// fmt.Println(job.id)
		wg.Add(1)
		go func(item job) {
			defer wg.Done()
			time.Sleep(time.Millisecond) // 模拟耗时的操作
			fmt.Println(item.id)  // 4 4 4 4
		}(item)
	}
}

func main() {
	task()
	wg.Wait()
}
