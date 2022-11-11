package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 小清单
var db *gorm.DB

// {"title":"杨俊不走神"}
type Todo struct {
	gorm.Model
	Title  string `form:"title" json:"title"` // 待办事项名称
	Status bool   `json:"status"`             // 是否完成的状态
}

func initDB() (err error) {
	dsn := "root:123@tcp(10.211.55.6:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// 初始化全局的db
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func main() {
	// 连接数据库
	if err := initDB(); err != nil {
		fmt.Println("connect mysql failed", err)
		panic(err)
	}
	// 表结构
	db.AutoMigrate(&Todo{})

	r := gin.Default()
	// 加载静态文件   前后端分离先注释
	// r.LoadHTMLFiles("./index.html")
	// r.Static("/static", "./static")

	// r.GET("/index", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", nil)
	// })

	// 注册路由，增删改查
	// 添加待办事项
	g := r.Group("/api/v1")
	{
		g.POST("/todo", createTodoHandler)
		g.PUT("/todo", updateTodoHandler)
		g.GET("/todo", getTodoHandler)
		g.DELETE("/todo/:id", deleteTodoHandler) // 路由参数
	}

	// 启动http server端
	r.Run(":8999")
}

// createTodoHandler 创建待办事项
func createTodoHandler(c *gin.Context) {
	// 1. 获取请求的参数
	var todo Todo
	if err := c.ShouldBind(&todo); err != nil {
		fmt.Println("invalid param", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "无效的参数", // 正常来说最好不要直接把后端错误返回给前端
		})
		return
	}
	// 2. 业务逻辑
	if err := db.Create(&todo).Error; err != nil {
		fmt.Println("db.Create failed", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常", // 正常来说最好不要直接把后端错误返回给前端
		})
		return
	}
	// 3. 返回响应
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		// "data": todo,
	})
}

// updateTodoHandler 更新待办事项
func updateTodoHandler(c *gin.Context) {
	// 1. 获取请求参数
	var todo Todo
	if err := c.ShouldBind(&todo); err != nil {
		fmt.Println("updateTodoHandler：获取请求参数失败", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "无效的参数",
		})
		return
	}
	// 2. 执行业务逻辑
	// 先检查数据（todo.ID）是否存在
	// 拿请求传过来的id去数据库里查询是否存在这条记录
	if err := db.First(&Todo{}, todo.ID).Error; err != nil {
		// if err == gorm.ErrRecordNotFound {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有这条记录
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "无效的参数",
			})
			return
		}
		// 其他错误
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	// 代码能执行到这里，说明数据库中确实存在 todo.ID 对应的记录
	// 接下来就去更新这条记录
	if err := db.Debug().Model(&todo).
		Update("status", todo.Status).
		Error; err != nil {
		// if err := db.Save(&todo).Error; err != nil {
		fmt.Println("updateTodoHandler：更新至数据库失败", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	// 3. 返回响应
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

// getTodoHandler 获取所有待办事项
func getTodoHandler(c *gin.Context) {
	// 1. 获取请求参数
	// 2. 执行业务逻辑
	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		fmt.Println("getTodoHandler: 查询数据库失败", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	// 3. 返回响应
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": todos,
	})
}

// deleteTodoHandler 删除待办事项
func deleteTodoHandler(c *gin.Context) {
	// 1. 获取请求参数
	// 127.0.0.1:8999/api/v1/todo/1
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("deleteTodoHandler：无效的参数", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "无效的参数",
		})
		return
	}
	// 2. 执行业务逻辑
	// 2.1 先查一下有没有这条记录
	if err := db.First(&Todo{}, id).Error; err != nil {
		// 没有这条记录的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "无效的参数",
			})
			return
		}
		fmt.Println("deleteTodoHandler：查询数据库失败", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	// 2.2 删除数据
	if err := db.Delete(&Todo{}, id).Error; err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}
	// 3. 返回响应
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
