package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HelloHandler(c *gin.Context) {
	zap.L().Info("req hello")
	c.JSON(200, gin.H{"code": 0, "msg": "你好"})
}
