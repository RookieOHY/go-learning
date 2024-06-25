package _4_iris

import (
	"github.com/kataras/iris/v12"
	"regexp"
)

/*
获取app对象 iris.Default()
使用中间件 iris.Default().Use(中间件函数名字)

参数获取
	ctx.Params() 请求参数，配个Get、Get类型
*/

func BaseUsage() {
	app := iris.Default()
	app.Use(imw)

	// 自定义配置你的app 结构体
	//configuration := iris.WithConfiguration(iris.Configuration{
	//	DisableStartupLog: true,
	//	Charset:           "UTF-8",
	//})

	// 读取yaml toml
	//withConfiguration := iris.WithConfiguration(iris.YAML("./iris.yml"))
	//withConfiguration := iris.WithConfiguration(iris.TOML("./iris.tml"))

	// 传递多个配置项 每一个cfg项都有对应的with，可以调用多个后传入 with 和 without
	//cfg01 := iris.WithTimeFormat("Mon, 01 Jan 2006 15:04:05 GMT")

	// 传递你的自定义kv以及访问
	cfg02 := iris.WithOtherValue("appK", "appV")

	app.Handle("GET", "/cfg", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": app.ConfigurationReadOnly().GetOther()["appK"]})
	})

	// 参数获取 路径参数可以是{名字:类型} 之后利用对应类型的方法来获取路径参数
	app.Get("/testPath/{id:uint64}", func(context iris.Context) {
		path := context.Params().GetUint64Default("id", 0)
		context.JSON(iris.Map{"id": path})
	})

	// 参数获取 内建函数 对参数的值进行限制
	app.Get("/testPath02/{id:uint64 max(1024)}", func(context iris.Context) {
		path := context.Params().GetUint64Default("id", 0)
		context.JSON(iris.Map{"id": path})
	})

	// 参数获取 自定义宏、注册和使用
	// 比如：经纬度expr
	latLonExpr := "^-?[0-9]{1,3}(?:\\.[0-9]{1,10})?$"
	latLonRegex, _ := regexp.Compile(latLonExpr)

	// 注册宏 函数名 和 实现
	app.Macros().Get("string").RegisterFunc("jwdLint", latLonRegex.MatchString)

	// 使用
	app.Get("/jwdLint/{j:string jwdLint()}/{w:string jwdLint()}", func(context iris.Context) {
		context.JSON(iris.Map{"j": context.Params().Get("j"), "w": context.Params().Get("w")})
	})

	// 其他例子1
	app.Macros().Get("string").RegisterFunc("range",
		func(minLength, maxLength int) func(string) bool {
			return func(paramValue string) bool {
				return len(paramValue) >= minLength && len(paramValue) <= maxLength
			}
		})

	app.Get("/max/{name:string range(1,2) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s `, name)
	})

	// 其他例子2
	app.Macros().Get("string").RegisterFunc("has", func(arr []string) func(string) bool {
		return func(s string) bool {
			for _, element := range arr {
				if s == element {
					return true
				}
			}
			return false
		}
	})

	app.Get("/username/{name:string has([ohy,kiwi]) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s `, name)
	})

	// 中间件 func(上下文){}
	// 放行 ctx.next()

	// http错误处理 404 500
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	app.Run(iris.Addr(":9999"), cfg02)
}

func internalServerError(context iris.Context) {
	context.JSON(iris.Map{"code": iris.StatusNotFound})
}

func notFound(context iris.Context) {
	context.JSON(iris.Map{"code": iris.StatusInternalServerError})
}

func imw(ctx iris.Context) {
	ctx.Application().Logger().Infof("请求前执行before%s", ctx.Path())
	ctx.Next()
}
