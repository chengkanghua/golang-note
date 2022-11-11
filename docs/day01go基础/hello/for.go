package main

import "fmt"

// for循环

func f7() {
	// 1.标准for循环
	for i := 0; i <= 10; i++ {
		fmt.Println(i) // 0 1 2 ... 10
	}
	// fmt.Println(i) // 不可访问i

	// 2.初始语句省略
	i := 0
	for ; i <= 10; i++ {
		fmt.Println(i) // 0 1 2 ... 10
	}
	fmt.Println(i) // ? 11

	// 3.初始语句和结束语句都可以省略
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++ // 10
	}
	fmt.Println(j) // ? 10

	// 4. 无限循环
	for {
		if j > 12 {
			break // 跳出循环
		}
		fmt.Println("...")
		j++
	}

	// for range 循环
	s := "golang"
	for i, v := range s {
		fmt.Printf("%v:%c \n", i, v)
	}
}
