package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	if err := initDB(); err != nil {
		fmt.Println("connect mysql failed ", err)
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLFiles("./login.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", loginHandler)
	r.Run(":8080")
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func loginHandler(c *gin.Context) {
	// 1. 从请求中获取用户的请求数据
	// 要么是form表单提交 要么是json格式提交
	var reqData Login
	if err := c.ShouldBind(&reqData); err != nil {
		fmt.Println(err)
		// 从请求里解析数据出错
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code": 1,
		// 	"msg":  "请求参数错误",
		// })
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求参数错误",
		})
		return
	}
	fmt.Printf("reqData:%#v\n", reqData)
	// c.JSON(http.StatusOK, reqData)
	fmt.Println("...")
	// 2. 对数据进行校验
	// if reqData.Username == "yangjun" && reqData.Password == "123456" {
	// 去MySQL数据库校验
	if u, err := QueryUser(reqData.Username, reqData.Password); err == nil {
		// 登陆成功
		fmt.Println(u)
		u.Desc = "test ..."
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "hello " + u.Username,
			"data": u, // !!!
		})
		return
	} else {
		// 登陆失败
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户名或密码错误",
		})
	}
	// 3. 返回响应
}

func initDB() (err error) {
	dsn := "root:123@tcp(10.211.55.6:3306)/db1?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect db failed ,err:%v\n", err)
		return
	}
	return
}

type User struct {
	Id       int    `db:"id" json:"-"`
	Username string `db:"username" json:"name"`
	Desc     string `json:",omitempty"`
}

func QueryUser(username, password string) (*User, error) {
	sqlStr := "select id,username from user where username=? and password=?"
	// 查库
	var u User
	err := db.Get(&u, sqlStr, username, password)
	if err != nil {
		fmt.Println(errors.Is(err, sql.ErrNoRows))
		fmt.Printf("get failed , err:%v \n", err)
		return nil, err
	}
	return &u, nil
}
