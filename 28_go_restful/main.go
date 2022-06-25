package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

/*
Go+restful风格来开发
*/
/*
①go自带的net实现查询用户列表
*/
func main01() {
	http.HandleFunc("/users", handleUsersJSON)
	http.ListenAndServe(":8080", nil)
}
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ID:1,NAME:RookieOHY")
		fmt.Fprintln(w, "ID:2,NAME:RookieOHY02")
		fmt.Fprintln(w, "ID:3,NAME:RookieOHY03")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not data")
	}
}

//返回json
var users = []User{
	{ID: 1, Name: "R1"},
	{ID: 2, Name: "R2"},
	{ID: 3, Name: "R4"},
}

func handleUsersJSON(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}

type User struct {
	ID   int
	Name string
}

/*-----------------------------------*/
/*
 而使用Gin实现CRUD
*/
func main() {
	r := gin.Default()
	//相对路径和方法名字
	r.GET("/users", listUser)
	r.GET("/users/:id", getUser)
	r.POST("/users", addUser)
	r.DELETE("/users/:id", deleteUser)
	//设置端口
	r.Run(":8080")
}

//删除用户
func deleteUser(context *gin.Context) {
	id := context.Param("id")
	i := -1
	//找到下标index
	for index, dItem := range users {
		if strings.EqualFold(id, strconv.Itoa(dItem.ID)) {
			//移除
			i = index
			break
		}
	}
	//执行删除
	if i >= 0 {
		users = append(users[:i], users[i+1:]...)
		fmt.Println(i)
		fmt.Println(users[:i])
		fmt.Println(users[i+1:])
		context.JSON(http.StatusNoContent, "")
	} else {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}
}

//新增用户
func addUser(context *gin.Context) {
	//从表单的实体中获取Name
	name := context.DefaultPostForm("name", "")
	if name != "" {
		addUser := User{
			ID:   len(users) + 1,
			Name: name,
		}
		//追加
		users = append(users, addUser)
		//json
		context.JSON(http.StatusCreated, addUser)
	} else {
		context.JSON(200, gin.H{
			"message": "请指定名字",
		})
	}
}

//根据id查询
func getUser(context *gin.Context) {
	//获取请求中uri的占位符id
	id := context.Param("id")
	//申明未初始化的User
	var user User
	//申明标志位（表示查询结果）
	flag := false
	//模拟查询数据库（这里遍历users匹配即可）
	for _, userItem := range users {
		//id是否一样
		//参数：来源id 现有的id
		if strings.EqualFold(id, strconv.Itoa(userItem.ID)) {
			user = userItem
			flag = true
			break
		}
	}
	if flag {
		context.JSON(200, user)
	} else {
		context.JSON(404, gin.H{
			"message": "用户不存在",
		})
	}
}

//查询列表
func listUser(context *gin.Context) {
	context.JSON(200, users)
}
