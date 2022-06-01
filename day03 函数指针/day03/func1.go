package main

import "fmt"

func f1() {
	fmt.Println("hello~")
}

func sum(x int, s string, y int) int {
	res := x + y
	return res
}

// 连续多个参数的类型一致，可以使用参数类型简写
func sum2(x, y int) int {
	return x + y
}

func calc(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc2(x int, y int) (sum int, sub int) {
	// 命名返回值相当于在函数内部声明了返回值变量
	// var (
	// 	sum int
	// 	sub int
	// )
	sum = x + y
	sub = x - y
	return
}

func f4() (m map[string]int) {
	// var m map[string]int
	fmt.Println(m == nil)
	m = make(map[string]int, 4)

	m["王俊翔"] = 180
	m["波塞冬"] = 150
	m["王文建"] = 160
	m["李硕"] = 200
	return
}

func f3(string) {
	fmt.Println("嘿嘿嘿")
}

// 可变参数  --> 把所有参数放到一个切片中
// intSum 求一堆int的和
func intSum(x ...int) int {
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 可变长参数必须放在函数参数的末尾
func sum3(name string, x ...int) int {
	fmt.Println(name)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func someFunc(x string) []int {
	var s1 []int     // 声明切片变量
	var s2 = []int{} // 声明并初始化切片变量

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	if x == "" {
		return nil
	}
	fmt.Println(num) // 函数中可以访问到全局变量num
	return []int{1, 2, 3}
}
