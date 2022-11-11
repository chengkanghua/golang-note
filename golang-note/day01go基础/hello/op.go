package main

import "fmt"

func f4() {
	x1 := 9
	x2 := 2
	fmt.Println(x1 + x2) // 11
	fmt.Println(x1 - x2) // 7
	fmt.Println(x1 * x2) // 18
	fmt.Println(x1 / x2) // 4
	fmt.Println(x1 % x2) // 1

	// 自增 和 自减 都是语句
	x2++
	// x2 = x2++  ❌
	fmt.Println(x2) // 3

	// 逻辑运算符
	flag1 := false
	flag2 := 1+3 >= 4

	fmt.Println(flag1 && flag2)
	fmt.Println(flag1 || flag2)
	fmt.Println(!flag2)

	// 位运算
	bb1 := 0b1001
	bb2 := 0b1100

	fmt.Printf("%b \n", bb1&bb2) // 与  0b1000
	fmt.Printf("%b \n", bb1|bb2) // 或  0b1101
	fmt.Printf("%b \n", bb1^bb2) // 异或 0b101

	fmt.Printf("1<<3:%b \n", 1<<3)      // 1000
	fmt.Printf("1>>3:%b \n", 0b1000>>3) // 1

	vv1 := 10
	vv1 += 20 // <==> vv1 = vv1 + 20
	vv1 /= 20 // <==> vv1 = vv1 / 20
	// ...
	fmt.Println(vv1)

}
