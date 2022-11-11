package main

import "fmt"

// append函数往切片中追加元素
// 使用append函数时必须接收返回值！！！

func appendDemo() {
	var s = []string{"北京"}
	// apend函数可能触发切片的扩容
	// 切片扩容之后会有一个新的底层数组，需要更新变量s
	s = append(s, "上海")
	fmt.Println(s) // [北京 上海]

	s2 := []string{"广州", "深圳"}
	s = append(s, s2...) // ...表示将s2拆开一个一个元素追加
	fmt.Println(s)       // [北京 上海 广州 深圳]

	// 零值切片可以直接在append中使用
	var s3 []int // nil
	fmt.Println(s3 == nil)
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3, 4)
	fmt.Println(s3)
}

// appendDemo2 使用append函数触发扩容
// 导致意想不到的事情发生
func appendDemo2() {
	var s = []string{"北京"}
	_ = append(s, "上海", "广州", "深圳")
	fmt.Println(s) // [北京]
}

// appendDemo3 append函数导致切片扩容示例
func appendDemo3() {
	var s = []string{"北京"}
	fmt.Println(len(s), cap(s)) // len = cap = 1
	s = append(s, "上海", "广州", "深圳")
	fmt.Println(len(s), cap(s)) // 4 4
	fmt.Println(s)              // [北京]
}

// deleteSlice 删除切片中的元素
func deleteSlice(idx int) {
	idx = 1
	var s = []int{1, 2, 3}
	s = append(s[:idx], s[idx+1:]...)
	fmt.Println(s)
}

func appendDemo4() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("%#v\n", a) // ["" "" "" "" "" "0" "1" "2" "3" "4" "5" "6" "7" "8" "9"]
	fmt.Println(len(a))    // 15
	fmt.Println(cap(a))    // 触发了自动扩容，不确定！

	// 格式化字符串
	name := "qimi"
	s := fmt.Sprintf("%s讲的真好", name)
	fmt.Println(s)
}
