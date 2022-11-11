package ex1

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 回文判断函数
// 编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。
// （回文：一个字符串正序和逆序一样，如“Madam,I’mAdam”、“油灯少灯油”等。）

// Palindrome 回文检测 列表转换的方法
func Palindrome(s string) bool {
	var s1 string
	var temp = strings.Split(s, "")
	//fmt.Println(utf8.RuneCountInString(s)) // 中文的len 方法，直接使用len 他会将中文转成字节码
	// utf8.RuneCountInString(s) 数一下字符串中的rune数
	for i := 0; i < utf8.RuneCountInString(s); i++ {
		s1 += temp[len(temp)-1-i]
	}
	return s1 == s
}

// 油灯少灯油 =》temp:[“油“，”灯“，”少“，”灯“，”油”]
// arr: "油灯少灯油"  => 得到一个逆序的字符串
// arr == s

func Palindrome2(s string) bool {
	// s -> []string
	// s:= "a:b:c" strings.Split(s, ":") => ["a", "b", "c"]
	var temp = strings.Split(s, "")

	for i, j := 0, len(temp)-1; i < j; {
		if temp[i] != temp[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// func main() {
// 	Palindrome("油灯少灯油")
// }

func ExamplePalindrome() {
	fmt.Println(("油灯少灯油"))
	// Output:
	// true
}
