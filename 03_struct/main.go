package main

import "fmt"

/*
结构体
*/
type person struct {
	age  uint
	name string
}

func main() {
	var p person
	fmt.Println(p)          //很显然，打印的是结构体的初值
	p2 := person{14, "偶欢愉"} // 不支持重新被赋值
	//p := person{14, "偶欢愉"}
	fmt.Println(p2)
	//fmt.Println(p)
	fmt.Println(p2.name, p2.age) //使用.操作符调用属性
	p3 := person{name: "ou", age: 16}
	fmt.Println(p3) //初始化时同时显式指定赋值哪一个字段
	p4 := person{age: 0}
	fmt.Println(p4) //初始化时同时为某一个字段赋值即可

}
