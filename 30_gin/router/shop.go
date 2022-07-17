package router

import "github.com/gin-gonic/gin"

func LoadShop(e *gin.Engine) {
	//注册路由
	e.GET("/addGoods", addGoods)
	e.GET("/pay", pay)
	e.GET("/sendMsg", sendMsg)
	e.GET("/receive", receive)
}

//对应的处理方法
func receive(context *gin.Context) {
	//收到商品
	context.JSON(200, "收到商品")
}

func sendMsg(context *gin.Context) {
	//订单短信
	context.JSON(200, "订单短信")
}

func pay(context *gin.Context) {
	//付款
	context.JSON(200, "付款")

}

func addGoods(context *gin.Context) {
	//加入购物车
	context.JSON(200, "加入购物车")

}
