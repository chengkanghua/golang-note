package main

import (
	"errors"
	"log"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 小清单
var db *gorm.DB

// Todo 待办事项结构体 <-> 对应的是数据库中的表
type Todo struct {
	gorm.Model
	Title  string `json:"title"`                // 待办事项名称
	Status bool   `form:"status" json:"status"` // 是否完成的状态
}

// APITodo 接口数据结构体
type APITodo struct {
	ID     uint   `form:"id" json:"id" gorm:"primarykey"`
	Title  string `form:"title" json:"title"`   // 待办事项名称
	Status bool   `form:"status" json:"status"` // 是否完成的状态
}

// initDB 初始化数据库
func initDB(dsn string) (err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

// ListHandler 查询列表
func ListHandler(c *gin.Context) {

	var todos []APITodo
	if err := db.Model(&Todo{}).Find(&todos).Error; err != nil {
		log.Println("getTodoHandler: 查询数据库失败", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "服务端异常，请稍后再试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": todos,
	})
}

// CreateHandler 创建事项
func CreateHandler(c *gin.Context) {
	var todo Todo
	var apiTodo APITodo
	if err := c.ShouldBind(&todo); err != nil {
		log.Println("invalid param failed:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "无效的参数",
		})
		return
	}
	if !(len(todo.Title) > 0) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "title不能为空",
		})
		return
	}

	// 创建一个代办事项
	if err := db.Create(&todo).Error; err != nil {
		log.Println("db.Create failed:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "服务端异常",
		})
		return
	}

	// 返回刚创建的代办事项
	err := db.Model(&todo).First(&apiTodo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "服务端异常",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"data": apiTodo,
	})
}

// UpdateHandler 更新事项
func UpdateHandler(c *gin.Context) {
	var todo Todo

	// 1.获取参数
	todoID := c.Param("id")

	// 2.校验id
	// 2.2 判断是否是数字
	for _, v := range todoID {
		if !unicode.IsDigit(v) {
			c.JSON(404, gin.H{
				"code": 1,
				"msg":  "404 not fount",
			})
			return
		}
	}

	// 2.2 判断是否存在
	err := db.Model(&Todo{}).
		First(&Todo{}, todoID).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有这条记录
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "请输入正确的id",
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

	// 3.绑定参数
	if err := c.ShouldBind(&todo); err != nil {
		log.Println("updateTodoHandler:获取请求参数失败", err)
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "无效的参数",
		})
		return
	}

	// 4.更新状态
	var todoObj Todo
	db.First(&todoObj) // todoObj.ID（有主键id）
	db.Model(&todoObj).Updates(map[string]interface{}{
		"title":  "打豆豆",
		"status": true,
	})

	if err := db.Model(&Todo{}).
		Where("ID = ?", todoID).
		Update("Status", todo.Status).
		Error; err != nil {
		log.Println("db.Update failed:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "服务端异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

// DestroyHandler 删除事项
func DestroyHandler(c *gin.Context) {
	// 1.获取参数
	todoID := c.Param("id")

	// 2.校验id
	// 2.2 判断是否是数字
	for _, v := range todoID {
		if !unicode.IsDigit(v) {
			c.JSON(404, gin.H{
				"code": 1,
				"msg":  "404 not fount",
			})
			return
		}
	}

	// 2.2 判断是否存在
	err := db.Model(&Todo{}).
		Find(&Todo{}, todoID).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有这条记录
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "请输入正确的id",
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

	if err := db.Delete(&Todo{}, todoID).Error; err != nil {
		log.Println("db.Delete failed:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "服务端异常",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"code": 0,
		"data": "success",
	})
}

func main() {
	// 连接数据库
	err := initDB("root:123456@tcp(192.168.1.2:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库连接失败!")
	}

	// gorm迁移
	db.AutoMigrate(&Todo{})

	r := gin.Default()
	g := r.Group("/api/v1")
	{
		g.GET("/todo", ListHandler)
		g.POST("/todo", CreateHandler)
		g.PUT("/todo/:id", UpdateHandler)
		g.DELETE("/todo/:id", DestroyHandler)
	}

	// 启动http server端
	r.Run(":18080")

}
