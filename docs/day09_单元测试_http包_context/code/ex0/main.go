package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// “我就不”的例子...
// 生成随机数然后多个goroutine消费随机数，计算随机数乘机的程序

var wg sync.WaitGroup
var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

type Job struct {
	value int64
}

type Result struct {
	sum int64
	job *Job
}

func main() {
	go producer(jobChan) // 生产者

	// 开启20个goroutine去消费数据
	for i := 0; i < 20; i++ {
		go consumer(jobChan, resultChan)
	}

	// 从结果的通道里接收值并打印
	for v := range resultChan {
		fmt.Printf("sum:%v job:%#v\n", v.sum, v.job)
	}
}

// producer 生成job 发送到jonChan中
func producer(jobChan chan<- *Job) {
	defer close(jobChan) // 执行不到

	for { // 没有退出条件
		newJob := &Job{
			value: int64(rand.Intn(10)),
		}
		jobChan <- newJob
	}
}

// consumer 从jobChan接收job,然后执行计算，最后把结果发送到resultChan
func consumer(jobChan <-chan *Job, resultChan chan<- *Result) {
	defer close(resultChan) // 执行不到

	for { // 没有退出条件
		// sum := 0 // int
		sum := int64(0) // var sum int64
		job := <-jobChan
		sum = job.value * job.value // 两数相乘可能会溢出  int8*int8  127*127
		res := &Result{
			sum: int64(sum),
			job: job,
		}
		resultChan <- res
	}
}
