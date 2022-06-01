package main

import (
	"fmt"
	"sort"
)

// 接口组合

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// type ReadWriter interface {
// 	Read()
// 	Write()
// }

type ReadWriter interface {
	Reader
	Writer
}

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// reverse 结构体中嵌入了Interface接口
type reverse struct {
	Interface
}

// 结构体嵌入匿名字段
// 找方法和属性的时候，先从自己的找，找不到再去内部匿名字段找

func (reverse) Len() int {
	return 100
}

func main() {
	// var x ReadWriter
	// x.Read([]byte{})

	// var i reverse

	// i.Len()

	var i Interface // 接口类型的零值是nil
	fmt.Println(i == nil)

	var p Person

	fmt.Println(p.city)      // 类型零值 空字符串
	fmt.Println(p.Interface) // nil
	p.showAddr()

	// 接口类型变量的初始化 --》 找一个实现了该接口类型的变量赋值进去
	p.Interface = sort.IntSlice{}

	p.Len() // ? p.Interface.Len()  nil.Len()

}

type Address struct {
	city string
}

func (a Address) showAddr() {
	fmt.Println(a.city)
}

type Person struct {
	name      string
	Address   // 嵌入匿名字段
	Interface // 嵌入一个匿名字段
}
