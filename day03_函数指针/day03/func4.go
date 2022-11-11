package main

import "fmt"

// 匿名函数

// 把匿名函数赋值给变量
func f33() func() {
	f1 := func() {
		fmt.Println("美好的周末就要结束啦~")
	}
	// var f2 func() = func() {
	// 	fmt.Println("美好的周末就要结束啦~")
	// }
	return f1
}

// 匿名函数立即执行
func f34() {
	func() {
		fmt.Println("美好的周末就要结束啦~")
	}()
}
