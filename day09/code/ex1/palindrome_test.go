package ex1

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	type testData struct {
		s    string
		want bool // 期望结果
	}

	// 测试用例，包含参数和期望的结果
	tt := []testData{
		{s: "Madam,I’mAdam", want: false},
		{s: "油灯少灯油", want: true},
		{s: "上海自来水来自海上", want: true},
		{s: "12321", want: true},
	}

	// 挨个执行测试用例
	for _, rc := range tt {
		got := Palindrome2(rc.s) // 调用我们的待测试函数
		if got != rc.want {      // 拿函数的返回结果与预期去比较
			t.Errorf("%s want %#v,but got:%#v", rc.s, rc.want, got)
		}
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Palindrome("Madam,I’mAdam") // 全循环版本
		Palindrome2("Madam,I’mAdam") // 半循环版本
	}
}
