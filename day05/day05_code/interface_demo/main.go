package main

import "fmt"

// 定义一个学生结构体类型  --> 我是谁
type student struct {
	name string
	age  int
}

func (s student) dream() {
	fmt.Printf("%s的梦想是找个对象。\n", s.name)
}

// 定义一个dreamer接口类型  --> 我能干什么
type dreamer interface {
	dream()
}

type writer interface {
	write()
}

// 自定义类型、结构体类型、都可以添加方法
type Myint int

func (m Myint) dream() {
	fmt.Println("Myint的梦想是下辈子做个人")
}

func main() {

	var s = student{name: "杨俊"} // 声明一个student类型的变量s

	makeDream(s)

	var x dreamer // 声明一个dreamer接口类型变量x

	// 接口是一种抽象的类型
	// 只要满足接口要求的方法就能当成对应接口的变量
	// 把student变量当成dreamer接口类型的变量
	x = s
	x.dream()

	var i Myint = 10
	x = i
	makeDream(i)

	d := dog{}
	x = d
	makeDream(d)

}

type dog struct{}

func (d dog) dream() {
	fmt.Println("狗的梦想是吃狗粮。")
}

// makeDream 让谁做梦的函数
func makeDream(x dreamer) {
	// 只关心传进来的参数能调用dream方法
	x.dream()
}
