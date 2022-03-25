package main

import "fmt"

// 自定义类型与类型别名

// 自定义类型
type MyFunc func(int, int) int

type MyInt int

type scoreMap map[string]int

// 类型别名：用来给类型起别名，方便理解
type NewInt = int

func demo1() {
	var b byte //
	b = 'a'
	fmt.Println(b)

	// var c uint8
	var c byte
	c = 'a'
	fmt.Println(c)

	// var r int32
	var r rune
	r = '中'
	fmt.Println(r)

	// MyInt和NewInt有什么区别
	var x MyInt = 100
	fmt.Printf("x:%T\n", x) // main.MyInt

	var y NewInt = 100
	fmt.Printf("y:%T\n", y) // int

	// x = MyInt(y)  // 类型强制转换
	// y = NewInt(x) // 类型强制转换

	f := 1.123
	i := int(f) // 浮点数可以强制转换成整数，但是会丢失精度
	fmt.Println(i)

}
