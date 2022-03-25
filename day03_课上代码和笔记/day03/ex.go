package main

import (
	"fmt"
	"sort"
	"strings"
)

func ex1() {
	// 写一个程序，统计一个字符串中每个单词出现的次数。
	// 比如：”how do you do”中how=1 do=2 you=1。
	s := "how do you do"

	// 1.用map存数据，key是单词，value是单词出现的次数
	// 2.将字符串分成一个一个的单词
	// 3.把上一步得到的单词挨个存放到map里
	// 4.遍历map打印结果

	// 迎刃而解
	var m map[string]int
	m = make(map[string]int)
	s1 := strings.Split(s, " ") // 切片
	for _, v := range s1 {
		// m[v] = 1  // m["do"] = 1
		num := m[v]
		m[v] = num + 1

		// m[v]++

		// if ok{
		// 	m[v] = num+1
		// }else{
		// 	m[v] = 0+1
		// }
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func ex2() {
	// 对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
	var a = [...]int{3, 7, 8, 9, 1}
	s := a[:]
	fmt.Println("before sort:", s)
	// 按什么规则对s排序
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println("after sort:", s)
}

func ex3() {
	m := make(map[string][]int) // 声明并初始化了一个map变量m
	s := []int{1, 2}            // 声明一个切片变量s
	s = append(s, 3)            // 向s追加一个元素
	fmt.Printf("%+v\n", s)      // [1 2 3]

	m["q1mi"] = s
	s3 := s

	// 用append从切片删除索引为1的元素 append(s[:idx], s[idx+1]...)
	s = append(s[:1], s[2:]...) // ...表示拆开切片
	// s[:1]     [1]
	// s[2:]...  3
	s[1] = 100 // 修改了底层数组

	fmt.Printf("s:%+v\n", s) // [1 3]

	fmt.Printf(`m["q1mi"]:%+v`+"\n", m["q1mi"]) // [1 3 3]
	fmt.Printf("s3:%+v\n", s3)                  // [1 3 3]

	fmt.Println(len(s), len(m["q1mi"]))
}

func ex4() {
	s := []int{1, 2, 3}

	s3 := s
	// s[1] = 3

	s = append(s[:1], 3)

	// 切片的容量够用就不会换底层数组

	// s = s[:1]  // [1]
	// s = append(s, 3)  // [1 3]

	// s[1] = 100

	fmt.Println(s, len(s), cap(s))    // [1 3 3] 2 3
	fmt.Println(s3, len(s3), cap(s3)) // [1 3 3] 3 3
}
