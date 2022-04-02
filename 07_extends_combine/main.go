package main

import "fmt"

/*
go语言的继承和组合: go中无父子关系 无继承 推荐和提倡使用组合（对于结构体和接口而言）
*/
/*1.接口之间的组合*/
type Name01 interface {
	ReadName01() string
}
type Name02 interface {
	ReadName02() string
}
type Name03 interface {
	Name01
	Name02
}

type address struct {
	city     string
	province string
}

/*2.结构体之间的组合：
外部类型不仅可以使用内部类型的字段，也可以使用内部类型的方法，就像使用自己的方法一样。
如果外部类型定义了和内部类型同样的方法，那么外部类型的会覆盖内部类型，这就是方法的覆写。

*/
//person是外部类型
type person struct {
	age  int
	name string
	//address是内部类型
	address
}

func main() {
	p := person{
		age:     0,
		name:    "ou",
		address: address{"beijing", "beijing"},
	}
	fmt.Println(p.age)
	fmt.Println(p.address.city)
	fmt.Println(p.city)

}
