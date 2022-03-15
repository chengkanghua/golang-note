package main

import "fmt"

func f9() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 结束本轮循环，继续下一次循环
		}
		fmt.Println(i)
	}
}
