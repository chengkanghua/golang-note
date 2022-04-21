package main

import (
	"fmt"
	"io/ioutil"
)

// bufio按行读取示例
func main() {
	context, err := ioutil.ReadFile("./src.txt")
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Print(string(context))
}
