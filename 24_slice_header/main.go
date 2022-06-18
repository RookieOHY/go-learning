package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*slice存在意义和slice的原理
数组之外为啥还需要切片的存在呢？
	切片比较高效
		1.在赋值、函数传参的时候，并不会把所有的元素都复制一遍，而只是复制 SliceHeader 的三个字段就可以了，共用的还是同一个底层数组。

数组的局限性？
	1.强类型语言的数组有两个属性：一个是长度，一个类型。一旦二者确定了，都不可以被改变了！
		无法添加比容量大的元素；对于大数据量的数组，应该确定申明的大小
	2.容量比较大的数组在作为参数传递时，由于值传递，造成大量的内存被开辟，容易造成浪费。

引入切片~
	1.切片的底层也是一个数组，本质是对数组的一个抽象封装。
	2.切片可以动态扩容。
		当原本的底层数组容量不够时，append操作会新建一个底层数组，之后将旧底层数组的元素拷贝至新的数组，最后返回一个指向新数组的切片。
*/
//①切片的扩容
func main01() {
	ss := []string{"RookieOHY", "RookieOHY2"}
	fmt.Println("切片的初始长度为", len(ss), "容量为", cap(ss))
	//append
	ss = append(ss, "RookieOHY3", "RookieOHY4", "RookieOHY5", "RookieOHY6")
	fmt.Println("切片扩容后长度为", len(ss), "容量为", cap(ss))
	//print
	fmt.Println(ss)
}

/*②切片的底层结构
type SliceHeader struct {
	Data uintptr //底层的数组
	Len  int
	Cap  int
}
*/
func main02() {
	//不同切片可能指向同一个底层数组
	ss := [2]string{"A", "B"}
	a := ss[0:1]
	b := ss[:]
	//打印切片a b 的底层数组data
	//如何打印
	//(因此，当切片作为函数传递时，尽量不要修改原数组的元素)
	//(切片作为参数传递时，值传递，存在一个副本，最多占用24字节的内存大小。)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&ss)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&a)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
}

/*
③探究切片作为函数参数时 和 外部的切片 是否用一个底层数组？
*/
func main03() {
	cd := [2]string{"c", "d"}
	//打印数组的地址
	fmt.Printf("外部数组的地址为%p\n", &cd)
	//打印数组方法
	printArrayAddress(cd)
	cds := cd[0:1]
	//打印外部切片的地址；切片数组的底层地址
	fmt.Printf("外部切片的地址为%p\n", &cds)
	fmt.Printf("外部切片的底层数组地址%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&cds)).Data)
	//打印内部切片的底层数组地址
	printSliceAddress(cds)
}
func printArrayAddress(cd [2]string) {
	fmt.Printf("内部数组的地址为%p\n", &cd)
}
func printSliceAddress(cds []string) {
	fmt.Printf("内部切片的底层数组地址为%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&cds)).Data)
}

/*
④string 和 []byte的互转（存在拷贝）体现slice的高效.
*/
func main04() {
	//两次的值传递 都存在拷贝操作 因此打印出来的内容虽然都是一样的 但是地址不一样
	//如果字符串非常大 ，复制时耗费的内存就很大
	s := "RookieOHY"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)

}

/*
⑤如何让④转换时，不进行值拷贝(使用unsafe.Pointer实现)
*/
func main() {
	s := "RookieOHY"
	b := []byte(s)
	s4 := *(*string)(unsafe.Pointer(&b))
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&b)).Data)
	fmt.Printf("s4的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s4)).Data)

}
