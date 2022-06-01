package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthParam struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// loginHandler 登录处理函数
func loginHandler(c *gin.Context) {
	// 1. 获取参数/校验
	var param AuthParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "参数错误",
		})
		return
	}
	// 2. 逻辑处理
	// 拿用户名和密码去数据库查询，能查到记录说明登陆成功
	var u Account
	err := db.
		Where("name=? and password = ?", param.Name, md5Secret(param.Password)).
		First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, Resp{
				Code: 1,
				Msg:  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "服务端异常，请售后再试",
		})
		return
	}
	// 登陆成功，接下来该发token
	// 生成token
	token, err := GenToken(u.Uid, u.Name)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 1, Msg: "服务端异常，请售后再试"})
		return
	}
	// 2.2 存redis
	

	// 3. 返回响应

	c.JSON(http.StatusOK, Resp{
		Code: 0,
		Data: token,
	})
}

// regHandler 注册处理函数
func regHandler(c *gin.Context) {
	var param AuthParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "参数错误",
		})
		return
	}
	// 拿到参数去注册用户
	// 去数据库创建一条记录
	var user Account
	// 返回 gorm.ErrRecordNotFound 表示name没有被注册
	err := db.Where("name = ? ", param.Name).First(&user).Error
	if err == nil {
		c.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "用户名已存在",
		})
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "服务端异常，请稍后再试",
		})
		return
	}
	// 说明没有查到name并且也没报错
	// 终于可以去创建用户了
	err = db.Create(&Account{
		Uid:      time.Now().Unix(), // TODO: 后续会讲雪花算法生成
		Name:     param.Name,
		Password: md5Secret(param.Password),
	}).Error

	if err != nil {
		c.JSON(http.StatusOK, Resp{
			Code: 1,
			Msg:  "服务端异常，请稍后再试",
		})
		return
	}

	c.JSON(http.StatusOK, Resp{
		Code: 0,
		Msg:  "注册成功",
	})
	return
}

func md5Secret(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	return hex.EncodeToString(h.Sum([]byte(MySecret)))
}
