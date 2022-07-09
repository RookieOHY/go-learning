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
	一些方法：
		获取参数：
			Param("key"):获取api参数
			DefaultQuery("key","默认值"):获取?后面的参数值；获取不到或者没有传递，设置值为默认值。
*/
func main() {
	//gin.Default()函数默认会返回一个Engine指针。（路由的创建）
	router := gin.Default()
	//func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes（HandlerFunc是func(*Context)的新定义。因此，下方handleType函数的入参类型应该为*Context）
	//参数：前者为请求映射的路径、后者为匿名函数或者为函数名
	router.GET("/", handleTypeMethod)
	router.GET("/user/:uid", getUser)
	router.GET("/user", getUser2)
	//设置服务的运行端口（默认为8080，源码可参考utils.go下的resolveAddress函数）
	router.Run(":9200")
}

//模拟查询用户2
func getUser2(c *gin.Context) {
	name := c.DefaultQuery("name", "bird bro")
	u2 := user{
		Id:   "2",
		Name: "RookieOHY02",
		Age:  25,
	}
	if name == u2.Name {
		c.JSON(http.StatusOK, u2)
	} else {
		c.String(http.StatusOK, "你好! "+name)
	}
}

//定义用户
type user struct {
	Id   string
	Name string
	Age  int
}

//模拟查询用户
func getUser(c *gin.Context) {
	//本质是遍历请求的每一个参数名字，如和uid匹配，返回对应key的value
	uid := c.Param("uid")
	u := user{
		Id:   "1",
		Name: "RookieOHY",
		Age:  25,
	}
	if uid == u.Id {
		c.JSON(http.StatusOK, u)
	} else {
		//gin.H 为新定义的类型。本质是：key为string类型、v为任意类型（空接口）的一个map
		c.JSON(http.StatusNotFound, gin.H{
			"message": uid + "对应的用户未注册~",
		})
	}
}

//入参*Context:表示Context的指针。而Context含义是一个上下文结构体（结构体是context包下context接口）。
//Request便是Context的一个成员（直接拿Request作为参数是否可行？本质上可行，但是如此设计的一些原因：可以使用到context包下context.go的一些现有的函数，如WithCancel, WithDeadline, WithTimeout, WithValue）
//如:context.WithDeadline(c,time.Now())
func handleTypeMethod(c *gin.Context) {
	//底层本质是接口Render负责字符串类型响应的渲染
	c.String(http.StatusOK, "hello gin!")
}
