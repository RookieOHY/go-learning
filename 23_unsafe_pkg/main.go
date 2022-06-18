package main

import (
	"fmt"
	"unsafe"
)

/*
unsafe包的一些用法(包含了一些直接操作go内存的操作，忽略且绕过了go的内存防护机制)
	默认下Go不允许2个指针类型互相转换的，这是处于安全的考虑。（如*int默认下无法转换为*float64）
*/
/*
① 指针类型无法强制转换的示例：程序在编译时便无法通过了！
*/
func main01() {
	//i := 10
	//ip := &i
	//Cannot convert an expression of the type '*int' to the type '*float64'
	//var fp *float64 = (*float64)(ip)
	//fmt.Println(fp)
}

/*
② unsafe.Pointer（一种可以表示任意类型的地址，在两个不同类型的强制转换中扮演中转的角色）实现互转。
	在Go中，属于ArbitraryType类型，表示任意类型，可以表示任何内存地址
	该类型的指针不可以做一些加法运算
*/
func main02() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)((unsafe.Pointer(ip)))
	*fp = *fp * 10
	fmt.Println(i)
}

/*
③	unsafe.uintptr
	通过它，可以对指针偏移进行计算，这样就可以访问特定的内存，达到对特定内存读写的目的
*/
type person struct {
	Name string
	Age  int
}

//demo里没有对p的某一个属性直接赋值，而是拿到对应的内存地址，操作内存进行赋值。
func main03() {
	p := new(person)
	//使用unsafe.Pointer获取Name（默认指向偏移量为0的内存地址，自然就定位到了Name属性）
	pName := (*string)(unsafe.Pointer(p))
	*pName = "RookieOHY"
	//先把p转换为unsafe.Pointer 再转换为uintptr 再利用进行偏移量计算 计算结果再转换为Pointer 最终获取到了Age属性
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 25
	//打印
	fmt.Println(*p)
}

/*
④unsafe.SizeOf 用于返回类型内存大小(和变量存储内存所占用的内存无关) 返回字节数
*/
func main() {
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(false))
	fmt.Println(unsafe.Sizeof(int(64)))
	fmt.Println(unsafe.Sizeof(int64(64)))
	fmt.Println(unsafe.Sizeof(int8(64)))
	strings := [1]string{"RookieOHY"}
	strings2 := [2]string{"RookieOHY"}
	//关于类型 数组长度不同 表示不同类型的字符串数组
	fmt.Printf("%T", strings)
	fmt.Printf("%T", strings2)

}
