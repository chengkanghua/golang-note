```go
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))  //defer在执行之前确定的变量值 x=1 y=2
	x = 10
	defer calc("BB", x, calc("B", x, y))  //defer在执行之前确定的变量值 x=10 y=2
	y = 20
}

-----------------
A 1 2 3
B 10 2 12
BB 10 12 22
AA 1 3 4

问题1: defer执行顺序是先定义的后执行,  先执行 BB AA   执行BB的时候就执行 B里的函数 在回道BB  然后AA 到A在回到AA
	我的想法应该是。B --》 BB --〉 A ———》 AA ，但实际情况确不一样， 这是什么原因呢

问题2: 第二个defer 再确定值的时候 y是第一次赋值的2 ，不是第二次赋值的20， 这和课堂的例子正好不一样？很疑惑？


--------------------------课堂例子
func sub(x, y int) int {
	return x - y
}

func deferDemo5() {
	x := 10
	defer func() {
		res := sub(x, sub(10, 2)) // x 100
		fmt.Println(res)
	}()
	x = 100
}

func main() {
	deferDemo5()    // 92
}

问题： 这里的x的值是第二次赋值的100 ，和上面的例子 y确是第一次的赋值2 ，why？



群内解答： 
问题1：defer不是要确定值才能执行的么，所以先A得出个结果，再B得出个结果，再BB，AA。
	defer 后面的参数如果是传了个函数调用的话需要先确定值，也就是先执行函数拿到结果

问题2: 第二个例子和第一个例子的区别是，第一个例子里面是传参，传的时候是什么值就是什么值，
	第二个例子里面是defer跟的是个匿名函数，这个匿名函数里面的变量x是使用的外部的变量



```

