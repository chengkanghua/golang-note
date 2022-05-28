package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 中间件

// context.Context
// type ctxKey string

// var (
// 	KeyUid  ctxKey = "uid"
// 	KeyName ctxKey = "name"
// )

// authMiddleware 从请求头获取 token 完成校验
func authMiddleware(c *gin.Context) {
	// 1. 从请求头获取token
	authHeader := c.Request.Header.Get("Authorization")
	// authHeader := c.Request.Header.Get("todo-app")
	fmt.Println(authHeader)

	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求头中Bearer Token为空",
		})
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请求头中Bearer Auth格式有误",
		})
		c.Abort()
		return
	}

	// 去黑名单查找是否有，有就提示token过期
	for index, value := range clice_tokens {
		fmt.Println(index, value)
		if value == parts[1] {
			c.JSON(http.StatusOK, gin.H{
				"code": 404,
				"msg":  "token已过期",
			})
			c.Abort()
			return
		}
	}

	// 解析token
	mc, err := ParseToken(parts[1])
	fmt.Println(mc)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "无效的token",
		})
		c.Abort()
		return
	}
	// 去redis根据userid查token判断是否一致

	// 将uid和name数据写入ctx
	c.Set(CtxNameKey, mc.Name)
	c.Set(CtxUidKey, mc.Uid)

	// c.Next() // 在这里去执行下一个函数
	// 2. 解析token
	// 3. 校验token
}
