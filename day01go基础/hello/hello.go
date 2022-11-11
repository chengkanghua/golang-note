package main

import "fmt"

var version string
var age11 = 18

// name12 := "小王子"  // 函数外语句必须以关键字开头

func main() {
	fmt.Println("Hello world!")

	/*
		多行注释
	*/
	// 变量的声明
	var name string // 声明变量

	// 批量声明
	var (
		age  int  // 0
		isOk bool // false
	)
	// var age int
	// var isOk bool

	age = 100 // 变量赋值
	fmt.Println(name, age, isOk)

	var age2 int = 18 // 声明变量并赋值
	fmt.Println(age2)

	// 没有指定类型？
	var name3, age3 = "jade", 28

	// var (
	// 	name3 string = "jade"
	// 	age3 int = 28
	// )

	fmt.Println(name3, age3)

	var age4 int8 = 28 // 如果不想用编译器推导的类型，就需要显式指定变量的类型
	fmt.Println(age4)

	// 双引号表示字符串，单引号表示字符

	var x byte = 'a'   // 字符
	var s string = "a" // 字符串
	fmt.Println(x, s)

	// 短变量声明
	s2 := "jade" // var s2 string s2="jade"
	fmt.Println(s2)
	s2 = "小王子"
	fmt.Println(s2)

	// var x2 string
	// x2 = 18 // 只能给变量赋正确类型的值
	// fmt.Println(x2)

	// 常量
	fmt.Println(c4)
	fmt.Println(d3, d4)

	f1()
	f2()

	f3()
	f4()

	f7()
	f8()

	// gotoDemo2()

	f10()
	f11()
}
