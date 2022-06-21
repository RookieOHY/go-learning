package main

/*
	Go做好工程的管理有助于保证代码的质量，提升代码的性能。
*/
/*
	①单元测试：测试的单元是一个完整的最小单元。
		比如 Go 语言的函数就是一个最小单元。
		当每个最小单元都被验证通过，那么整个模块、甚至整个程序就都可以被验证通过。
		改动者是单元测试代码的编写者。
	②Go 函数main.go对应的单元测试一定为main_test.go
	③基础测试：测试代码的性能

*/
func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
