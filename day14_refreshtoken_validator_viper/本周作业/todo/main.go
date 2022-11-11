package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 小清单
var db *gorm.DB

// {"title":"杨俊不走神"}
type Todo struct {
	gorm.Model
	Title  string `form:"title" json:"title"`           // 待办事项名称
	Status bool   `json:"status"`                       // 是否完成的状态
	Uid    int64  `gorm:"uid;not null;default:0;index"` // 根据这一列能知道是谁的待办事项
}

// Account 用户表
type Account struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Uid int64 `gorm:"uid,unique"` // 用户id 唯一标识

	Name     string `gorm:"name,unique"` // 用户名
	Password string `gorm:"password"`

	NickName string `gorm:"nick_name"` // 昵称随便改
	Status   *bool  `gorm:"status"`
}

func initDB() (err error) {
	// dsn := "root:123@tcp(10.211.55.6:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
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
	db.AutoMigrate(&Account{})

	r := gin.Default()
	r.Use(AllowAll())

	// 加载静态文件
	r.LoadHTMLFiles("./index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/register", regHandler)
	r.POST("/login", loginHandler)
	r.POST("/logout", logoutHandler)

	// 注册路由，增删改查
	// 添加待办事项
	g := r.Group("/api/v1", authMiddleware) // 给路由组添加auth中间件
	{
		g.POST("/todo", createTodoHandler)
		g.PUT("/todo", updateTodoHandler)
		g.GET("/todo", getTodoHandler)           // ? authMiddleware -> getTodoHandler
		g.DELETE("/todo/:id", deleteTodoHandler) // 路由参数
	}

	// 启动http server端
	r.Run(":8999")
}

func AllowAll() gin.HandlerFunc {
	cfg := cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	cfg.AllowAllOrigins = true
	return cors.New(cfg)
}
