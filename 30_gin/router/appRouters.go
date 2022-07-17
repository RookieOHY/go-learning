package router

import "github.com/gin-gonic/gin"

//1.将func(engine *gin.Engine) 定义为一种新类型
type Option func(engine *gin.Engine)

//2.定义Option
var options = []Option{}

//3.数组里面添加内容
func Include(opts ...Option) {
	options = append(options, opts...)
}

//4.初始化路由
func Init() *gin.Engine {
	engine := gin.Default()
	for _, option := range options {
		//每一个新定义的函数类型Option(本质是以入参为engine的可变参数的函数)
		option(engine)
	}
	return engine
}
