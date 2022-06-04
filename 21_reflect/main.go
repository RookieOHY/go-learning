package main

import (
	"fmt"
	"reflect"
)

/*反射基础知识点*/
/*
①利用反射来获取变量的类型
*/
func main01() {
	i := 3
	//ValueOf将任意类型的对象转换为reflect.Value对象(该函数返回的值i的值的拷贝)
	vi := reflect.ValueOf(i)
	fmt.Printf("vi转换后类型为%T", vi)
	//reflect.Value的Interface方法可以将reflect.Value对象转换为原来的类型
	vi02 := vi.Interface().(int)
	fmt.Printf("vi原来的类型为%T", vi02)
}

/*②利用反射修改原来的值*/
func main02() {
	i := 4
	ip := reflect.ValueOf(&i)
	//使用Elem获取指针变量指向的值
	v := ip.Elem()
	fmt.Println(v)
	flag := v.CanSet()
	//判断是否执行修改
	fmt.Println(flag)
	//修改指针变量指向的值
	ip.Elem().SetInt(25)
	fmt.Println(v)
}

/*③修改结构体中属性的值
什么样的结构体的属性可以被修改呢？
	属性是非私有的结构体才可以被修改。

*/
type person struct {
	Name string
	Age  int
}

func main03() {
	//新建一个变量
	p := person{
		Name: "ohy",
		Age:  24,
	}
	//获取对应reflect.Value类型的指针变量
	pp := reflect.ValueOf(&p)
	//获取指针变量pp指向的值
	ppv := pp.Elem()
	fmt.Println("ppv是指针变量pp的值，他的值为", ppv)
	//获取属性和修改属性的值
	ppv.Field(0).SetString("RookieOHY")
	ppv.Field(1).SetInt(25)
	fmt.Println("修改后的值为", ppv)
}

/*④获取底层结构类型
哪些属于底层结构类型：
	接口、结构体、指针
	如p的底层类型是struct、&p的底层类型是指针
*/
func main() {
	p := person{
		Name: "ohy",
		Age:  24,
	}
	pp := reflect.ValueOf(&p)
	pv := reflect.ValueOf(p)
	//指针变量的底层类型
	fmt.Println(pp.Kind())
	//变量的底层类型
	fmt.Println(pv.Kind())
}
