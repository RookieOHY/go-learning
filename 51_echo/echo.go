package _1_echo

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

// user

type User struct {
	Name  string `json:"name" xml:"name" query:"name" form:"name"`
	Email string `json:"email" xml:"email" query:"email" form:"email"`
}

func EchoMain() {
	e := echo.New()
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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
	// 自定义选项
	//
	e.Start(":9999")
}
