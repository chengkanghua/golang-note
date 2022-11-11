package main

import (
	"fmt"

	"gomod_demo/app/blog"
	"gomod_demo/calc"
)

// 怎么样在项目里导入其他的包？
// 一个稍大的项目（project）内部又会分为很多的包（package）

// 额外补充两个问题：
// 1. calc下面能不能go mod init?
// 答：可以的，很多大的开源项目，里面的package也可以单独被拿出来引用
// 比如middleware这个包想支持被外部的人直接import，就执行下列命令：
// cd middleware
// go mod init gomod_demo/middleware
// 会在 gomod_demo/middleware 目录下生成 go.mod 文件

// 2. project下一个目录下面又分了多个目录

func main() {
	// 调用项目下面calc包中的Add函数
	ret := calc.Add(10, 20)
	fmt.Println(ret)

	// 调用项目下面app/blog/blog.go中的函数
	blog.Demo()
}
