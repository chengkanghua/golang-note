package main

import (
	"fmt"
)

// 结构体的继承

type Animal struct {
	name string
}

func (a Animal) Move() {
	fmt.Printf("%s会动...\n", a.name)
}

type Dog struct {
	Animal // 嵌入Animal结构体来实现类似继承的效果
	leg    int
}

func (d Dog) wang() {
	fmt.Println("汪汪汪~")
}

func demo12() {
	var d = Dog{
		leg:    4,
		Animal: Animal{name: "旺财"},
	}
	d.wang()
	d.Move()
}
