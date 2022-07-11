package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

/*
gin相关知识点：
	相关命令：
		强制下载（如果之前已经下载过了，执行更新）和安装gin包：go get -u github.com/gin-gonic/gin
		引入包：import "github.com/gin-gonic/gin"
	最基本的go服务：
		由gin和其他基本库来构建
	一些方法：
		获取uri参数：
			Param("key"):获取api参数(占位符)
		获取get请求头参数：
			DefaultQuery("key","默认值"):获取?后面的参数值；获取不到或者没有传递key，设置值为默认值。
			Query("key"):获取?名字为key对应的值
			GetQuery("key")：效果同上，但是该方法存在2个返回值，前者为key对应的值，后者是一个布尔值。
		获取post请求的参数：
			PostForm: 同get
			DefaultPostForm：同get
			GetPostForm: 同get
		请求参数的对象绑定：
			ShouldBind：支持json、普通参数，表单参数的对象绑定
		重定向：
			http重定向：c.Redirect()
			路由：r.HandleContext(c)
	本质上：使用gin,需要自己做一层封装
*/

type userInfo struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

/*gin的中间件*/
func main06() {

}

/*gin路由和路由组*/
func main() {
	route := gin.Default()
	route.Any("/testAny", testAny)
	//设置空路由
	route.NoRoute(testNoToute)
	route.Run(":9200")
	//路由组 todo
	//路由组嵌套 todo
}

//404
func testNoToute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"messgae": "404",
	})
}

//测试
func testAny(c *gin.Context) {
	method := c.Request.Method
	c.JSON(http.StatusOK, gin.H{
		"message": "method type is" + method,
	})
}

/*gin的请求转发和重定向*/
func main05() {
	route := gin.Default()
	//http重定向
	route.GET("/rookieohy", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.rookieohy.top")
	})
	//路由重定向
	route.GET("/toTest2", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		route.HandleContext(c)
	})
	route.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "router is test2",
		})
	})
	route.Run(":9200")
}

/*gin的文件上传*/
func main04() {
	route := gin.Default()
	route.POST("/upload", uploadFile)
	route.POST("/multiUpload", uploadFiles)

	route.Run(":9200")
}

//测试文件的上传2
func uploadFiles(c *gin.Context) {
	//获取表单对象
	form, _ := c.MultipartForm()
	//获取多个文件对象
	files := form.File["files"]
	//遍历保存
	for _, file := range files { //取files中的所有图片
		dest := path.Join("C:/upload/", file.Filename) //保存路径
		c.SaveUploadedFile(file, dest)
	}
	c.JSON(200, gin.H{
		"message": "save files success",
	})

}

//测试文件的上传
func uploadFile(c *gin.Context) {
	//获取单个文件
	file, err := c.FormFile("simpleFile")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "upload file error",
		})
		return
	} else {
		//保存至某一个目录
		log.Println(file.Filename)
		//不会自动创建目录
		dir := fmt.Sprintf("C:/upload/%s", file.Filename)
		c.SaveUploadedFile(file, dir)
		c.JSON(http.StatusOK, gin.H{
			"message": "upload file success",
		})
	}
}

/*gin的参数绑定(绑定表单和绑定json)*/
func main03() {
	router := gin.Default()
	router.POST("/testBind", testBind)
	router.Run(":9200")
}

//测试参数的绑定
func testBind(c *gin.Context) {
	var userInfo userInfo
	err := c.ShouldBind(&userInfo)
	if err == nil {
		c.JSON(http.StatusOK, userInfo)
	}
}

/*post请求demo*/
func main02() {
	router := gin.Default()
	router.POST("/login", login)
	router.Run(":9200")
}

func login(c *gin.Context) {
	//方式1
	//username := c.PostForm("username")
	//password:=c.PostForm("password")
	//方式2(没有username和password参数名才会使用默认值)
	//username := c.DefaultPostForm("username", "我是username默认值")
	//password:=c.DefaultPostForm("password","我是password默认值")
	//方式3
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	data := map[string]interface{}{
		"name": username,
		"pwd":  password,
		"msg":  "success",
	}
	c.JSON(http.StatusOK, data)
}

/*get请求demo*/
func main01() {
	//gin.Default()函数默认会返回一个Engine指针。（路由的创建）
	router := gin.Default()

	//为模板设置自定义函数
	//router.SetFuncMap(template.FuncMap{
	//	"safe": func(str string) template.HTML{
	//		return template.HTML(str)
	//	},
	//})

	//func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes（HandlerFunc是func(*Context)的新定义。因此，下方handleType函数的入参类型应该为*Context）
	//参数：前者为请求映射的路径、后者为匿名函数或者为函数名
	router.GET("/", handleTypeMethod)
	router.GET("/user/:uid", getUser)
	router.GET("/user", getUser2)
	//加载一些静态的资源
	//router.LoadHTMLGlob("templates/**/*")
	//转发到某一模板
	//router.GET("/posts/index",handleIndex)
	//设置服务的运行端口（默认为8080，源码可参考utils.go下的resolveAddress函数）
	router.Run(":9200")
}

//响应静态资源
//func handleIndex(c *gin.Context) {
//	c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
//		"title":"<a href='http://www.rookieohy.top'>rookieohy的博客</a>",
//	})
//}

//模拟查询用户2
func getUser2(c *gin.Context) {
	name := c.DefaultQuery("name", "bird bro")
	//name := c.Query("name")
	//name, flag := c.GetQuery("name")
	//	//if !flag {
	//	//	name="查询不到"
	//	//	fmt.Println("query param [name] not exist!")
	//	//}
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

//定义用户(属性名字要求大写以暴露出去)
type user struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
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
