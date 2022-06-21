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

/*
	基准测试：
		必须是BenchMark+函数名
		一定导出的
		被测试的代码逻辑要放在循环
		无返回值
	基准测试一些命令：
		go test -bench=. 相对路径

	结果分析：
		-16 --> GOMAXPROCS 逻辑cpu的数量
		1576354 --> for循环的次数
		653.1 ns/op --> 每次循环花费的纳秒数
		0 B/op --> 每次循环分配的内存(单位字节)
		0 allocs/op --> 每次循环分配内存的次数


*/
func BenchmarkFibonacci(b *testing.B) {
	//构建测试用例也是需要时间的，但是按理不应该算上
	n := 10
	//启用内存统计 计算每一次操作分配内存的次数和大小
	b.ReportAllocs()
	//使用重置定时器重置时间(还有StartTimer 和 StopTimer)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibonacci(n)
	}
}

/*
并发基准测试
	把n分配给多个goroutine并发执行
*/

func BenchmarkFibonacciRunConcurrent(b *testing.B) {
	n := 10
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fibonacci(n)
		}
	})

}

/*
注意点：
	编码-->单元测试-->覆盖率分析-->普通基准测试-->并发基准测试-->修改或者调整你的代码
	单元测试是保证代码质量的好方法，但单元测试也不是万能的.使用它可以降低 Bug 率，但也不要完全依赖。
	除了单元测试外，还可以辅以 Code Review、人工测试等手段更好地保证代码质量。
*/
