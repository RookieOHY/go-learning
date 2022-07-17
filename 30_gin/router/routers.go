package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetUpRoute 设置路由和返回
func SetUpRoute() *gin.Engine {
	engine := gin.Default()
	engine.GET("/get", testRoute)
	return engine
}

func testRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "测试路由封装",
	})
}
