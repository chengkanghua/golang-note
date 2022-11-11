package main

import "fmt"

// 指针接收者和值接收者

type Mover interface {
	Move()
}

type Dog struct {
	name string
}

// Move 值接收者实现Move方法
// func (d Dog) Move() {
// }

// Move 指针接收者实现Move方法
func (d *Dog) Move() {
	fmt.Println(d.name)
}

// 使用指针接收者实现接口
// 接口变量可以接收指针类型但是不能接收值类型 （不是任何值都能取地址）

// 使用值接收者实现接口
// 接口变量既能接收指针类型又能接收值类型（有了地址就能取值）

type Myint int

func (i *Myint) Move() {
}

func main() {
	var x Mover

	var d Dog // 结构体值
	d.Move()  // &d.Move()

	// Myint(10).Move() // 10是字面量 没有办法取地址

	n := Myint(10)
	n.Move() // n是变量，可以取地址

	// x = &Myint(10) // 字面量10

	fmt.Println(x)

}
