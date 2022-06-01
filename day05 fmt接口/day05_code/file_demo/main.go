package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

func main() {

	bufio_read()
}

func file_read() {
	f, err := os.Open("./main.go") // 默认以读的方式打开文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // 调用f.Close()是安全的

	// 读取文件内容
	var data [4096]byte

	n, err := f.Read(data[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", data[:n])

	// f.Close() // 关闭文件
}

func bufio_read() {
	f, err := os.Open("./main.go") // 默认以读的方式打开文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // 调用f.Close()是安全的

	reader := bufio.NewReader(f)
	for {
		s, err := reader.ReadString('\n') // 每次读取到换行符 一行一行的读
		if err == io.EOF {
			// 读完了
			fmt.Print(s)
			return
		}
		fmt.Print(s)
	}
}

func ioutil_read() {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func file_write() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	str := "hello 沙河"
	// fmt.Fprint()
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("hello 小王子") //直接写入字符串数据
}

func bufio_write() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

func ioutil_write() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
