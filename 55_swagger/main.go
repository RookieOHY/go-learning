package main

import (
	"errors"
	_ "errors"
	"github.com/swaggo/swag/example/celler/httputil"
	"net/http"
	_ "net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/controller"
	_ "github.com/swaggo/swag/example/celler/docs"
	_ "github.com/swaggo/swag/example/celler/httputil"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			这是一个服务的api文档
//	@version		1.0
//	@description	文档描述
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API 开发者名字
//	@contact.url	https://subhee.top
//	@contact.email	204130199@qq.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
			accounts.POST(":id/images", c.UploadAccountImage)
		}
		bottles := v1.Group("/bottles")
		{
			bottles.GET(":id", c.ShowBottle)
			bottles.GET("", c.ListBottles)
		}
		admin := v1.Group("/admin")
		{
			admin.Use(auth())
			admin.POST("/auth", c.Auth)
		}
		examples := v1.Group("/examples")
		{
			examples.GET("ping", c.PingExample)
			examples.GET("calc", c.CalcExample)
			examples.GET("groups/:group_id/accounts/:account_id", c.PathParamsExample)
			examples.GET("header", c.HeaderExample)
			examples.GET("securities", c.SecuritiesExample)
			examples.GET("attribute", c.AttributeExample)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
