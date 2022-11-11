package main

import "fmt"

var (
	ddd = [3]string{"周公瑾"}
)

func f2() {
	var x = [3]int{1, 3, 5}

	y := x // 把x的值赋值给y

	y[0] = 100

	fmt.Println(y) // [100 3 5]
	fmt.Println(x) // [1 3 5]

	// 函数传参
	f3(x)
	fmt.Println(x) // [1 3 5]

	// 用f4的返回值给x赋值
	x = f4(x)      // x = [200 3 5] 赋值语句
	fmt.Println(x) // [200 3 5]

	// 函数中修改全局变量
	fmt.Println(ddd) // [周公瑾  ]
	f5(ddd)          // 把全局变量当成参数传给函数
	fmt.Println(ddd) // [周公瑾  ]

}

// f3 定义一个函数
// 参数类型为 [3]int
func f3(v [3]int) { // 通过函数传参，把x的值赋值给了v
	v[0] = 200
	fmt.Println("函数f3内v：", v) // [200 3 5]
}

func f4(v [3]int) [3]int {
	v[0] = 200
	return v
}

func f5(v [3]string) {
	v[0] = "大都督"
}

func f6() {
	ddd[0] = "大都督" // 直接修改全局变量
}
