package main

import "fmt"


type Person1 struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person1) SetDreams(dreams []string) {
	tmp := make([]string, len(dreams))
	copy(tmp, dreams)
	p.dreams = tmp
}

func demo14() {
	p1 := Person1{name: "jade", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉" // 修改切片变量,修改了底层数组

	fmt.Println(p1.dreams) // ?
}
