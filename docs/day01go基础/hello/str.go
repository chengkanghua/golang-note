package main

import (
	"fmt"
	"strings"
)

func f2() {
	// filename 表示windows下一个文件路径
	filename := "C:\\go\\hello\\hello.exe"
	fmt.Println(filename)

	s11 := "永远不要高估自己"
	fmt.Println(s11)

	s12 := "\"永远不要高估自己\""
	fmt.Println(s12)

	// 多行字符串
	s13 := `多行
字符串
	测\n试
	`
	fmt.Println(s13)

	// 字符串操作
	fmt.Println(len(s11))
	// 字符串拼接
	name1 := "jade"
	value1 := "过年好"
	fmt.Println(name1 + value1)

	ret := fmt.Sprintf("大家好，%s祝大家%s", name1, value1)
	fmt.Println(ret)

	// strings
	s14 := "你:好:呀"
	fmt.Println(strings.Split(s14, ":"))

	fmt.Println(strings.Contains(s14, "你"))
	fmt.Println(strings.HasPrefix(s14, "你:")) // true
	fmt.Println(strings.HasSuffix(s14, "啊"))  // false

	fmt.Println(strings.Index(s14, ":"))     // 3
	fmt.Println(strings.LastIndex(s14, ":")) // 7

	// 拼接
	slice1 := []string{"你", "我", "他"}
	fmt.Println(strings.Join(slice1, "-"))

	// 字符和字符串
	y1 := '中' // 字符
	y2 := "中" // 字符串
	fmt.Println(y1, y2)

	// byte 和rune
	fmt.Println([]rune(s14))
	fmt.Println([]byte(s14))
	// for range循环
	idx := 0
	for _, r := range s14 { // rune表示一个汉字
		if r == ':' {
			fmt.Println(idx)
			break
		}
		idx++
	}
}
