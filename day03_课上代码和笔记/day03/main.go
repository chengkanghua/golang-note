package main

func main() {
	// ex2()
	// ex3()
	// ex4()

	// f1()
	// sum(1, 2)
	// intSum(1, 2, 3)
	// f4()
	// someFunc("")

	// fmt.Println(num)
	// f11()
	// f12() // ?

	// f13()
	// f15()

	// f17(10, 20, add) // 把函数当成参数传进去

	// f := adder()
	// res := f() // 执行
	// fmt.Println(res)
	// res = f()        // 执行
	// fmt.Println(res) // ?

	// f2只要在，那么那就一直会占着adder2里面的x这个变量
	// f2 := adder2()  // 拿到返回的匿名函数
	// r2 := f2(10)    // 执行匿名函数
	// fmt.Println(r2) // 0+10=10
	// r2 = f2(100)    // 执行匿名函数
	// fmt.Println(r2) // 10+100=110

	// f3 := adder3(10)
	// f4 := adder3(100)

	// r3 := f3(1)
	// r4 := f4(1)
	// fmt.Println(r3, r4)

	// filename := "xxx"
	// jpgFunc := makeSuffixFunc(".jpg")
	// res := jpgFunc(filename) // 如果不是.jpg结尾的就添加一个.jpg后缀
	// fmt.Println(res)

	// txtFunc := makeSuffixFunc(".txt")
	// res = txtFunc(filename) // 如果不是.txt结尾的就添加一个.jpg后缀
	// fmt.Println(res)

	// fmt.Println(deferDemo1())
	// fmt.Println(deferDemo2())
	// fmt.Println(deferDemo3())
	// fmt.Println(deferDemo4())
	// deferDemo5()

	// defer func() {
	// 	r := recover()
	// 	if r != nil {
	// 		fmt.Println("main:recover", r)
	// 		return
	// 	}
	// 	defer sub(10, 2)
	// 	fmt.Println("一切正常")
	// }()

	// f56()

	// f66()

	// x := 1
	// f67(&x) // Go语言中函数传参全都是值拷贝
	// fmt.Println(x)

	f68()
}
