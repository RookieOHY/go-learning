package shop

import "github.com/gin-gonic/gin"

func listGoods(ctx *gin.Context) {
	ctx.JSON(200, "商品列表")
}
func listStore(ctx *gin.Context) {
	ctx.JSON(200, "门店列表")
}
