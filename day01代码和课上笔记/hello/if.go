package main

import "fmt"

// if条件判断分支

func f5() {
	score := 89 // 假设从数据库中查询出一个同学的分数
	if score > 90 {
		fmt.Println("A")
	} else if score > 65 {
		fmt.Println("勉强留下")
	} else {
		fmt.Println("明年再来")
	}
	fmt.Println(score)
}

func f6() {
	// score只在if分支中有效
	// 因为它只在if分支中声明了score，外部不可见
	if score := 89; score > 90 {
		fmt.Println("A")
	} else if score > 65 {
		fmt.Println(score)
		fmt.Println("勉强留下")
	} else {
		fmt.Println("明年再来")
	}
	// fmt.Println(score)
}
