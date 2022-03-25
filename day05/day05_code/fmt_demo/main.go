package main

import "fmt"

// fmt

type Student struct {
	Name string
	Age  int
}

func main() {
	// fmt.Print("你输入什么就打印什么")
	// fmt.Println()
	// fmt.Printf("id:%v", 101)

	stu := Student{
		Name: "杨俊",
		Age:  28,
	}

	// fmt.Println(stu)               // {杨俊 28}
	// fmt.Printf("%v\n", stu)        // {杨俊 28}
	// fmt.Printf("%+v\n", stu)       // {Name:杨俊 Age:28}
	// fmt.Printf("%#v\n", stu)       // main.Student{Name:"杨俊", Age:28}
	// fmt.Printf("%% %v\n", stu.Age) // % 28

	// // %s
	// var b = []byte{'h', 'e', 'l', 'l', 'o'}
	// fmt.Println(b)
	// fmt.Printf("%s \n", b)

	// bs, err := json.Marshal(stu)
	// fmt.Printf("bs:%s err:%v\n", bs, err)

	// fmt.Println(&stu)
	// fmt.Printf("%p \n", &stu) // 0xc00000c030

	// // 宽度标识符
	// n := 12.3456
	// fmt.Printf("%f\n", n)    // 12.345600
	// fmt.Printf("%9f\n", n)   // 12.345600
	// fmt.Printf("%.2f\n", n)  // 12.35
	// fmt.Printf("%9.2f\n", n) //     12.35
	// fmt.Printf("%9.f\n", n)  //        12

	// f1 := 12.3                // 12.30
	// fmt.Printf("%.2f \n", f1) // 12.30

	// s := "杨俊"
	// fmt.Printf("%s\n", s)
	// fmt.Printf("%5s\n", s)
	// fmt.Printf("%-5s\n", s)
	// fmt.Printf("%5.7s\n", s)
	// fmt.Printf("%-5.7s\n", s)
	// fmt.Printf("%5.2s\n", s)
	// fmt.Printf("%05s\n", s)

	// f, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Fprintln(os.Stdout, "hello") // fmt.Println("hello")

	// str := "nb"
	// fmt.Fprintf(f, "你接下来要说的是：%s", str) // fmt.Println("hello")

	// 学生姓名：年龄：
	// s2 := "学生姓名" + stu.Name + "年龄：" + stu.Age
	s2 := fmt.Sprintf("学生姓名：%s 年龄：%d", stu.Name, stu.Age)
	fmt.Println(s2)

	scan_demo()

}

func scan_demo() {
	// scan

	var (
		name string
		age  int
	)
	fmt.Scanf("姓名:%s 年龄:%d", &name, &age)
	fmt.Println(name, age)
}
