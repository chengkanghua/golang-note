package main

import "fmt"

/*
1. 使用“面向对象”的思维方式编写一个学生信息管理系统。
	1. 学生有id、姓名、年龄、分数等信息
	2. 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
*/

// obj := New管理系统()
// obj.StudentList()

// 1. 有学生的结构体
// 2. 管理系统能够存储学生信息
// 3. 管理系统得有查看学生列表、添加学生、编辑学生、删除学生的功能。

// Student 学生结构体类型
type Student struct {
	Id    int64
	Name  string
	Age   int8
	Score int8
}

// Manager 学生管理系统结构体
type Manager struct {
	StuInfo map[int64]*Student
}

// ShowAll 展示所有学生信息
func (m *Manager) ShowAll() {
	fmt.Println("展示所有学生信息")
	for id, stu := range m.StuInfo {
		fmt.Printf("id:%v name:%v age:%v score:%v\n",
			id, stu.Name, stu.Age, stu.Score)
	}
}

// AddStu 添加学生信息
func (m *Manager) AddStu() {
	fmt.Println("添加学生信息")
	// 用户输入id 姓名 年龄 分数
	fmt.Printf("请依次输入id 姓名 年龄 分数 并用空格分隔:")
	var (
		id    int64
		name  string
		age   int8
		score int8
	)
	fmt.Scanln(&id, &name, &age, &score)
	// 此处忽略可能因用户输入导致的错误

	// 判断id是否存在
	if _, ok := m.StuInfo[id]; ok {
		fmt.Println("学生信息已存在")
		return
	}

	// 创建新学生(可以使用构造函数 NewStudent)
	newStu := Student{
		Id:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
	// 将新增的学生纳入学员管理系统
	m.StuInfo[id] = &newStu
	fmt.Println("添加成功")
}

// EditStu 根据id编辑学生信息
func (m *Manager) EditStu() {
	fmt.Println("编辑学生信息")
	// 获取用户输入的id
	var id int64
	fmt.Print("请输入学生id:")
	fmt.Scanln(&id)
	if _, ok := m.StuInfo[id]; !ok {
		fmt.Println("查无此人")
		return
	}
	// 代码能走到这里说明有这个学生
	fmt.Print("请按格式依次输入 姓名 年龄 分数：")
	var (
		name  string
		age   int8
		score int8
	)
	fmt.Scanln(&name, &age, &score)
	// 忽略可能出现的输入错误
	m.StuInfo[id].Name = name
	m.StuInfo[id].Age = age
	m.StuInfo[id].Score = score
	fmt.Println("修改成功")

}

// DelStu 根据id删除学生信息
func (m *Manager) DelStu() {
	fmt.Println("删除学生信息")

	// 获取用户输入的id
	var id int64
	fmt.Print("请输入学生id:")
	fmt.Scanln(&id)
	if _, ok := m.StuInfo[id]; !ok {
		fmt.Println("查无此人")
		return
	}

	// 从map中删除key-value
	delete(m.StuInfo, id)
	fmt.Println("删除成功")
}

func main() {
	// 进入系统

	// 初始化Manager
	mgr := Manager{
		StuInfo: make(map[int64]*Student, 8),
	}

	for {
		// 打印菜单
		fmt.Print(`
		欢迎使用路飞学生信息管理系统v1.0
		菜单：
			1. 查看所有学生信息
			2. 增加学生信息
			3. 编辑学生信息
			4. 删除学生信息
			5. 退出

		`)

		// 获取用户输入
		var input int
		fmt.Print("请输入：")
		fmt.Scanln(&input) // 读取到非int数据会报错，input就是默认的0

		// 根据用户输入的菜单执行对应的方法
		switch input {
		case 1:
			// 调用管理者的方法
			mgr.ShowAll()
		case 2:
			mgr.AddStu()
		case 3:
			mgr.EditStu()
		case 4:
			mgr.DelStu()
		case 5:
			return
		default:
			fmt.Println("无效的输入")
		}
	}
}
