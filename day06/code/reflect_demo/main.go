package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 反射

type Student struct { // {"name": "". "age": 16}
	Name string
	Age  int
}

func main() {

	// var x Order
	// json.Marshal(x)

	// var a Student
	// reflectType(a)

	// var b Order
	// reflectType(b)

	// f1()
	marshalDemo()
}

// 利用反射获取到实际的类型
func reflectType(x interface{}) {
	v := reflect.TypeOf(x) // 拿到实际的类型
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

type Order struct { // {"id": 1111123}
	Id int64 `json:"id" xml:"xxx"`
}

func marshalDemo() {
	// s := `{"id": 1234567}`
	s := `id=1234567`

	var o Order // 声明了一个结构体变量o

	// 做了啥?
	// 1. 读取json字符串中 key - value 信息
	// 2. 拿着读取到的key 去结构体里面 依次查找字段，找到结构体中同名字段
	// 3. 把value赋值给结构体的这个字段
	// json.Unmarshal([]byte(s), &o)

	kvList := strings.Split(s, "=")
	key := kvList[0]
	value := kvList[1]
	// golang strconv.ParseInt 是将字符串转换为数字的函数
	// 参数1 数字的字符串形式， 参数2 10进制 参数3 bit大小 也就是int64
	intValue, _ := strconv.ParseInt(value, 10, 64)
	fmt.Println(key, value)

	t := reflect.TypeOf(&o)
	v := reflect.ValueOf(&o)
	for i := 0; i < t.Elem().NumField(); i++ {
		// t.Elem().Field(i)  // 字段
		filed := t.Elem().Field(i)
		if filed.Tag.Get("json") == key {
			switch filed.Type.Kind() {
			case reflect.Int64:
				// 设置值
				v.Elem().Field(i).SetInt(intValue)
			}
			// 找到了那个字段
		}
	}

	fmt.Printf("o:%#v\n", o)
}

// loadStudentData 从文件中加载数据到Student结构体
func loadStudentData(filename string, obj interface{}) error {
	// 0. 前置检查
	// 0.1 obj必须是指针类型，不是就返回错误
	// 0.2 obj必须是结构体指针, 不是就返回错误
	// 1. 一行一行读取文件中的内容，解析成键值对 （读一行处理一行）
	// 2. 根据key去结构体里找字段名
	// 3. 找到之后赋值
	return nil
}
