package main

import (
	"fmt"
	"io"
	"os"
)

// 使用接口实现一个简单的日志库
// 既能往终端输出也能往文件输出日志

// log.Error()
// log.Warning()
// log.Info("...")

type Log struct {
	// Output os.Stdout // 标准输出
	Output io.Writer // 日志文件
}

// NewLog 构造函数中指定要输出的地方
// 需要传入一个ouput参数，告诉我将日志输出到哪里
func NewLog(output io.Writer) *Log {
	return &Log{
		Output: output,
	}
}

func (l *Log) Error(s string) {

}

func (l *Log) Warning(s string) {

}

func (l *Log) Info(s string) {
	// 把要记录的日志信息s 输出
	// 有可能会输出到 os.Stdout
	// 有可能出输出到 os.File
	// ....

	// 我只需要写东西...
	fmt.Fprintln(l.Output, s)

}

func main() {
	// 往终端输出日志
	// logger := NewLog(os.Stdout)
	// 往文件里面输出日志
	f, err := os.OpenFile("./app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("创建日志文件失败, err:", err)
		return
	}
	defer f.Close() // 程序退出前关闭文件

	logger := NewLog(f) // 把打开的文件对象传入构造函数

	logger.Info("程序启动啦")

	logger.Warning("感觉程序要出问题啦")
	logger.Error("程序出错误啦")
}

// 1. 日志构造函数中可以指定日志级别 level
// 2. error级别的日志还可以输出到另外一个文件 app.err.log
// 3. 日志文件可以限制大小，最大500M,支持自动切割日志文件
