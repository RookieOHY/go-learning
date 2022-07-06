package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin相关知识点：
	相关命令：
		强制下载（如果之前已经下载过了，执行更新）和安装gin包：go get -u github.com/gin-gonic/gin
		引入包：import "github.com/gin-gonic/gin"
	最基本的go服务：
		由gin和其他基本库来构建
*/
func main() {
	//gin.Default()函数默认会返回一个Engine指针。（路由的创建）
	router := gin.Default()
	//func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes（HandlerFunc是func(*Context)的新定义。因此，下方handleType函数的入参类型应该为*Context）
	router.GET("/", handleTypeMethod)
	//设置服务的运行端口（默认为8080，源码可参考utils.go下的resolveAddress函数）
	router.Run(":9200")
}

//入参*Context:表示Context的指针。而Context含义是一个上下文结构体（结构体是context包下context接口）。
//Request便是Context的一个成员（直接拿Request作为参数是否可行？本质上可行，但是如此设计的一些原因：可以使用到context包下context.go的一些现有的函数，如WithCancel, WithDeadline, WithTimeout, WithValue）
//如:context.WithDeadline(c,time.Now())
func handleTypeMethod(c *gin.Context) {
	//底层本质是接口Render负责字符串类型响应的渲染
	c.String(http.StatusOK, "hello gin!")
}
