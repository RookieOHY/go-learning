package _9_panic_handling

import "fmt"

/*
	Go程序异常处理最佳实践
		当前协程 和 recover() 是一对一的。
		不可以擦除其他协程的异常。
*/

// PanicSimulate 模拟异常抛出 和 协程中异常堆栈擦除
func PanicSimulate(key int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	if key == 2 {
		panic("key==2,异常了")
	}
	fmt.Println("is ok,key ==", key)
}

//启用3个协程
func runRoutines() {
	for i := 0; i < 3; i++ {
		go PanicSimulate(i)
	}
	for {
		select {
		default:
			//fmt.Println("run..........")
		}
	}
}
