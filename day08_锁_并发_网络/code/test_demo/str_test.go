package main

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	// 造一些假数据
// 	// 调用我们自己的函数
// 	res := Split("acbcc", "b")
// 	// 拿到结果和预期的作比较
// 	if len(res) != 2 {
// 		t.Fatal("测试失败")
// 	}
// 	// 又增加一个测试用例
// 	res2 := Split("上海自来水来自海上", "水")
// 	if len(res2) != 2 {
// 		t.Fatal("测试失败")
// 	}
// 	if res2[1] != "来自海上" {
// 		t.Log("测试失败", res2[1])
// 	}
// 	// 上海自来水来自海上 水

// 	// 在新增一个测试用例
// 	res3 := Split("acbcc", "bc")
// 	if len(res3) != 2 {
// 		t.Fatal("测试失败")
// 	}
// 	if res3[1] != "c" {
// 		t.Log("测试失败", res3[1])
// 	}
// }

func TestSplit(t *testing.T) {
	// 测试数据结构体，根据测试的函数自己定义
	type testData struct {
		s    string
		sep  string
		want []string
	}

	tt := []testData{
		{s: "acbcc", sep: "b", want: []string{"ac", "cc"}},
		{s: "上海自来水来自海上", sep: "水", want: []string{"上海自来", "来自海上"}},
		{s: "acbcc", sep: "bc", want: []string{"ac", "c"}},
		{s: "xx1x", sep: "1", want: []string{"xx", "x"}},
		{s: "aaa", sep: "bbb", want: []string{"aaa"}},
	}

	for _, tc := range tt {
		got := Split(tc.s, tc.sep)
		// 拿到结果和预期作比较
		// 自己写判断条件
		// if got[1] != tc.want[1] {
		// 	t.Errorf("测试失败, want:%#v but got:%#v", tc.want, got)
		// }
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("测试失败, want:%#v but got:%#v", tc.want, got)
		}
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
