package main

import "fmt"

/*
结构体和结构体之间互相引用（嵌套结构体）
*/
type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

func main() {
	p := person{
		age:  15,
		name: "ou",
		addr: address{"北京", "北京"},
	}
	fmt.Println(p)
}
