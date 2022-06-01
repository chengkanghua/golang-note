package main

import "fmt"

// 指针

func f66() {
	s := "永远不要高估自己" // s是一个字符串变量
	p := &s         // 取变量s的内存地址
	// 问：p是什么类型？ --> 指针类型 --> 字符串指针
	fmt.Printf("p:%T\n", p)

	i := 10
	p2 := &i // 指针类型  --> int指针
	fmt.Printf("p2:%T\n", p2)
	b := [3]int{1, 2, 3}

	p3 := &b // 指针类型  --> [3]int指针
	fmt.Printf("p3:%T\n", p3)
}

func f67(x *int) {
	*x = 100
}

func f68() {
	var a *int
	fmt.Println(a)
	a = new(int) // *int
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)

	var x int // int型变量x
	px := &x  // *int
	fmt.Println(px == nil)

	var b map[string]int
	b = make(map[string]int)  // 申请内存
	b["杨俊"] = 100
	fmt.Println(b)
}
