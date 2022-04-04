package main

import (
	"fmt"
	"time"
)

/**
13.Go协程：Go中操作的是协程，协程比线程更加轻量。可以轻易的启动上万个协程
	语法：go 函数名
	协程如何通信：使用channel
*/
func main01() {
	go fmt.Println("我是协程1")
	//fmt.Println("我是协程2")
	panic("我出现异常了")
	//在main协程里面，新启动一个协程。若主main协程挂了，整一个程序将挂掉
	time.Sleep(time.Second) //让main协程等待一秒执行
}

//channel可以达到main01中的延迟效果，如果，ch中没有值，它会一直等待，直到有值
func main02() {
	ch := make(chan string)
	go func() {
		fmt.Println("我是协程1")
		time.Sleep(10 * time.Second)
		ch <- "我是写入ch的字符串"
	}()
	fmt.Println("我是main协程")
	str := <-ch
	fmt.Println("ch中接收到的值为：", str)
}

//有缓冲channel和无缓冲的channel:使用make创建的channel默认是无缓冲的，容量为0,不可以存储数据
func main() {

}
