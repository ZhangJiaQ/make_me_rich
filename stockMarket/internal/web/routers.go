package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", homeHandler)
	// 在这里添加其他的路由和处理函数
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Stock Market Web Server!",
	})
}
