package main

import "fmt"

// 切片共享底层数据

func slice5() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(b) // [1 2 3]
	b[1] = 200     // 改的底层数组里的值
	fmt.Println(b) // [1 200 3]
	fmt.Println(a) // [1 200 3]
}

func slice6() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(b) // [1 2 3]
	a[1] = 200     // 改的底层数组里的值
	fmt.Println(a) // [1 200 3]
	fmt.Println(b) // [1 200 3]
}

func copyDemo() {
	a := []int{1, 2, 3}
	// var b = make([]int, 0, len(a))
	b := make([]int, 0)
	copy(b, a)     // 把切片a中的值拷贝到切片b中
	fmt.Println(b) // ? 为什么？
}

func copyDemo2() {
	a := []int{1, 2, 3}
	// var b = make([]int, 0, len(a))
	b := make([]int, len(a)) // 直接按目标切片的长度进行初始化
	copy(b, a)               // 把切片a中的值拷贝到切片b中
	fmt.Println(b)           // ?
	b[1] = 200

	fmt.Println(a) // ?
	fmt.Println(b) // [1 200 3]
}
