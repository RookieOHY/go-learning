package users

import "github.com/gin-gonic/gin"

func login(context *gin.Context) {
	context.JSON(200, "登录")

}

func listUser(context *gin.Context) {
	context.JSON(200, "用户列表")
}
