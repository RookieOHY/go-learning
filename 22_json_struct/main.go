package main

import (
	"encoding/json"
	"fmt"
)

/*json 和 struct之间的互相转换*/
/*
①利用json包Marshal函数的，将结构体对象转换为json
*/
func main() {
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
	Name string
	Age  int
}
