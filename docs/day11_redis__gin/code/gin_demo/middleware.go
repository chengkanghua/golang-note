package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost(format string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/xyz" {
			c.Next() // -> c.Abort()
			return
		}

		start := time.Now()
		fmt.Println(start.Format(format))
		// 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Set("name", "小王子")
		// 调用该请求的剩余处理程序
		c.Next()

		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}
