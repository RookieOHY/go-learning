package main

import "fmt"

/*类型的断言*/
type person struct {
	age  int
	name string
}
type address struct {
	city     string
	province string
}

//1.实现同一个接口
func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

//2.类型断言示例
func main() {
	var s fmt.Stringer
	p1 := person{
		age:  0,
		name: "ou",
	}
	s = p1
	p2, ok := s.(*person)
	if ok {
		fmt.Println(p2)
	} else {
		fmt.Println("接口的值s不是person类型")
	}

}
