package main

import (
	"fmt"
	"sync"
	"time"
)

/**
14.go的sync包学习
*/
//使用互斥锁保证多个协程操作同一份内存中的数据时，是资源竞争是安全的，符合预想中的结果的
//sync.Mutex
var (
	sum   int
	mutex sync.RWMutex
	//mutex sync.Mutex
)

func add(i int) {
	mutex.Lock() //加锁
	sum += i
	defer mutex.Unlock()
}
func main01() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	//必要：加上延时，让main协程可以读到结果
	time.Sleep(5 * time.Second)
	fmt.Println(sum)
}

/*----------------------------------------------*/
//sync.RWMutex:读写锁，应用于并发读写场景
func readSum() int {
	mutex.RLock() //避免了多个读协程阻塞（互斥锁），读写锁（一起读，不要互相阻塞了）
	defer mutex.RUnlock()
	b := sum
	return b
}
func main02() {
	for i := 0; i < 100; i++ {
		fmt.Println("调用add()")
		go add(10)
	}
	//time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		go fmt.Println("和为:", readSum())
	}
	time.Sleep(2 * time.Second)
}

/*-----------------------------------------*/
//分析time.sleep作用：保证子协程先于main协程执行完毕。除非你可以人为的估测n个协程的执行总花费的时间，保证结果的main协程退出之前执行完毕。
//否则不推荐人为的设置等待延时。
func main() {
	run()
}
func run() {
	var watch sync.WaitGroup
	watch.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			defer watch.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer watch.Done()
			fmt.Println(readSum())
		}()
	}
	//只有110个协程全部done(),代表了110个协程都存在返回结果，那么才可以让run退出
	watch.Wait()
}
