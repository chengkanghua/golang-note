package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

var (
	NeedPtrErr       = errors.New("必须传入指针类型")
	NeedStructPtrErr = errors.New("必须传入结构体指针类型")
)

type Student struct {
	Name  string  `info:"name"`
	Age   int     `info:"age"`
	Score float64 `info:"score"`
}

func (s Student) Study(title string) {
	fmt.Printf("%s在学%s\n", s.Name, title)
}

func (s Student) Play(num int) {
	fmt.Printf("%s玩了%d小时\n", s.Name, num)
}

func Do(obj interface{}, methodName string, arg interface{}) {
	tInfo := reflect.TypeOf(obj)
	vInfo := reflect.ValueOf(obj)

	fmt.Println("方法数量：", tInfo.NumMethod()) // 看看结构体有多少方法
	// 根据传入方法名去结构体里找对应的方法
	m := vInfo.MethodByName(methodName)
	if !m.IsValid() || m.IsNil() {
		fmt.Println("调用方法失败")
		return
	}
	// 进行方法调用
	// 处理参数
	argValue := reflect.ValueOf(arg)
	m.Call([]reflect.Value{argValue})

}

// 从文件加载数据到Student结构体
func loadStudentData(filename string, obj interface{}) error {
	// 0，检查 obj interface{}
	// obj必须是指针类型
	// 如果不是，就返回错误
	tInfo := reflect.TypeOf(obj)
	vInfo := reflect.ValueOf(obj)
	// tInfo.Name()  // 类型名称
	// tInfo.Kind()  // 种类
	if tInfo.Kind() != reflect.Ptr {
		return NeedPtrErr
	}
	// obj必须是结构体指针
	// 如果不是，就返回错误
	// tInfo.Elem()  // 根据指针取到对应的值
	if tInfo.Elem().Kind() != reflect.Struct {
		return NeedStructPtrErr
	}

	// 1,一行行读取文件中的内容，解析键值对，读一行处理一行
	// 1.1 读文件
	b, err := ioutil.ReadFile(filename) // 文件不大的情况下
	if err != nil {
		fmt.Printf("打开文件失败，err：%v", err)
		return err
	}
	// []byte -> string
	s := string(b)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue // 跳过配置文件中的空行
		}
		// 按=分隔，拿到键值对
		list := strings.Split(line, "=") // [""]
		// 去掉首尾的空格
		key := strings.TrimSpace(list[0])
		value := strings.TrimSpace(list[1])
		// 2，根据key 去结构体里找字段的 info tag == key
		// 2.1 遍历结构体的所有字段
		for i := 0; i < tInfo.Elem().NumField(); i++ {
			field := tInfo.Elem().Field(i) // 拿到一个具体的结构体字段
			// field.Name == key
			if field.Tag.Get("info") == key {
				// 说明我们找到了对应配置文件中的那个字段
				// 3，找到后赋值
				vField := vInfo.Elem().Field(i)
				switch vField.Type().Kind() {
				case reflect.String:
					vField.SetString(value)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					// 注意：value是字符串
					// strconv.Atoi()  // string -> int
					int64Value, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return err
					}
					vField.SetInt(int64Value)
				case reflect.Float32, reflect.Float64:
					float64Value, err := strconv.ParseFloat(value, 64)
					if err != nil {
						return err
					}
					vField.SetFloat(float64Value)
				}

			}

		}
	}
	return nil

}

func main() {

	var stu Student
	err := loadStudentData("info.txt", &stu)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", stu)

	n, err := strconv.ParseInt("嘿嘿嘿", 10, 64) // 把“嘿嘿嘿”转换成int64
	if err != nil {                           // 转换失败
		fmt.Println(err)
	}
	// 转换失败n就是默认类型的零值
	fmt.Println(n)

	// 把字符串转换为int类型
	intValue, err := strconv.Atoi("123")
	fmt.Println(intValue, err)
	// 把int转为字符串
	v := 68 // 68 -> "68"

	// s := string(v)  // "D"  ascii码

	s := strconv.Itoa(v)

	fmt.Printf("%#v\n", s) // "68"

	s2 := fmt.Sprintf("%v", v) // "68"
	fmt.Printf("%#v\n", s2)

	// 通过反射调用方法
	Do(stu, "Study", "Go语言")
	Do(stu, "Play", 2)

}
