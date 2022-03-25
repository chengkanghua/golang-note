package main

import "strings"

// 闭包

func adder() func() int {
	var x int

	// 函数内部使用了它外部函数的变量x
	f := func() int {
		x++
		return x
	}

	// 把匿名函数当成返回值返回了
	return f
}

func adder2() func(int) int {
	var x int
	f := func(y int) int {
		x += y // x = x+y
		return x
	}
	return f
}

// 定义一个累加器
// x的值会被返回的函数一直占用着
func adder3(x int) func(int) int {
	f := func(y int) int {
		x += y // x = x+y
		return x
	}
	return f
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
