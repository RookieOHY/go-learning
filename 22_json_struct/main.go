package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
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
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}

/*
③反射获取结构体的struct tag 信息,让对应的字段可以实现json字符串和结构体对象之间的互相转换
*/
func main03() {
	p := person{
		Name: "RookieOHY",
		Age:  25,
	}
	pt := reflect.TypeOf(p)
	//遍历pt的tag
	for i := 0; i < pt.NumField(); i++ {
		//返回的对象是StructField（属性对象）：包含了属性的所有信息
		sf := pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n", sf.Name, sf.Tag.Get("json"))
		fmt.Printf("字段%s上,bson tag为%s\n", sf.Name, sf.Tag.Get("bson"))
	}
}

/*
④struct 转json
*/
func main04() {
	p := person{
		Name: "RookieOHY",
		Age:  25,
	}
	//分别获取属性和值的对象
	vp := reflect.ValueOf(p)
	tp := reflect.TypeOf(p)
	//拼接字符串的方式来构建json字符串
	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")
	for i := 0; i < tp.NumField(); i++ {
		//获取属性的tag
		fieldName := tp.Field(i).Tag.Get("json")
		jsonBuilder.WriteString("\"" + fieldName + "\"")
		jsonBuilder.WriteString(":")
		//获取值
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", vp.Field(i)))
		//最后一个不需要加上逗号
		if i < tp.NumField()-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String())
}
