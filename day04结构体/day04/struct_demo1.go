package main

import "fmt"

// 结构体

// Student 定义一个学生类型
type Student struct {
	name     string
	age      int8
	married  bool
	mapScore map[string]int
}

// Order 定义一个订单类型
type Order struct {
	id         int64
	skuId      int64 // 商品id
	userId     int64
	createTime int64
}

func demo3() {
	// 声明一个Student类型的变量stu1
	var stu1 Student
	stu1.name = "杨俊" // 变量.属性（字段）
	stu1.age = 28
	// 结构体属性不赋值，默认就是属性对应类型的零值
	fmt.Printf("%+v\n", stu1)
	fmt.Printf("%#v\n", stu1)
	// 结构体中的map也需要初始化后才能使用
	stu1.mapScore = make(map[string]int, 2)
	stu1.mapScore["语文"] = 100
	stu1.mapScore["数学"] = 100
	stu1.mapScore["英语"] = 80
	fmt.Println(stu1.name)

	// 直接声明一个匿名结构体变量
	var user1 struct {
		name string
		age  int
	}
	fmt.Printf("user1:%#v\n", user1)

	// 在函数内部定义一个结构体类型User
	type User struct {
		name string
		age  int
	}
	var user2 User
	fmt.Printf("user2:%#v\n", user2)
	// 结构体指针
	var user3 = &user2 // user3存储的是指针，结构体指针 *User
	fmt.Printf("user3:%T\n", user3)

	p1 := new(int)                   // ? *int
	fmt.Printf("p1:%v %T\n", p1, p1) // 0xc000018120 *int
	p2 := new(User)                  // *User
	// %p 打印出指针类型的值
	fmt.Printf("p2:%p %T\n", p2, p2) // ?
}

// func demo4() {
// 	var user3 User
// }

// 结构体字面量初始化
func demo5() {
	stu2 := Student{
		name: "王俊翔",
		age:  26,
		mapScore: map[string]int{
			"语文": 6,
			"数学": 100,
		},
	}
	fmt.Printf("%+v\n", stu2)

	stu3 := Student{} // map[string]int{}
	fmt.Printf("%+v\n", stu3)

	stu4 := &Student{} // 取地址  --》 new(Student) --> 结构体指针
	(*stu4).name = "李硕"
	stu4.age = 18 // Go语言中提供的语法糖，支持 结构体指针类型.属性 简写
	fmt.Printf("%+v\n", stu4)

	// var stu5 *Student // nil
	// var stu5 = new(Student)
	var stu5 = &Student{}
	stu5.name = "jade" // (*nil).name =
	fmt.Printf("%+v\n", stu5)
	stu5 = &Student{
		name: "大都督",
	}
	stu5 = new(Student)

	// var x *int       // nil
	var x = new(int)
	*x = 100 // (*nil) = 100
	fmt.Println(x)

	// 列表初始化
	// 必须按结构体定义时候的属性顺序依次赋值
	var stu6 = Student{
		"吴勇",
		18,
		false,
		map[string]int{"语文": 1},
	}
	fmt.Println(stu6)
}
