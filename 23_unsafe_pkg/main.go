package main

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
② unsafe.Pointer实现互转
*/
func main() {

}
