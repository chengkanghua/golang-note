package main

import (
	"fmt"
	"time"
)

// timeDemo 时间对象的年月日时分秒
func timeDemo() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
}

// 时间间隔
func durationDemo() {
	// time.Duration // 时间间隔类型  time.Duration 表示1纳秒

	time.Sleep(10 * time.Second) // sleep10秒

	num := 5
	// time.Sleep(num * time.Second) //不可以直接乘
	time.Sleep(time.Duration(num) * time.Second) // 5秒  time.Duration(num)表示把数字强制转换成纳秒
}

func timeOp() {
	now := time.Now()

	t2 := now.Add(time.Hour) // 1小时之后
	fmt.Println(t2.Sub(now)) // 60 * time.Minute   t2-当前时间=间隔时间1小时
}

func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

// formatDemo 时间格式化
func formatDemo() {
	now := time.Now()
	// 格式化的模板为 2006-01-02 15:04:05

	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制

	// fmt.Println(now.Format("Y.m.d H:M:S.000 PM Mon Jan"))
	fmt.Println(now.Format("2006.01.02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96

	// 只格式化时分秒部分
	fmt.Println(now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02"))
}

// parseDemo 指定时区解析时间
func parseDemo(s string) {
	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	timeObj, err := time.Parse("2006/01/02", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0000 UTC

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	timeObj, err = time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0800 CST
}

func main() {
	// timeDemo()
	// timeOp()
	// formatDemo()
	parseDemo("2022/10/07") // time.Parse("2022/10/7", "2006/01/02")
}
