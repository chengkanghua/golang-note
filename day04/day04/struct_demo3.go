package main

import "fmt"

// 结构体嵌入

type Address struct {
	Province   string
	City       string
	updateTime int64
}

type Email struct {
	updateTime int64
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string // 性别
	Address        // 结构体类型
	Email
}

func demo11() {
	u1 := User{
		Name:   "jade",
		Gender: "男",
		Address: Address{
			Province: "北京",
			City:     "北京",
		},
	}
	fmt.Printf("%+v\n", u1)

	u1.City = "上海"
	// 匿名嵌入的字段冲突，需要写明
	u1.Email.updateTime = 160000000

	fmt.Println()
}
