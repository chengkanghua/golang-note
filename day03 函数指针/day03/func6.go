package main

import (
	"fmt"
	"os"
)

// defer

func f44() {
	// 打开文件
	file, _ := os.Open("./func6.go")

	// 在函数结束的时候再帮我调用file.Close()
	defer file.Close()

	// 从文件中读取数据
	var data []byte
	file.Read(data)
	fmt.Printf("%s", data)
	// 关闭文件
	// file.Close()
}

func deferDemo1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值=x=5 -> defer操作 -> 底层RET返回
}

func deferDemo2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x(5)  -> defer操作  -> 底层的RET返回
}

func deferDemo3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值=y=x=5  -> defer操作 -> 底层RET返回
}

func deferDemo4() (x int) {
	defer func(x int) {
		x++
	}(x) // x当成参数传进匿名函数中
	return 5
}

func deferDemo5() {
	x := 10
	defer func() {
		res := sub(x, sub(10, 2))
		fmt.Println(res)
	}()

	x = 100

}
