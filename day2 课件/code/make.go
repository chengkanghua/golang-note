package main

import "fmt"

// 字面量初始化(花括号)
func slice4() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)  // [1 2 3]
	s2 := []int{99: 1}
	fmt.Println(s2)
}

// 使用make函数初始化切片
func make1() {
	var s []int
	fmt.Println(s == nil)

	// 使用make函数初始化
	// make([]T, len, cap) cap可以省略,cap=len
	// s = make([]int, 2) // len = cap = 2
	// s = make([]int, 2, 4) // len = cap = 2
	s = make([]int, 2, 4)          // len =2, cap = 4
	fmt.Println(s, len(s), cap(s)) // [0 0] 2 4
	fmt.Println(s == nil)          // false

	s1 := make([]int, 0)
	fmt.Println(s1 == nil) // false

	s2 := make([]int, 0, 15)          // 一次把内存申请到位
	fmt.Println(s2, len(s2), cap(s2)) // ?

	// 如果你确定一个切片中最终要存储的元素个数，那么你最好一次把内存申请到位
}
