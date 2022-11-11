package main

import "fmt"

// f1 多维数组
func f1() {
	var x = [3]int{1, 2, 3}
	fmt.Println(x)
	var xx = [3][2]string{
		{"北京", "石家庄"},
		{"上海", "苏州"},
		{"成都", "重庆"}, // 注意：最外层的花括号换行则这里必须加,
	}
	// var xx = [3][2]string{{"北京", "石家庄"}, {"上海", "苏州"}, {"成都", "重庆"}}

	// fmt.Println(xx)
	// fmt.Printf("%#v\n", xx)

	// 二维数组的遍历
	for i := 0; i < len(xx); i++ {
		tmp := xx[i]
		// 第一层
		fmt.Printf("xx[%v]:%v\n", i, xx[i])
		// 第二层
		for j := 0; j < len(tmp); j++ {
			fmt.Printf("\t xx[%v][%v]:%v\n", i, j, tmp[j])
		}
	}
	fmt.Println(xx)
	fmt.Printf("%#v\n", xx)

	// 二维数组的遍历
	for i := 0; i < len(xx); i++ {
		tmp := xx[i]
		// 第一层
		fmt.Printf("xx[%v]:%v\n", i, xx[i])
		// 第二层
		for j := 0; j < len(tmp); j++ {
			fmt.Printf("\t xx[%v][%v]:%v\n", i, j, tmp[j])
		}
	}

	for i1, v1 := range xx {
		fmt.Printf("xx[%v]: %v\n", i1, v1)
		// 第二层
		for i2, v2 := range v1 {
			fmt.Printf("\txx[%v][%v]:%v\n", i1, i2, v2)
		}
	}
}
