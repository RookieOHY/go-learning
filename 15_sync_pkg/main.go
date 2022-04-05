package main

import (
	"fmt"
	sync "sync"
	"time"
)

/*
15.sync包的学习
*/
//sync.Once：执行一次
func doOnce() {
	var doNoce sync.Once
	svc := func() {
		fmt.Println("我是只执行一次的业务方法")
	}
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			doNoce.Do(svc)
			ch <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-ch
	}

}
func main01() {
	doOnce()
}

//sync.Cond:唤醒处于等待中协程，表示所有的协程都可以开始执行（如裁判和短跑运动员）;也可以阻塞协程
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var waitGroup sync.WaitGroup
	waitGroup.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer waitGroup.Done()
			fmt.Println(num, "号选手已就位！")
			cond.L.Lock()
			//阻塞所有的协程
			cond.Wait()
			fmt.Println(num, "跑了")
			cond.L.Unlock()
		}(i)
	}
	//设置延时，尽量保证10个协程都进入wait
	time.Sleep(2 * time.Second)
	go func() {
		defer waitGroup.Done()
		fmt.Println("裁判就位了")
		fmt.Println("比赛开始跑")
		cond.Broadcast()
	}()
	//阻塞等待
	waitGroup.Wait()
}
