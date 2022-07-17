package shop

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/listGoods", listGoods)
	e.GET("/listStore", listStore)
}
