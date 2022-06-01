package service

import "fmt"

var Name string

func InitName(name string) {
	Name = name
}

func f1() {
	fmt.Println(Name)
}
