package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"sync"
	"time"
)

/*
使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
	开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
	开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	主 goroutine 从resultChan取出结果并打印到终端输出
*/

// CreateNum 生成一个int64的随机数
func CreateNum() (num int64) {
	maxNum := int64(math.MaxInt64)                         // 获取int64的最大值
	result, _ := rand.Int(rand.Reader, big.NewInt(maxNum)) // 使用crypto/rand生成真随机数
	num = result.Int64()                                   // 把big.int类型转换为int64类型
	return
}

// ComputeNum 计算num各个位数的和
func ComputeNum(num int64) (sumNum int) {
	numStr := strconv.Itoa(int(num)) // 把int类型转换为string
	numStrSlice := []rune(numStr)    // 把字符串转为rune类型

	// 计算num各个位数的和
	for i := 0; i < len(numStrSlice); i++ {
		numInt, _ := strconv.Atoi(string(numStrSlice[i]))
		sumNum += numInt
	}
	return
}

// run 运行函数
func run(amount int) {
	wg := sync.WaitGroup{}

	jobChan := make(chan int64, 8)  // 创建并初始化一个int64的channel,用于存放随机数
	resultChan := make(chan int, 8) // 创建并初始化一个int的channel,用于接收结果

	// 开启1个goroutine执行 (生成 amount个 随机数并发送到jobChan)
	wg.Add(1)
	go func() {
		defer close(jobChan)
		for i := 0; i < amount; i++ {
			jobChan <- CreateNum() // 第9次阻塞....
		}
		wg.Done()
	}()

	// 开启24个goroutine执行(计算值并发送到resultChan)
	go func() {
		for i := 0; i < 24; i++ {
			wg.Add(1)

			go func() {
				for i := range jobChan { // 退出不了
					resultChan <- ComputeNum(i) // 接收jobChan的值并计算各个位数的和
				}
				wg.Done()
			}()
		}

		wg.Wait()         // 等待goroutine执行完毕
		close(resultChan) // 关闭resultChan
	}()

	// 接收resultChan的值并打印
	i := 0
	for v := range resultChan {
		fmt.Println(i, "值:", v)
		i++
	}
}

func main() {
	start := time.Now() // 获取当前时间
	run(5)
	fmt.Println(time.Since(start)) // 计算程序耗时
}
