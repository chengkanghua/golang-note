package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 小清单的增删改查

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
	// 2.1 获取当前用户的uid
	v, _ := c.Get(CtxUidKey)
	uid := v.(int64)
	if uid <= 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "登录异常,请重新登录",
		})
		return
	}
	todo.Uid = uid
	// 2.2 数据库插入一条记录
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

// updateTodoHandler 更新待办事项   {"id":6,"status":true}
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
	v, _ := c.Get(CtxUidKey)
	uid := v.(int64)
	if uid <= 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "登录异常,请重新登录",
		})
		return
	}
	// 先检查数据（todo.ID）是否存在
	// 拿请求传过来的id去数据库里查询是否存在这条记录
	if err := db.Where("id = ? and uid = ?", todo.ID, uid).First(&Todo{}).Error; err != nil {
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
	// 从ctx 获取uid
	v, _ := c.Get(CtxUidKey)
	uid := v.(int64)
	if uid <= 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "登录异常，请重新登录",
		})
		return
	}
	// 根据请求的uid查询对应用户的待办事项
	if err := db.Where("uid = ?", uid).Find(&todos).Error; err != nil {
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
	v, _ := c.Get(CtxUidKey)
	uid := v.(int64)
	if uid <= 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "登录异常，请重新登录",
		})
		return
	}
	// 2.1 先查一下有没有这条记录
	if err := db.Where("id = ? and uid = ?", id, uid).First(&Todo{}).Error; err != nil {
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
