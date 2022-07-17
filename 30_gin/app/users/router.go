package users

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/listUser", listUser)
	e.POST("/login", login)
}
