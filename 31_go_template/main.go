package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
	go模板引擎的使用：
		主要为text/template、html/template
		参考李文周老师的博客：https://www.liwenzhou.com/posts/Go/go_template/
*/
type user struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", testTemplate)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Printf("http server start failed! err is %v\n", err)
		return
	}
}

func testTemplate(writer http.ResponseWriter, request *http.Request) {
	//读取模板和返回解析后数据
	file, err := template.ParseFiles("./go.tmpl")
	if err != nil {
		fmt.Printf(" Parse file failed! err:%v\n", err)
		return
	}
	m1 := map[string]interface{}{
		"name": "RookieOHY02",
		"age":  25,
	}
	u1 := user{Name: "RookieOHY", Age: 24}
	err2 := file.Execute(writer, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
	if err2 != nil {
		fmt.Printf(" render failed! err:%v\n", err2)
		return
	}
}
