package main

import (
	"fmt"
	"holiday/summer" // 第一级是go.mod里面的名字

	"github.com/q1mi/hello"
)

func main() {
	fmt.Println("打工人没有假期")

	hello.SayHi()

	// 调用本project下面的package
	summer.Diving()

}
