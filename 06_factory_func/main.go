package main

import "fmt"

/*工厂函数：用于创建自定义结构体*/
type person struct {
	age  int
	name string
}

func NewPerson(name string) *person {
	return &person{name: name}
}

func main01() {
	//使用工厂函数实现新建一个person
	newPerson := NewPerson("ou")
	fmt.Println(newPerson.name)
}

/*工厂函数：新建一个接口error*/
//工厂函数，返回一个error接口，其实具体实现是*errorString
func New(text string) error {
	return &errorString{text}
}

//结构体，内部一个字段s，存储错误信息
type errorString struct {
	s string
}

//用于实现error接口
func (e *errorString) Error() string {
	return e.s
}

func main() {
	err := New("错误信息")
	if err == nil {
		return
	}
	fmt.Println(err)
}
