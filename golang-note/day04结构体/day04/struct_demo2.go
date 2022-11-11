package main

import (
	"fmt"
	"unsafe"
)

// 结构体的内存布局
// 结构体属性占用的内存空间是连续的

// MyStruct2 定义一个空结构体
type MyStruct2 struct{}

func demo6() {
	// 查看结构体大小
	var v1 MyStruct
	// unsafe.Sizeof()  // 获取变量占用内存的大小
	fmt.Println(unsafe.Sizeof(v1)) // 多大?
	// 拿到属性a的内存地址
	fmt.Println(&(v1.a))
	// 拿到属性b的内存地址
	fmt.Println(&(v1.b))
	fmt.Println(&(v1.c))
	fmt.Println(&(v1.d))

	// 空结构体不占用内存空间
	var v2 MyStruct2
	fmt.Println(unsafe.Sizeof(v2)) // 0

	// 空结构体的应用场景1

	// var set map[string]int

	// var set1 map[string]struct{}

	// 给一堆人名，要去重
	nameList := []string{"张三", "李四", "王五", "jade", "张三", "王五"}

	type empty struct{}
	// 利用map的key不能重复
	// var nameMap = make(map[string]bool)
	// var nameMap = make(map[string]empty)

	var nameMap = make(map[string]struct{})
	for _, name := range nameList {
		nameMap[name] = struct{}{}
	}

	// 经过去重的name
	for key := range nameMap {
		fmt.Println(key)
	}

	// set1 := map[string]struct{}{}
	// fmt.Println(set1)
}

func demo7() {
	var v1 MyStruct         // 结构体变量的内存地址 = 第一个属性的地址
	fmt.Printf("%p\n", &v1) // 0xc0000180a4
	fmt.Println(&v1.a)      // 0xc0000180a4
	fmt.Println(&(v1.a))    // 0xc0000180a4
}

type MyStruct struct {
	a int8 // 1byte
	b int8 // 1byte
	c int8 // 1byte
	d int8 // 1byte
}

// type MyStruct3 struct {
// 	a int8  // 1byte
// 	b int32 // 4bytes
// 	c int8  // 1byte
// 	d int64 // 8bytes
// }

type MyStruct3 struct {
	a int8  // 1byte
	c int8  // 1byte
	b int32 // 4bytes
	d int64 // 8bytes
}

func demo8() {
	var v3 MyStruct3
	fmt.Println(unsafe.Sizeof(v3)) // ? 14  24?? 编译器自动帮我们做了内存对齐
}

type MyStruct4 struct {
	n struct{} // 0
	m int8     // 1
}

// type MyStruct4 struct {
// 	m int8     // 1
// 	n struct{} // 0
// }

func demo9() {
	var v4 MyStruct4
	fmt.Println(unsafe.Sizeof(v4)) // 1
}

// newMyStruct3 构造函数
func newMyStruct3(a int8) MyStruct3 {
	return MyStruct3{
		a: a,
	}
}	
