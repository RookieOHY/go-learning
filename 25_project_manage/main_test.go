package main

import "testing"

/*
	单元测试的一些命令：
		① go test -v 测试代码的相对路径 例如： go test -v .\25_project_manage\
			测试结果和期望值比较，这就是单元测试想要达到的目的。
		②  go test -v --coverprofile=fb.cover .\25_project_manage\
			返回对应的覆盖率以及生成对应的覆盖率文件
*/
//单元测试的函数名为Test+被测试的函数名；入参一定是 *testing.T 指针类型
//单元测试的函数不会有任何的返回值
func TestFibonacci(t *testing.T) {
	//下面是一些测试用例
	fsMap := map[int]int{}
	fsMap[0] = 0
	fsMap[1] = 1
	fsMap[2] = 1
	fsMap[3] = 2
	fsMap[4] = 3
	fsMap[5] = 5
	fsMap[6] = 8
	fsMap[7] = 13
	fsMap[8] = 21
	fsMap[9] = 34
	for k, v := range fsMap {
		fib := Fibonacci(k)
		if v == fib {
			t.Logf("结果正确:n为%d,值为%d", k, fib)
		} else {
			t.Errorf("结果错误：期望%d,但是计算的值是%d", v, fib)
		}

	}

}
