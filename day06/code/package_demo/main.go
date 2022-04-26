package main

import (
	"encoding/json"
	"fmt"
	"sort"
	// j "encoding/json" // 起别名 慎用
	// . "encoding/json" // 起别名 慎用
)

func reverseDemo() {
	var x = []int{8, 5, 7, 88, 57}
	sort.Ints(x) // 排序
	fmt.Println(x)

	x2 := sort.Reverse(sort.IntSlice(x))
	sort.Sort(x2)
	fmt.Println(x2)
}

// 结构体首字母小写 别的包导入是无法使用的， 但是下面的直接把值赋值过去的可以，
type student struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}

func jsonDemo() {
	var s = student{
		Name: "杨俊",
		age:  26,
	}
	// b, err := j.Marshal(s)
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("%s\n", b) //{"name":"杨俊"}  因为结构体age首字母小写，json包读不到所以解析不到
}

func main() {
	// reverseDemo()
	jsonDemo()
}
