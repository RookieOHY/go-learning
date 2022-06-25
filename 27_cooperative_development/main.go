package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
协作开发二三事：
	使用模块化开发提高开发效率.
		一个项目=多个模块=多个包
	新建一个模块或者项目：
		go mod init 文件夹名字（执行成功，生成一个go.mod）
	引用第三方模块（如：gin）
		先安装（下载）、再引用
			go get -u github.com/gin-gonic/gin
			import即可
		(如果使用的不是goland，可能需要使用go mod tidy将第三方的模块添加go.mod文件缺失的包)
		(Go 语言的工具链比如 go mod tidy 命令，就可以帮助我们自动地维护、自动地添加或者修改 go.mod 的内容)
	建议：模块名字最好是域名/模块名字
	包概念拾遗：
		若干个.go文件可以组成一个包。
		包和包之间可以互相应用。
		包下的任意一个.go文件都必须申明属于哪一个包名字。
		私有和共有体现作用域：go文件成员的名字是否大写。（大写，任意包下文件都可以引用该成员。否则，只能在该包所引用）
		分类：
			main包，编译器会主动找到main包，把main包编译成可执行文件，所有的服务，都必须有一个main包。

*/
func main01() {
	fmt.Println("我执行了")
}
func init() {
	fmt.Println("我比main先执行了")
}

/*引用第三方模块：gin*/
func main() {
	fmt.Println("先导入fmt包，才能使用")
	r := gin.Default()
	r.Run()
}
