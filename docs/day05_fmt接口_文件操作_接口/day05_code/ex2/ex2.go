package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)

	// x := []int{1, 2, 3}

	// x := []int{
	// 	1,
	// 	2,
	// 	3,
	// }

	// stus := []student{
	// 	student{name: "小王子", age: 18},
	// 	student{name: "娜扎", age: 23},
	// 	student{name: "大王八", age: 9000},
	// }

	stus := []student{
		{name: "杨俊", age: 18},
		{name: "吴勇", age: 23},
		{name: "jade", age: 9000},
	}

	// 将切片中的学生赋值到map中
	// 编程世界里的刻舟求剑
	// for _, stu := range stus {
	// 	m[stu.name] = &stu
	// }

	for i := 0; i < len(stus); i++ {
		stu := stus[i]
		m[stu.name] = &stu
	}

	// 遍历map中的键值对
	for k, v := range m {
		fmt.Printf("k:%v v:%p \n", k, v)
	}

	// 杨俊 => 杨俊
	// 吴勇 => 吴勇
	// jade => jade
}
