package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// gin框架返回响应
// 想要写一个中间件把一些特定的请求的响应数据记录下来
// 张三请求了 /add 接口，我要把返回的响应数据记录到数据库中

type param struct {
	X int `json:"x" binding:"required"`
	Y int `json:"y" binding:"required"`
}

func addHandler(c *gin.Context) {
	fmt.Println("addHandler")
	var p param
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数有误",
			"err":  err.Error(),
		})
		return
	}
	res := p.X + p.Y
	c.JSON(200, gin.H{
		"code": 200,
		"data": res,
	})
}

type bodyLogWriter struct {
	gin.ResponseWriter               // 嵌入gin框架ResponseWriter
	body               *bytes.Buffer // 我们记录用的response
}

// Write 写入响应体数据
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // 我们记录一份
	return w.ResponseWriter.Write(b) // 真正写入响应数据
}

// recordResponseMiddleware 记录响应体的中间件
func recordResponseMiddleware(c *gin.Context) {
	fmt.Println("in recordResponseMiddleware")
	// 在gin框架写入响应数据前
	// 把 c.Writer 替换成我们自定义
	newWriter := &bodyLogWriter{
		body:           bytes.NewBuffer([]byte{}),
		ResponseWriter: c.Writer,
	}
	c.Writer = newWriter // 使用我们自定义的类型替换默认的
	c.Next()             // addHandler -> c.Writer c.Request.Response.Body
	// 记录addHandler返回的响应数据是什么
	// 1. 取到返回的响应数据
	// c.Request.Response.Body  // 响应体 io.ReadCloser接口类型
	statusCode := c.Writer.Status()
	if statusCode == 400 {
		b, _ := ioutil.ReadAll(newWriter.body)
		// gin框架需要把这个响应数据通过网络返回给请求客户端,读到io.EOF了
		// 你在这里把本应该返回给浏览器的响应数据给读完了就空指针了
		fmt.Printf("---> %s\n", b)
	}
	// 2. 存数据库（打印出来）
	fmt.Println("out recordResponseMiddleware")
}

func main() {
	r := gin.Default()
	r.POST("/add", recordResponseMiddleware, addHandler)
	r.Run(":8999")
}
