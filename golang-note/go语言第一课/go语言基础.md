

# 入门必备

## go语言特点

1 静态类型、编译型的开源语言

```go
package main

func main(){
	var num1 int =1
}

```



```go
package main

func main(){
  var num1 int =1
  num2 := 2  //编译器推倒变量类型 :=
}

```

2 脚本化的语法 ，支持多重编程范式

​	函数式 & 面向对象

3 原生，给力的并发编程支持



go语言的劣势

1 语法糖并没有Python和ruby 那么多

2 目前的程序运行速度还不及c

3 第三方函数库暂时不像绝对主流的编程语言那样多



## linux的安装方式

http://golang.org/dl/  下载最新版本的go语言二进制档案包

```
tar -zxf go1.4.2.linux-amd64.tar.gz -C /usr/local

设置4个环境变量： GOROOT  GOPATH. GOBIN   PATH
配置profile。（～/.bash_profile 或 /etc/profile)

export GOROOT = /usr/local/go
export GOPATH = ~/golib:~/goproject
export GOBIN = ～/gobin
export PATH=$PATH:$GOROOT/bin:$GOBIN

source /etc/profile

```



# 基本规则

常用概念和定义

​	工作区和GOPATH

```
			工作区是放置Go源码文件的目录
			 一般情况下，go源码文件都需要放在工作区中
			但是对于命令源码文件来说，这不是必须的

每一个工作区的结构都类似下图所示
	/home/hypermind/golib:
		src/       用于源码文件
		pkg/       用于存放归档文件  名称.a后缀名
		bin/       存放go程序可执行文件

GOOS GOARCH      如 linux_amd64

```





​	源码文件的分类和含义

```
名称.go 

命令源码文件、库源码文件
测试源码文件

命令源码文件
	声明自己属于main 代码包，包含无参数声明和结果声明的main函数
	被安装后，相应的可执行文件会被存放到GOBIN指向的目录或当前工作区目录 bin/下
	
库源码文件
	不具备命令源码文件的那两个特征的源码文件
	
测试源码文件
	不具备命令源码文件的那两个特征的源码文件


```



​	代码包的相关知识

```go
代码包的作用
	编译和归档Go程序的最基本单位


```





# 基础命令

```
go run main.go
	用于运行命令源码文件，
	
其内部操作步骤： 先编译源码文件再运行
	编译 --》临时目录 --》运行
	
-v : 列出被编译的代码包的名称
-work
-x
-n
https://golang.google.cn/cmd/go/


go build
	用于编译源码文件或代码包
	
go install 
	用于编译并安装代码包或源码文件
	
go get 
	用于从远程代码仓库上下载并安装代码包
	
	 

```

