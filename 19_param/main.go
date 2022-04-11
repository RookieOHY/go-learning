package main

import "fmt"

/*19.参数传递*/
/*①值做为形参，实现了某一个接口，那么值对应的指针也实现了该接口*/
type address struct {
	province string
	city     string
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}
func printAddress(s fmt.Stringer) {
	fmt.Println(s.String())
}
func main01() {
	address := address{
		province: "北京",
		city:     "北京",
	}
	printAddress(address)  //对象做形参
	printAddress(&address) //对象的指针形参
}

/*②接口的指针类型作为参数*/
func main02() {
	var intfc fmt.Stringer = address{
		province: "北京",
		city:     "北京",
	}
	//接口对象
	printAddress(intfc)
	//取接口对象的指针
	//interfc := &intfc
	//printAddress(interfc) //指向具体类型的指针可以实现一个接口，而指向接口的指针永远不可能实现该接口
}

/*③修改结构体的属性值*/
/*----------------------------------*/
func main() {
	p := person{name: "张三", age: 18}
	modifyPerson(&p)
	fmt.Println("person name:", p.name, ",age:", p.age)

}
func modifyPerson(p *person) {
	p.name = "李四"
	p.age = 20
}

type person struct {
	name string
	age  int
}
