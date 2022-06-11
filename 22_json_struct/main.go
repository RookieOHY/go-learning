package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*json 和 struct之间的互相转换*/
/*
①利用json包Marshal函数的，将结构体对象转换为json
②②利用struct tag 标记结构体的字段
*/
func main01() {
	p := person{
		Name: "RookieOHY",
		Age:  25,
	}
	//返回[]byte
	result, err := json.Marshal(p)
	if err == nil {
		fmt.Println("result is :", string(result))
	}
	//相反
	respJSON := " {\"Name\":\"RookieOHY2\",\"Age\":25}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/*
③反射获取结构体的struct tag 信息,让对应的字段可以实现json字符串和结构体对象之间的互相转换
*/
func main() {
	p := person{
		Name: "RookieOHY",
		Age:  25,
	}
	pt := reflect.TypeOf(p)
	//遍历pt的tag
	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n", sf.Name, sf.Tag.Get("json"))
	}
}
