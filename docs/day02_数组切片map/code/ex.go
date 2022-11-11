package main

import "fmt"

// 练习1:求数组元素的和
func arraySum() {
	var a = [...]int{1, 3, 5, 7, 8, 10}
	sum := 0
	for _, value := range a {
		sum = sum + value
		// sum += value  // 简写
	}

	fmt.Println(sum)
}

// 练习2：求数组中元素和为8的元素的下标(索引)
// 1 + 7 = 8;索引(0,3)
// 3 + 5 = 8;索引(1,2)
// 1.拿到数组中的每一个元素 ：遍历数组
// 2.找到元素和为8的那两个元素 ： 数学运算和比较运算
// 3.把符合要求的索引打印出来  ： fmt.Println()
func arraySum2() {
	var a = [...]int{1, 3, 5, 7, 8, 10}

	// 正向思路
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				// 找到啦
				fmt.Println(i, j)
			}
		}
	}

	// 反向思路
	for i, v := range a {
		x := 8 - v // 要找的目标值
		for j := i + 1; j < len(a); j++ {
			if a[j] == x {
				fmt.Println(i, j)
			}
		}
	}
}
