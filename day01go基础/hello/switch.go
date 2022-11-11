package main

import "fmt"

// switch

func f8() {
	finger := 3 // 从外界获取的一个值
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("🖕🏻")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入")
	}

	num := 9
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}

	switch {
	case num%2 != 0:
		fmt.Println("奇数")
	case num%2 == 0:
		fmt.Println("偶数")
	default:
		fmt.Println("num=0")
	}
}
