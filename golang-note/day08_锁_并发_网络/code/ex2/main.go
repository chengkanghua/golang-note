package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 上周作业
/*
使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
主 goroutine 从resultChan取出结果并打印到终端输出
*/

func f1() {
	rand.Seed(time.Now().Unix()) // 设置随机因子
	v := rand.Int63()
	fmt.Println(v)
}

// 两个channel
// 两个任务：生产随机数的；计算和的

// result 表示结果的结构体
type result struct {
	num int64 // 随机数
	sum int64 // 计算得到的和
}

// genNumber 生成随机数int64
func genNumber() <-chan int64 {
	var jobChan = make(chan int64, 24)
	rand.Seed(time.Now().Unix()) // 设置随机因子
	// 在后台一直执行
	go func() {
		for {
			// v := rand.Int63()
			// v := rand.Intn(9999) // int
			v := rand.Int63()
			select {
			case jobChan <- int64(v):
			default:
				time.Sleep(time.Millisecond)
			}
		}
	}()
	return jobChan
}

// sum 从jobChan取出int64的值计算每位数字的和 发送到resultChan
func sum(ch <-chan int64, resChan chan result) {
	// 循环的从ch取数字，然后计算和
	for v := range ch {
		// v := <-ch // 6341
		r := result{
			num: v, // 原始数字记录
		}
		var res int64 = 0
		// 计算每个位置上的数字之和
		for v > 0 {
			res += v % 10
			v = v / 10
		}

		r.sum = res // 把计算的结果记录

		resChan <- r
	}

	// 6341 -> "6341" -> []int8 -> 求和

	// 6341 % 10 = 1
	// 6341 / 10 = 634
	// 634 % 10 = 4
	// 634 / 10 = 63
	// 63 % 10 = 3
	// 63 / 10 = 6
	// 6 % 10 = 6

}

func main() {
	// f1()

	jobChan := genNumber()
	resChan := make(chan result, 10)

	// 开启24个goroutine去干活
	for i := 0; i < 24; i++ {
		// go func() {
		// 	resChan := sum(jobChan)
		// }()
		go sum(jobChan, resChan)
	}

	// 从resultChan里接收值，打印结果
	// for res := range resChan {
	for {
		res := <-resChan
		time.Sleep(time.Second)
		fmt.Printf("数字:%d 和:%d\n", res.num, res.sum)
	}
}
