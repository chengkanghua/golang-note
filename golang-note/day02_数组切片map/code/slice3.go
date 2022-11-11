package main

import "fmt"

// 切片的遍历

func slice7() {
	s := []int{1, 2, 3, 4}

	// 索引遍历
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	// for range遍历
	for i := range s {
		fmt.Println(i)
	}
	for i, v := range s {
		fmt.Println(i, v)
	}
	for _, v := range s {
		fmt.Println(v)
	}
}
