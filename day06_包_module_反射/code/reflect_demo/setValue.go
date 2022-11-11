package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 因为要让函数修改我们传进来的值，所以只能传指针
	// 反射里面拿到 内存地址 得到的是一个指针类型
	// 而我们需要的是值的类型，有了类型再调用具体的方法去修改值
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func f1() {
	var num int64 = 100
	// reflectSetValue1(num)
	reflectSetValue2(&num)
	fmt.Println(num)
}
