package main

import (
	"fmt"
	"sync"
	"time"
)

/*17.go的一些并发场景demo*/
/*①select timeout:适用于网络请求的超时*/
func main01() {
	resp := make(chan string)
	go func() {
		//假设响应8s后到底
		time.Sleep(8 * time.Second)
		resp <- "请求成功"
	}()
	//若5s后请求没响应，默认视作超时
	select {
	case data := <-resp:
		fmt.Println("相应结果", data)
	case <-time.After(8 * time.Second):
		fmt.Println("请求超时了")
	}
}

/*----------------------------------------------*/
/*②Pipeline模式：多个channel之间互相协作完成一系列操作*/
//购买
func buy(num int) <-chan string {
	out := make(chan string)
	go func() {
		//购买完毕 关闭
		defer close(out)
		for i := 1; i < num; i++ {
			out <- fmt.Sprint("零件-", i)
		}
	}()
	return out
}

//组装
func build(in <-chan string) <-chan string {
	out2 := make(chan string)
	go func() {
		//组装完毕 关闭
		defer close(out2)
		for c := range in {
			out2 <- "组装" + c
		}
	}()
	return out2
}

//打包
func pack(in <-chan string) <-chan string {
	out3 := make(chan string)
	go func() {
		//打包 关闭
		defer close(out3)
		for i := range in {
			out3 <- "打包" + i
		}
	}()
	return out3
}
func main02() {
	//购买10个零件
	buyCh := buy(10)
	//安装10个零件
	buildCh := build(buyCh)
	//打包10个零件
	packCh := pack(buildCh)
	for ch := range packCh {
		fmt.Println(ch)
	}
}

/*③扇入和扇出模式：适用于流水模式某一个环节产能无法跟上的场景：如buy环节购买的零件过多，build环节产能不够，导致pack环节产能过剩*/
//扇入：汇聚多个channel的数据
func mergeChannel(ins ...<-chan string) <-chan string {
	var waitGroup sync.WaitGroup
	out := make(chan string)
	//启动多个协程并发获取ins的数据
	//遍历一个ch中的值
	p := func(ch <-chan string) {
		defer waitGroup.Done()
		for v := range ch {
			out <- v
		}
	}
	//监听n个协程
	waitGroup.Add(len(ins))
	//启动n个协程
	for _, cs := range ins {
		//读取
		go p(cs)
	}
	//等待n个协程执行完毕
	go func() {
		waitGroup.Wait()
		//关闭ch
		close(out)
	}()
	//返回聚合的结果
	return out
}
func main03() {
	buyCh := buy(100)
	buildCh01 := build(buyCh)
	buildCh02 := build(buyCh)
	buildCh03 := build(buyCh)
	//扇入
	buildCh := mergeChannel(buildCh01, buildCh02, buildCh03)
	result := pack(buildCh)
	for v := range result {
		fmt.Println(v)

	}
}

/*-----------------------------------------------*/
/*④Futures模式：协程没有互相依赖和协程直接互相依赖的场景*/
/*如：做火锅时一个很大的任务：需要洗菜，烧水。二者没有依赖，但是刷火锅时就需要二者都先准备好*/
func vegetable() <-chan string {
	vegesCh := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegesCh <- "菜洗干净了"
	}()
	return vegesCh
}
func water() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "水滚了"
	}()
	return water
}
func main() {
	ch1 := vegetable()
	ch2 := water()
	//如果没有读取到值 一直阻塞
	vegesResult := <-ch1
	waterResult := <-ch2
	fmt.Println("结果：", vegesResult, waterResult)
}
