package main

import "fmt"

/*
定义一个方法，有值类型接收者和指针类型接收者两种。二者都可以调用方法，因为 Go 语言编译器自动做了转换，所以值类型接收者和指针类型接收者是等价的。
但是在接口的实现中，值类型接收者和指针类型接收者不一样
*/
type person struct {
	age     int
	name    string
	address address
}
type address struct {
	province string
	city     string
}

/*接口*/
////Stringer是Go的一个接口，属于fmt包。
//type Stringer interface {
//	String() string
//}

//接口实现1
func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

//接口实现2
func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

//s 表示接口的实现类型 那么 Person就应该
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	printString(person{age: 15, name: "ou"})
	printString(address{city: "北京", province: "北京"})
}
