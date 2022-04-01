package main

import "fmt"

func main() {
	// ch := make(chan int, 1)
	// go func() {
	// 	fmt.Println("等接收值....")
	// 	x := <-ch // 等着从通道中接收值
	// 	fmt.Println("接收到值了", x)
	// }()

	// time.Sleep(3 * time.Second)
	// ch <- 10
	// // ch <- 20
	// close(ch)
	// // ch <- 10
	// fmt.Println("发送成功")

	f3()
}

func f2(ch chan int) {
	// for {
	// 	// v := <-ch
	// 	// fmt.Println(v) // 接收完1和2之后，会一直接收到0
	// 	v, ok := <-ch
	// 	if !ok {
	// 		fmt.Println("通道已关闭")
	// 		break
	// 	}
	// 	fmt.Printf("v:%#v ok:%#v\n", v, ok)
	// }

	// 使用for range循环执行接收操作
	// 直到通道被关闭且数据被接收完再退出循环
	for v := range ch {
		fmt.Println(v)
	}
}

func f3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("元素数量：", len(ch))
	close(ch)
	f2(ch)
}
