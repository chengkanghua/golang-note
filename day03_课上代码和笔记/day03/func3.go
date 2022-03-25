package main

import (
	"fmt"
)

type MySlice []int

type MyMap map[string][]int

// 自定义类型
// 类型名：f
// 具体是一个没有参数没有返回值的函数类型
type myFunc func() // 等到讲web编程会用到

func f22() {
	var x int
	x = 100
	fmt.Println(x)

	var z MySlice
	z = []int{1, 2, 3}
	fmt.Println(z)

	var fx myFunc
	fx = func() { fmt.Println("写成一行我蒙圈!") }

	fx()
}

// 将函数作为一个类型能够提高代码的可读性
// 可以为其定义方法

type Fa func(x, y int)
type Fb func(name string, score int)

func ff(f1 Fa, f2 Fb) func() {
	// ...
	return func() {}
}

// 函数签名 --> 函数定义（声明）的格式,与参数名称和返回值名称无关
// 函数签名一样 --> 函数的参数、返回值的类型和个数、顺序都要一样

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(a, b int) int {
	return a - b
}

func f15() {
	var x calculation
	fmt.Printf("%T\n", x) // main.calculation
	fmt.Println(x == nil) // true
	x = add               // 把add赋值给x
	res := x(10, 20)      // ???
	fmt.Println(res)

	add(10, 20)
}

func fi(name string, age int) {}

func fj(age int, name string) {}

type MyFF func(string, int)

func f16() {
	var mf MyFF
	mf = fi
	// mf = fj  // 函数签名不一致
	mf("ddd", 1)
}

func f17(x, y int, op calculation) int {
	res := op(x, y)

	return res
}

// 命名返回值
// 1.函数内部声明了一个变量res
// 2.返回值是res
func f18(x, y int, s string) (res func(int, int) int) {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	}
	return  // 默认就把res返回
}

func f19(x, y int, s string) func(int, int) int {
	var res func(int, int) int // res = nil
	switch s {
	case "+":
		return add
	case "-":
		return sub
	}
	return res
}
