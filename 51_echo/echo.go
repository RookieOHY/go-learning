package _1_echo

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

// user

type User struct {
	Name  string `json:"name" xml:"name" query:"name" form:"name" validate:"required"`
	Email string `json:"email" xml:"email" query:"email" form:"email" validate:"required,email"`
}

// MyCtx 结构体持有echo.Context 并且实现2个函数
type MyCtx struct {
	echo.Context
}

func (receiver *MyCtx) foo() {
	println("foo方法执行了")
}

func (receiver *MyCtx) bar() {
	println("bar方法执行了")
}

type (
	RequestValidator struct {
		// 最好申明 指针类型
		validator *validator.Validate
	}
)

func (receiver *RequestValidator) Validate(i interface{}) error {
	return receiver.validator.Struct(i)
}

func EchoMain() {
	e := echo.New()

	// https://github.com/go-playground/validator 借助它实现数据校验,定义结构体持有 validator并且实现方法
	// echo使用定义好的Validator
	e.Validator = &RequestValidator{validator: validator.New()}
	e.POST("/valid", func(context echo.Context) error {
		u := new(User)
		err := context.Bind(u)
		if err != nil {
			return err
		}
		//调用僬侥结构体的校验函数 执行校验
		err = context.Validate(u)
		if err != nil {
			return err
		}
		// 返回
		return context.JSON(200, u)
	})

	e.GET("/echo", func(context echo.Context) error {
		return context.String(200, "echo返回内容")
	})
	// 获取url里面的参数 如：/user/1 /user/ohy 等
	e.GET("/user/:id", func(context echo.Context) error {
		id := context.Param("id")
		return context.String(http.StatusOK, "获取的id："+id)
	})
	// 获取?k=v&k=v
	e.GET("/user", func(context echo.Context) error {
		name := context.QueryParam("name")
		age := context.QueryParam("age")
		return context.String(http.StatusOK, "name："+name+",age:"+age)
	})
	// 获取表单 application/x-www-form-urlencoded
	// 获取表单 multipart/form-data
	e.POST("/form", func(context echo.Context) error {
		id := context.FormValue("id")
		value := context.FormValue("value")
		return context.String(http.StatusOK, "id:"+id+",value:"+value)
	})

	// json 绑定对象
	e.POST("/user/save", func(context echo.Context) error {
		result := new(User)
		err := context.Bind(result)
		if err != nil {
			return err
		}
		return context.JSON(http.StatusCreated, result)
	})

	// 中间件：root group route 三种
	// root 级别
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// group todo
	g := e.Group("/mw")
	g.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == "ohy" && password == "123456" {
			return true, nil
		}
		return false, nil
	}))

	// router 级别
	mw := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			println("req is /mw")
			return next(context)
		}
	}

	e.GET("/admin/mw", func(context echo.Context) error {
		return context.String(http.StatusOK, "成功")
	}, mw)
	// 上下文 echo.Context 表示http请求的上下文
	// Context 是接口 可以自己实现 或者 定义结构体持有
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			// 新建 MyCtx 返回
			ctx := &MyCtx{context}
			return handlerFunc(ctx)
		}
	})
	// 接收使用定义的中间件
	e.GET("/ctx", func(context echo.Context) error {
		// 这里直接断言
		ctx := context.(*MyCtx)
		// 调用ctx定义的方法
		ctx.foo()
		ctx.bar()
		// 处理请求
		return ctx.String(http.StatusOK, "OK")
	})

	// 配合http.Cookie实现读写cookie
	e.GET("/createCookie", func(context echo.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = "RookieOHY"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		context.SetCookie(cookie)
		return context.String(http.StatusOK, "create cookie ok")
	})

	e.GET("/readCookie", func(context echo.Context) error {
		cookie, err := context.Cookie("username")
		if err != nil {
			return err
		}
		println(cookie.Name)
		println(cookie.Value)

		// 获取所有的cookie
		cookies := context.Cookies()
		for _, c := range cookies {
			println("c->" + c.Value)
		}

		return context.String(http.StatusOK, cookie.Value)
	})

	// 响应总结：字符串、json、html、xml、file、blob
	e.GET("/respType", func(context echo.Context) error {
		// 字符串
		//err := context.String(200, "string")
		// html
		//err := context.HTML(200,"<h1>string</h1>")
		// json 或者 美化的json
		user := &User{
			Name:  "returnName",
			Email: "returnName@qq.com",
		}
		//err := context.JSON(200,user)

		// json流 适用于大对象返回 效率高
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		context.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(context.Response()).Encode(user)
	})

	// Hook 响应前和后执行的逻辑
	e.GET("/hook", func(context echo.Context) error {
		context.Response().Before(func() {
			println("请求前打印...")
		})
		context.Response().After(func() {
			println("请求后打印...")
		})
		return context.String(200, "hook")
	})

	// 路由对象：可以设置路由的名字
	// 路由组：具有相同前缀的可以定位一组路由
	group := e.Group("/group")
	group.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == "ohy" && password == "123456" {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/group/test", func(context echo.Context) error {
		return context.String(200, "group ok")
	})

	e.Start(":9999")
}
