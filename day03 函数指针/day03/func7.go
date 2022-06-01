package main

import "fmt"

func f55() {
	// close()

	x := new(int) // new(类型)
	fmt.Println(x)
}

func f56() {
	defer func() {
		r := recover() // 恢复奔溃的现场
		fmt.Println("我很好~知道了~", r)
	}()
	var m map[string]int
	// panic("我要崩溃了...")
	m["杨俊"] = 100 // panic 程序崩了
	fmt.Println("美好的周末要结束了~")
}
