package main

func main() {
	// var x [3]int8 // 索引范围：0 1 2
	// fmt.Println(x)
	// fmt.Println(x[0])
	// x[0] = 10
	// fmt.Println(x[0])
	// fmt.Println(x[1])
	// x[2] = -11
	// fmt.Println(x[2])

	// fmt.Printf("%T \n", x) // 格式化打印
	// fmt.Println(x)         // 直接打印

	// var y [2]string
	// fmt.Println(y)
	// fmt.Printf("%#v\n", y)
	// fmt.Printf("%+v\n", y)

	// var c [2]bool
	// c[1] = true
	// fmt.Println(c)
	// // fmt.Println(c[2])
	// // fmt.Println(c[-1])       // 不支持负数索引
	// fmt.Println(c[len(c)-1]) // 使用len(array)-1取最后一个元素

	// var s = "安小枫"
	// fmt.Println(len(s)) // 内置len函数支持获取字符串的长度等等
	// fmt.Println(len(c)) // 内置len函数支持获取切片的长度等等

	// // 数组的初始化
	// var x1 = [3]int{1, 0, -1}
	// fmt.Println(x1)
	// // 使用索引赋值进行初始化
	// var x99 = [100]int{98: -2, 99: 1}
	// fmt.Println(x99)
	// // 事先不知道要有几个元素
	// var yy = [...]string{"安小枫", "jade", "大都督"} // 编译器自动的去帮我数一下有几个元素
	// fmt.Println(yy, len(yy))

	// var z = [...]int{1: 1, 3: 3, 7: 7}
	// fmt.Println(z, len(z))

	// const size = 100 // 定义常量
	// // var size = 100
	// var xx = [size]int{} // 支持使用常量作为数组的长度
	// fmt.Println(xx)

	// 数组的遍历
	// var zz = [...]int{1, 3, 5, 7, 9}
	// // 1. for i索引遍历
	// for i := 0; i < len(zz); i++ { // i的范围：0，1，2，3，4
	// 	fmt.Println(zz[i]) // 根据索引依次取出数组中的每一个元素
	// }

	// // 2. for range 遍历
	// for i := range zz { // 从zz中依次取出每一个元素赋值给变量i
	// 	fmt.Println(i)
	// }

	// for i, v := range zz { // 从zz中依次取出索引和值，分别赋值给i和v
	// 	fmt.Println(i, v)
	// }

	// for _, v := range zz { // 利用特殊的 _ ，把索引丢弃
	// 	fmt.Println(v)
	// }

	// f1()
	// f2()
	// arraySum()
	// arraySum2()

	// slice1()
	// slice2()
	// slice3()

	// slice4()
	// make1()

	// slice5()
	// slice6()
	// copyDemo()
	// copyDemo2()

	// slice7()

	// appendDemo()
	// appendDemo2()

	// deleteSlice(1)
	// appendDemo4()

	// mapDemo1()
	mapDemo2()
}
