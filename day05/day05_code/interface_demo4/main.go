package main

import (
	"fmt"
)

// 空接口

type Any interface{}

func main() {

	// var x Any
	var x interface{} // 定义空接口类型变量x

	fmt.Println(x) // 接口的默认零值就是nil
	x = 100
	fmt.Printf("%T \n", x)
	x = "杨俊"
	fmt.Printf("%T \n", x)
	x = true
	fmt.Printf("%T \n", x)
	x = struct{}{}
	fmt.Printf("%T \n", x)
	x = new(int)
	fmt.Printf("%T \n", x)

	// 类型断言
	b, ok := x.(bool)
	// 猜对了就把接口变量转为对应的类型
	// 猜错了b默认零值
	fmt.Println(b, ok)

	p, ok := x.(*int)
	fmt.Println(p, ok) // 0xc0000180c0 true

	// map[string]int  // {"name":"杨俊", "age": 28}
	// 2.空接口类型作为map的value的类型 实现数据的兼容
	m := map[string]interface{}{
		"name": "杨俊",
		"age":  28,
	}
	fmt.Println(m)
}

type Dog struct {
	Name string
}

func foo(x interface{}) { // 空接口作为函数的参数类型
	if v, ok := x.(Dog); ok {
		v.Name = "旺财"
	}



}
