package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin demo

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// r := gin.New()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {

		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.Any("/heiheihei", func(c *gin.Context) {

		// 根据请求的方法做处理
		c.JSON(http.StatusOK, gin.H{
			"method": c.Request.Method,
		})
	})

	r.GET("/user/search", func(c *gin.Context) {

		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 加载模板
	r.LoadHTMLFiles("./templates/upload.html")
	// 上传文件
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		// 接收用户上传文件的post请求
		f, err := c.FormFile("avatar")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		// 保存到指定目录 dst
		dst := "/Users/liwenzhou/Desktop/go/day11/code/gin_demo/upload/" + f.Filename
		c.SaveUploadedFile(f, dst)

		// 返回响应
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", f.Filename),
		})
	})

	// 重定向
	r.GET("/abc", func(c *gin.Context) {
		// c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
		c.Request.URL.Path = "/upload"
		r.HandleContext(c)
	})

	r.GET("/student", StatCost("2006-01-02"), studentHandler)
	r.GET("/school", schoolHandler)

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "没有这个网页")
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":8082")
}

func getStudentData() gin.H {
	return gin.H{
		"name": "王石头",
	}
}

func studentHandler(c *gin.Context) {
	// 1. 解析参数、校验参数
	// 2. 获取数据 拼装数据（业务逻辑）
	data := getStudentData()
	// 3. 返回响应
	c.JSON(http.StatusOK, data)
}

func schoolHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"school":  "xxxx",
		"student": getStudentData(),
	})
}
