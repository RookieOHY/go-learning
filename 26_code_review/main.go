package main

import "fmt"

/*
Go 代码审查方法：
	扫描的内容：
		静态扫描检查（业务无关；只关注代码）
			--> 如：未使用的常量、函数有返回值但是没有被使用、拼写问题、死代码、代码简化检测、命名中带下划线、冗余代码等。
		性能优化
			--> 前提是保证是正确的，做的性能优化才有意义。
	扫描工具：
		golangci-lint、golint、gofmt、misspell
	golangci-lint一些配置项：
		根据自己的需要，参考官方的配置选择适合自己配置内容。
		配置案例.golangci.yml
		推荐将golangci集成具体的ci流程中：gitlabci；github action；jenkins等
	堆内存和栈内存：
		栈内存由编译器自动分配和释放。
			一般存储：局部变量、参数。
			函数创建时，内存就会自动申请，函数返回时，内存释放。
		堆内存生命周期大于栈内存。
			一般是函数的返回值被别的地方使用时，编译器会被值转移到堆上。
			堆内存不会被编译器释放，需要的是GC来释放。
		对比：
			栈内存效率大于堆内存
	内存逃逸：
		原本应该被存储在栈上的变量，因为一些原因转移到了堆上
	优化技巧：
		项目的代码要尽可能避免内存逃逸现象，才不会被GC拖累程序。

*/
/*
①内存逃逸分析demo
	使用
		go build -gcflags="-m -l" .\26_code_review\httpMethod.go 编译测试是否存在内存逃逸
	结果
		26_code_review\httpMethod.go:39:8: new(string) escapes to heap
	结论
		函数返回指针类型的变量，一定会发生内存逃逸。
		被已经逃逸的指针引用的变量也会发生逃逸.
		slice、map 和 chan，被这三种类型所引用的指针也会发生逃逸
*/
func main() {
	//newString02()
	//fmt.Println("RookieOHY")
	m := map[int]*string{}
	s := "RookieOHY"
	m[0] = &s

}
func newString() *string {
	//申请一块内存
	s := new(string)
	*s = "RookieOHY"
	fmt.Printf("%T", s)  //变量s的类型
	fmt.Printf("%T", *s) //变量s的类型

	return s
}

/*
②逃逸到堆内存的变量不能马上被回收，只能通过垃圾回收标记清除，增加了垃圾回收的压力，所以要尽可能地避免逃逸。
让变量分配在栈内存上，这样函数返回时就可以回收资源，提升效率。
一下demo为优化版 返回实际的变量 不会发生内存逃逸
*/

func newString02() string {
	//申请一块内存
	s := new(string)
	*s = "RookieOHY"
	return *s
}
