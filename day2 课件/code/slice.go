package main

import "fmt"

// 切片 slice

func slice1() {
	// 声明切片变量但是未初始化（没申请内存空间）
	var s1 []int
	var s2 []bool
	var s3 []string

	fmt.Printf("%#v \n", s1) // []int(nil)
	fmt.Printf("%#v \n", s2) // []bool(nil)
	fmt.Printf("%#v \n", s3) // []string(nil)

	// var s4 []int
	// fmt.Println(s1 == s4) // 切片类型不能互相比较
	fmt.Println(len(s1)) // 求切片的长度（元素个数）
	fmt.Println(cap(s1)) // 求切片的容量（底层数组能存的元素个数）

	var dalichuqiji []string // 元素代表员工
	// 租了10个工位: cap
	// 1老板+4个员工=5: 长度（数量）
	// 增加了5个员工 ： 切片中又添加5个元素 len=cap=10

	// 又招了8个人 ： 切片中又添加8个元素
	// 公司换租一个更大的办公地址 <-> 切片底层的数组换了
	fmt.Println(dalichuqiji)

	var a1 [3]int
	fmt.Printf("%#v \n", a1) // [3]int{0 0 0}
}

func slice2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]         // s := a[low:high]
	fmt.Println(s)      // [2, 3]
	fmt.Println(len(s)) // 2
	fmt.Println(cap(s)) // 4

	s2 := a[3:]                       // 从索引为3开始切，切到最后
	fmt.Println(s2, len(s2), cap(s2)) // [4 5] 2 2
	s3 := a[:4]                       // 从索引为0开始切，切到索引为4 （左包含右不包含/顾头不顾尾）
	fmt.Println(s3, len(s3), cap(s3)) // [1 2 3 4] 4 5
	s4 := a[:]                        // 数组转切片
	fmt.Println(s4, len(s4), cap(s4)) // [1 2 3 4 5] 5 5
	// // fmt.Println(a == s4)              // a:数组类型 s4:切片类型 不同的不能比较

	// 数组是支持切片表达式
	a1 := [5]int{1, 2, 3, 4, 5}
	// a[low:high]   0 <= low < high <= len(a)
	s5 := a1[:5]
	fmt.Println(s5)

	// 字符串也支持切片表达式
	str := "hello"
	s6 := str[:3]            // 得到字符串
	fmt.Printf("%T\n", s6)   // string
	fmt.Println(s6, len(s6)) // "hel" 3

	// 切片再切片,high的取值范围<=cap
	a22 := [5]int{1, 2, 3, 4, 5}
	// a[low:high]   0 <= low < high <= len(a)
	s22 := a22[0:1]                      // 切片
	fmt.Println(s22, len(s22), cap(s22)) // [1] 1 5
	s33 := s22[:5]                       // 0 <= low < high <= cap(s22)
	fmt.Println(s33, len(s33), cap(s33)) // [1 2 3 4 5] 5 5

	s44 := s22[4:5]                      // 0 <= low < high <= cap(s22)
	fmt.Println(s44, len(s44), cap(s44)) // [5] 1 1

	a22[4] = 500     // 修改数组中的值
	fmt.Println(s33) // [1 2 3 4 500]
	fmt.Println(s44) // [500]
}

// slice3 切片完整表达式
func slice3() {
	// 默认切片的容量是从切片的开始索引到数组的最后
	// max: 影响切片的容量
	// max:想象成high能取到的最大值
	// 最终切片的容量：max-low
	a := []int{1, 2, 3, 4, 5}
	// a[low:high:max]
	s1 := a[1:2:3]                    // 0 <= low <= high <= max <= cap(a)
	fmt.Println(s1, len(s1), cap(s1)) // [2] 2 2
}

