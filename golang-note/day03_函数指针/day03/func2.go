package main

import "fmt"

// 全局变量尽量少用
// 用的话变量名也要有区分度

var num = 100
var name = "杨俊"

func f11() {
	fmt.Println(num)
	num = 200 // 赋值
}

func f12() {
	fmt.Println(num)
}

func f13() {
	fmt.Println(name) // 全局变量
	name = "周公瑾"      // 修改全局变量

	name := "王俊翔" // 又重新声明了一个函数内部有效的变量
	fmt.Println(name)
}

func f14() {
	var m = map[string]int{
		"李硕": 100,
		"杨俊": 200,
	}

	// v和ok是不是只在if条件判断语句中有效
	if v, ok := m["波塞冬"]; ok {
		fmt.Println(v)
	}

	v, ok := m["波塞冬"]
	if ok {
		fmt.Println(v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	switch i := 0; i {
	case 100:
		
	}

	fmt.Println(v)
}
