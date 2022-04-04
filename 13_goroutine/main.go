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

//有缓冲channel和无缓冲的channel:
//无缓冲：使用make创建的channel默认是无缓冲的，容量为0,不可以存储数据。channel接收数据和发送数据是同时进行的，可以称之为同步channel。
//有缓冲：make创建channel的时候同时指定容量.特点是先进先出，类似于队列。大小表示，最多可以存储5个int类型的对象
func main03() {
	ch := make(chan int, 5)
	//打印有缓冲channel的长度和容量
	ch <- 5
	ch <- 4
	fmt.Println("当前有缓冲channel的大小为:", len(ch), ",容量为：", cap(ch))
}

//关闭channel:如果channel已经关闭了，那么不允许向channel发送数据，否则会报异常。
func main() {
	ch := make(chan int, 5)
	go func() {
		ch <- 5
		ch <- 4
		time.Sleep(10 * time.Second)
		close(ch)
		ch <- 3 //出现异常
	}()
	for i := 0; i < 2; i++ {
		v := <-ch //每次只能读取一个
		fmt.Println(v)
	}

}

//单向
