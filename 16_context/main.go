package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
16.Context学习:单协程取消和多协程取消
*/
/*死循环监控*/
func watchDog(name string) {
	for {
		select {
		default:
			fmt.Println(name, "正在监听...")
		}
		time.Sleep(time.Second)
	}
}
func main01() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		watchDog("监控摄像头 ")
	}()

	wg.Wait()
}

/*---------------------------------------*/
/*如何让上面的死循环协程安全的退出*/
/*方案：使用select+channel*/
/*检测ch是否存在值，如果存在值，让监控协程退出*/
func watchDog02(watchChannel chan bool, name string) {
	for {
		select {
		case <-watchChannel:
			fmt.Println("监控退出啦！")
			return
		default:
			fmt.Println(name, "监控正在执行...")
		}
		//延时1s
		time.Sleep(time.Second)
	}
}
func main02() {
	var watchGroup sync.WaitGroup
	watchGroup.Add(1)
	watchChannel := make(chan bool)
	go func() {
		defer watchGroup.Done()
		watchDog02(watchChannel, "摄像头")
	}()
	time.Sleep(5 * time.Second)
	//5s后发送stop指令
	watchChannel <- true
	watchGroup.Wait()
}

/*-------------------------------*/
/*使用Context清理单、多个协程*/
/*根节点的Context为空Context,节点下存在其他的Context:
定时取消功能的Context
取消的Context
可传键值对的Context
*/
func watchDog03(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): //Done会返回一个只读的channel,如果channel中存在值，也就是Context已经发起了取消信息，代表协程取消了，改释放内存资源了。
			fmt.Println("协程取消了")
			return
		default:
			fmt.Println(name, " 监控正在执行...")
		}
		time.Sleep(time.Second)
	}

}
func main03() {
	var watchGroup sync.WaitGroup
	watchGroup.Add(3)
	//使用go内置的函数来获取Context
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer watchGroup.Done()
		watchDog03(ctx, "摄像头")
	}()
	go func() {
		defer watchGroup.Done()
		watchDog03(ctx, "摄像头02")
	}()
	go func() {
		defer watchGroup.Done()
		watchDog03(ctx, "摄像头03")
	}()
	time.Sleep(5 * time.Second)
	stop()
	//等待
	watchGroup.Wait()
}

/*-------------------------------------*/
/*Context的传值：Context的值可以供给其他的多个协程使用*/
func main() {
	var watchGroup sync.WaitGroup
	watchGroup.Add(1)
	//使用go内置的函数来获取Context：Background()生成根节点Context
	ctx, stop := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, "userId", 2)
	go func() {
		defer watchGroup.Done()
		getUser(valueCtx)
	}()
	time.Sleep(5 * time.Second)
	stop()
	//等待
	watchGroup.Wait()
}

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("我退出了")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("用户的id：", userId)
			time.Sleep(time.Second)
		}
	}
}
