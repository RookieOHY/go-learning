package router

import "github.com/gin-gonic/gin"

func LoadUsers(e *gin.Engine) {
	//注册关于用户业务的路由
	e.POST("/register", register)
	e.POST("/login", login)

}

//对应的处理方法
func login(context *gin.Context) {
	context.JSON(200, "登录")

}

func register(context *gin.Context) {
	context.JSON(200, "注册")

}
