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
func main03() {
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

/*④函数中值类型无法修改结构体对象的属性的原因：
本质是结构体对象和函数形参的内存地址不同，对于结构体的形参，本质上是原来的变量的拷贝，因此他们的内存地址是不一样
虽然他们的属性和值一样。
*/
/*------------*/

/*
⑤使用map来表示一个对象。针对make或者是字面量来创建的map,本质上都会调用 src/runtime/map.go的makemap，
该函数返回的是一个指针类型的变量*hmap。因此本质上make map返回的是*hmap。
因此，无论是main函数还是modify函数，本质上使用的都是指针作为参数，导致修改行为是执行了的。
*/
func main() {
	personMap := make(map[string]int)
	personMap["名字"] = 1
	personMap["年龄"] = 2
	fmt.Printf("personMap地址为：%p\n", personMap)
	modifyMap(personMap)
	fmt.Println(personMap)

}
func modifyMap(personMap map[string]int) {
	personMap["名字"] = 3
	fmt.Printf("personMap地址为：%p\n", personMap)
}
