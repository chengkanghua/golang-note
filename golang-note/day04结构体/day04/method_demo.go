package main

import "fmt"

// 方法

type Person struct {
	name string
	age  int
}

// newPerson 构造函数
func newPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

// dream 给person定义一个方法
// p ==》 this， self
func (p Person) dream(s string) {
	fmt.Printf("%s的梦想是%s\n", p.name, s)
}

// 使用结构体值作为接收者定义方法
// func (p Person) guonian() {
// 	p.age++ // 值拷贝，改副本
// }

// 使用结构体指针作为接收者定义方法
func (p *Person) guonian() {
	p.age++
}

func demo10() {
	p1 := newPerson("尼勒克", 60)
	// p := Person{
	// 	name: "尼勒克",
	// 	age:  60,
	// }
	// p1.dream("吃喝拉撒")
	p1.guonian()
	fmt.Println(p1.age) // ?
}
