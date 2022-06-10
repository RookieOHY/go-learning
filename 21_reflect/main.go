package main

import (
	"fmt"
	"io"
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

func (p person) String() string {
	return fmt.Sprintf("Name is %s,Age is %d", p.Name, p.Age)
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
使用Kind()方法来获取：
	Kind方法返回Kind类型的常量
*/
func main04() {
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

/*⑤关于变量类型本身的反射，推荐使用reflect.Type()
reflect.Type和reflect.Value有所不同：前者是接口，后者是结构体。二者拥有一些相同的方法。
*/
func main() {
	p := person{
		Name: "RookieOHY",
		Age:  25,
	}
	//反射获取任意变量的类型
	var a = 29
	str := "我是字符串"
	pt := reflect.TypeOf(p)
	at := reflect.TypeOf(a)
	strt := reflect.TypeOf(str)
	fmt.Println("类型p：", pt)
	fmt.Println("类型a：", at)
	fmt.Println("类型str：", strt)
	//获取person的属性以及方法
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("属性名字：", pt.Field(i).Name)
		fmt.Println("属性类型：", pt.Field(i).Type)
	}
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法名字：", pt.Method(i).Name)
	}
	//获取person中指定的方法
	mn, flag := pt.MethodByName("String")
	if flag == true {
		fmt.Println("获取指定的方法名字成功! ", mn.Name)
	}
	//判断是否实现了某一个接口（如判断person是否实现了fmt.Stringer 和 io.Writer）
	//获取接口类型
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	iot := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("person是否实现了Stringer接口：", pt.Implements(st))
	fmt.Println("person是否实现了Writer接口：", pt.Implements(iot))

}
