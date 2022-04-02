package main

import "fmt"

//go的异常(对于可以预知的异常，可以捕获或者抛出)
//01.抛出异常会导致程序停止运行

func test(a int) {
	if a == 0 {
		//func panic(v interface{}) 这是一个空接口，支持任意的类型
		panic("我是异常 我是测试")
	}
}

func main01() {
	test(0)
}

//02.即使出现panic出现异常，也需要释放资源。
//02.1 使用内置的recover函数捕获panic异常
func main02() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	test(0)
}
func multipartDefer() {
	defer fmt.Println("stack 1")
	defer fmt.Println("stack 2")
	defer fmt.Println("stack 3")
	fmt.Println("stack 4")
}

//03.是否可以定义多个defer,如果可以，执行顺序是什么样的？
func main() {
	multipartDefer() //查看执行结果：可以定义多个defer捕获，捕获的顺序为先进后出
}
