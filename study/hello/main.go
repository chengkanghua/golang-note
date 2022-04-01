package main

import (
	"fmt"
)

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d Dog) Move() {
	fmt.Println("狗会动")
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Println("汽车在跑")
}

func main() {
	var n Mover = &Dog{Name: "旺财"}
	v, ok := n.(*Dog)
	fmt.Println(v, ok)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵"
	} else {
		fmt.Println("类型断言失败")
	}

}
