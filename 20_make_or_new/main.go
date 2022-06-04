package main

import "fmt"

/*
	make、new的使用场景。
	默认的，go语言自动帮我们管理栈内存。而堆内存需要自己管理。
	程序的数据、程序本身大部分内存在堆上
*/
/*①值类型和指针类型的初始化问题：
默认的对于值类型的变量，不手动初始化，go默认帮我们初始化、自动分配内存。之后可以随意的进行变量的赋值和计算等操作
而对于指针类型，不手动初始化，go不会帮助其初始化，不分配对应的内存。因此无法对变量执行赋值和计算等操作
*/
func main01() {
	var s string
	s = "RookieOHY"
	fmt.Println(s)
	//报错：panic: runtime error: invalid memory address or nil pointer dereference
	//解析：s1是指针类型的变量，没有初始化默认为nil,没有指向的内存。无法对一个没有指向确定内存地址的变量作赋值等操作。
	var s1 *string
	*s1 = "RookieOHY"
	fmt.Println(s1)
}

/*---------*/
/*
②new函数改造①中的错误
	new函数的作用：为一个指针类型的变量（因为还没有被分配内存）开辟一块内存并且初始化零值。
*/
func main02() {
	var s *string
	//new函数为指针类型的变量开辟了一块内存空间,默认初始化了，返回指向该内存的指针
	s = new(string)
	//空字符串
	fmt.Println(*s)
	//修改s所指向的内存
	*s = "RookieOHY uses new-function to solve it!"
	fmt.Println(*s)
	fmt.Println(s)
}

/*----*/
/*
③使用工厂+new初始化一个非零值的xxx类型的指针
*/
func main03() {
	p := NewPerson("RookieOHY", 25)
	fmt.Println("p的name", p.name, "age", p.age)
}
func NewPerson(name string, age int) *person {
	p := new(person)
	p.name = name
	p.age = age
	return p
}

type person struct {
	name string
	age  int
}

/*
④make函数：
	make 函数只用于 slice、chan 和 map 这三种内置类型的创建和初始化，因为这三种类型的结构比较复杂.
	也是这三种类型的工厂函数。
*/
