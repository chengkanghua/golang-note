package main

import (
	"strings"
)

// Split 按指定字符串切割，
// "acbcc" "b"
// 上海自来水来自海上 水
func Split(s, sep string) (result []string) {
	count := strings.Count(s, sep)
	// 数一下s里面有多少个sep ,算出最终切割的结果有多少项
	// 然后把内存一次申请到位
	result = make([]string, 0, count+1)

	i := strings.Index(s, sep) // 2

	for i > -1 {
		result = append(result, s[:i]) // []string{"ac"}
		s = s[i+len(sep):]             // "cc"
		i = strings.Index(s, sep)      // -1
	}
	result = append(result, s) // []string{"ac", "cc"}
	return
}
