package main

import "fmt"

// 九九乘法表

func f10() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}

}

// 10万零1个数字, 除了一个数字只出现1次之外，其他数字都出现了2次，
// 问，如何找出只出现一次的那个数字
func f11() {
	nums := []int{17, 4, 3, 3, 9, 11, 9, 11, 17}
	if len(nums)%2 == 0 {
		return
	}
	ret := nums[0]
	for _, num := range nums[1:] {
		ret ^= num // 二进制的异或
	}
	fmt.Println(ret)

	// 01 10 10 

}
