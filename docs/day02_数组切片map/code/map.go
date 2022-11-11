package main

import "fmt"

// map

func mapDemo1() {
	var m map[string]int

	fmt.Println(m == nil)    // true
	m = make(map[string]int) // 只要初始化就可以赋值
	fmt.Println(m == nil)    // false

	m["jade"] = 300 // 设置值
	fmt.Println(m)
	weight := m["jade"] // 取值
	fmt.Println(weight)

	// 字面量初始化
	m2 := map[string]string{
		"username": "七米",
		"password": "1234", // 花括号换行则此处必须加逗号
	}
	fmt.Printf("%#v\n", m2)

	name := m2["name"]
	fmt.Println(name)
	// v, ok 取值  ; ok是一个变量名，只不过大家约定成俗在这里用ok
	// ok = true表示map中有这个key， ok=false表示map没有这个key
	// 如果没有这个key，此时v=对应类型的零值
	v, ok := m2["name"] // 类似与for range 可以用一个变量也可以两个变量接收
	fmt.Println(v, ok)

	_, ok = m2["name"]
	fmt.Println(ok)
}

// map的遍历（map是无序的）
func mapDemo2() {
	m := map[string]int{
		"jade": 300,
		"ddd":  180,
		"嚯嚯嚯":  160,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// 只取map中的key
	for k := range m {
		fmt.Println(k)
	}

	// 只取map中的value
	for _, v := range m {
		fmt.Println(v)
	}

	// 判断map中是否存在某个key
	_, ok := m["小盆子"]
	fmt.Println(ok)

	// 从map中删除键值对
	delete(m, "jade") // 没有返回值
	fmt.Println(m)
}
