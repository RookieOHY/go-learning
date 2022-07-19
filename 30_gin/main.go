package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go-learning/30_gin/app/shop"
	"go-learning/30_gin/app/users"
	"go-learning/30_gin/router"
	"log"
	"net/http"
	"path"
	"time"
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
		文件上传：
			默认单个文件的大小为32<<20等于32M
		重定向：
			http重定向：c.Redirect()
			路由：r.HandleContext(c)
		中间件：
			在处理逻辑前面和后面的逻辑，都归为一个组件->中间件处理。
				如：鉴权、分页、业务耗时、记录日志。
		大型项目的路由拆分：
			简单的文件或者包
			按照业务拆分：如xx01.go,xx02.go。其中的路由注册和处理方法都写在一个xx.go文件中。
			更加细致的业务拆分：如shop目录下router.go 和 handler.go分别表示shop业务的路由注册和处理方法。
	本质上：使用gin,需要自己做一层封装
*/

type userInfo struct {
	//关于binding:required:表示这是一个必填字段，如果不传，或者为空，会报错
	Username string `json:"usernameJ" form:"usernameF" uri:"usernameU" binding:"required"`
	Password string `json:"passwordJ" form:"passwordF" uri:"passwordU" binding:"required"`
}

/*G.gin源码解读分析*/

/*F.gin+casbin实现权限管理*/

/*E.gin解析token*/

/*D.gin生成验证码*/

/*C.gin+air实现代码的实时加载*/

/*B.gin的日志文件输出控制台、文件*/

/*A.gin的参数校验*/
func main() {

}

/*⑫原生的http包操作session(多个http请求之间的关系)*/
//初始化一个存储cookie的对象
var cookieStore = sessions.NewCookieStore([]byte("00000000"))

func main12() {
	http.HandleFunc("/save", saveSession)
	http.HandleFunc("/get", getSession)
	http.HandleFunc("/delete", deleteSession)
	http.ListenAndServe(":9200", nil)
}

//删除session
func deleteSession(writer http.ResponseWriter, request *http.Request) {
	session, err := cookieStore.Get(request, "sessionName")
	if err != nil {
		log.Println("获取session失败")
		return
	}
	//删除：将session中数据的最大存储时间设置为小于0
	session.Options.MaxAge = -1
	session.Save(request, writer)
}

//获取session
func getSession(writer http.ResponseWriter, request *http.Request) {
	session, err := cookieStore.Get(request, "sessionName")
	if err != nil {
		log.Println("获取session失败")
		return
	}
	//取值
	name := session.Values["name"]
	log.Println("获取到的name为{}", name)
}

//设置session
func saveSession(writer http.ResponseWriter, request *http.Request) {
	session, err := cookieStore.Get(request, "sessionName")
	if err != nil {
		log.Println("生成session失败")
		return
	}
	//存数据
	session.Values["name"] = "RookieOHY"
	//保存
	session.Save(request, writer)
}

/*⑪基于cookie实现页面权限的控制*/
func main11() {
	r := gin.Default()
	//注册2个路由
	r.GET("/login", testLogin)
	r.GET("/home", AuthMiddleWare(), testHome)
	//端口
	r.Run(":9200")
}

//AuthMiddleWare 页面需要鉴权
func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, _ := context.Cookie("cookie")
		if cookie == "RookieOHY" {
			//放行，执行路由对应的方法
			context.Next()
			log.Println("有权限访问home")
			return
		}
		//不等于RookieOHY,报错：无权限访问home
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "无权访问home",
		})
		//不放行
		context.Abort()
		log.Println("无权限访问home")
	}
}

func testHome(context *gin.Context) {
	context.JSON(200, "this is home page")
}

func testLogin(context *gin.Context) {
	//设置cookie
	context.SetCookie("cookie", "RookieOHY", 60, "/",
		"localhost", false, true)
	// 返回信息
	context.JSON(200, "this is login page")
}

/*⑩gin操作cookie*/
func main10() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 服务端要给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})
	r.Run(":9200")
}

/*⑨gin的同步和异步*/
func main09() {
	r := gin.Default()
	//使用goroutine异步处理
	r.GET("/testAsync", func(c *gin.Context) {
		//不能使用原来的上下文对象，而是需要对应的副本
		cpCtx := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + cpCtx.Request.URL.Path)
		}()
	})
	//同步
	r.GET("/testSync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

	r.Run(":9200")
}

/*⑧gin的路由封装*/
func main08() {

	//1.所有的路由和方法都定义在一个routers.go文件汇总
	//route := router.SetUpRoute()
	//if route != nil {
	//	fmt.Printf("测试路由的封装~")
	//}

	//2.按照业务区分为一个一个.go文件
	//engine := gin.Default()
	//router.LoadUsers(engine)
	//router.LoadShop(engine)

	//3.一个业务对应一个目录
	// 加载多个APP的路由配置
	router.Include(shop.Routers, users.Routers)
	// 初始化路由
	r := router.Init()
	r.Run(":9200")
}

/*⑦gin的中间件*/
func main07() {
	//全局使用start
	//不使用默认的中间件Logger和Recovery
	route := gin.New()
	//route.Use(StatCost())
	//route.GET("/testGet", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"message":"test middleware successful",
	//	})
	//})
	//全局使用end
	//局部使用start
	//route.GET("/testGet2",StatCost(), func(c *gin.Context) {
	//	name:=c.MustGet("name")
	//	c.JSON(http.StatusOK,gin.H{
	//		"message":name,
	//	})
	//})
	//局部使用end

	//路由组注册中间件
	testGroup := route.Group("/testRouteGroup")
	testGroup.Use(StatCost())
	{
		testGroup.GET("/v1", func(c *gin.Context) {
			c.JSON(http.StatusOK, c.MustGet("name"))
		})
	}

	route.Run(":9200")
}

// StatCost 定义计算业务执行时间的中间件
func StatCost() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Set("name", "RookieOHY")
		//放行执行请求
		context.Next()
		//获取结束时间
		cost := time.Since(start)
		//打印耗时
		log.Println(cost)
	}
}

/*⑥gin路由和路由组*/
func main06() {
	route := gin.Default()
	route.Any("/testAny", testAny)
	//设置空路由
	route.NoRoute(testNoToute)
	//路由组
	v1 := route.Group("/v1")
	{
		v1.GET("/add", func(c *gin.Context) {
			//对应的业务
			c.JSON(http.StatusOK, gin.H{
				"message": "add success!",
			})
		})
		v1.GET("/add2", func(c *gin.Context) {
			//对应的业务
			c.JSON(http.StatusOK, gin.H{
				"message": "add2 success!",
			})
		})
		v1.GET("/add3", func(c *gin.Context) {
			//对应的业务
			c.JSON(http.StatusOK, gin.H{
				"message": "add3 success!",
			})
		})
	}
	//路由组嵌套
	userGroup := route.Group("/users")
	{
		userGroup.GET("/getById", nil)
		userGroup.GET("/update", nil)
		//嵌套
		run := userGroup.Group("/run")
		{
			run.GET("/withXx01", nil)
			run.GET("/withXx02", nil)

		}
	}

	route.Run(":9200")
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

/*⑤gin的请求转发和重定向*/
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

/*④gin的文件上传*/
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

/*③gin的参数绑定(绑定表单和绑定json)*/
func main03() {
	router := gin.Default()
	router.POST("/testBind", testBind)
	router.GET("/testBindUri/:usernameU/:passwordU", testBindUri)
	router.Run(":9200")
}

//测试参数的绑定2
func testBindUri(c *gin.Context) {
	var userInfo userInfo
	err := c.ShouldBindUri(&userInfo)
	if err == nil {
		c.JSON(http.StatusOK, userInfo)
	}
}

//测试参数的绑定
func testBind(c *gin.Context) {
	var userInfo userInfo
	//可以绑定任意类型
	//err := c.ShouldBind(&userInfo)
	//绑定json
	//err := c.ShouldBindJSON(&userInfo)
	//绑定form
	err := c.Bind(&userInfo)
	if err == nil {
		c.JSON(http.StatusOK, userInfo)
	}
}

/*②gin的post请求demo*/
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

/*①gin最基本的demo*/
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
	router.GET("/users/:uid", getUser)
	router.GET("/users", getUser2)
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
