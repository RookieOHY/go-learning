package main

import (
	"fmt"
	"reflect"
)

/*
17.go中的指针理解
	含义：指针是数据类型，用来存储一个内存地址，该地址指向存储内存中的对象（字符串、整数、函数、结构体），帮助快速找到内存中的数据
	指针=>内存地址=>对象
	书的页码=>对应页码的内容
*/
func main01() {
	var name = "RookieOHY"
	fmt.Println("name的内存地址->", &name)
	//打印变量的类型
	fmt.Printf("%T", &name)
	fmt.Println(reflect.TypeOf(name))
}

//指针变量的声明
func main02() {
	var a *int
	fmt.Printf("%T", a)
	//使用new
	b := new(int)
	fmt.Printf("%T", b)
}

//指针的常用操作：
func main03() {
	var name = "RookieOHY"
	nameP := &name
	nameV := *nameP //获取指针指向的值: * 指针变量
	fmt.Println("nameP指针指向的值为:", nameV)
	*nameP = "RookieOHY2" //*nameP重新赋值等于修改了指针nameP指向的值
	nameV2 := *nameP
	fmt.Println("nameP指针指向的值修改为:", nameV2)
	fmt.Println(&nameV)
	fmt.Println(&nameV2)
	fmt.Println(name)
}

/*---------------------------------*/
var age int = 18

func modifyAge(age int) {
	age++
}

func modifyAge2(age *int) {
	*age = 19 //修改指针指向的值
}

//指针与函数参数:使用形参来改变外部的实参的值
func main() {
	modifyAge(1)
	fmt.Println(age)
	modifyAge2(&age) //指针类型
	fmt.Println(age)
}
