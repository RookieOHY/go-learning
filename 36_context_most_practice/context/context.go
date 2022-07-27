package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ContextWithCancel withCancel最佳实践：监听和取消协程
func ContextWithCancel() {
	//获取一个bg 上下文对象
	bgctx := context.Background()
	//返回参数：前者是一个上下文对象 后者是一个函数
	ctx, cancel := context.WithCancel(bgctx)
	//看下这个ctx到底是什么
	fmt.Println(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	//4s后主动关闭（公用同一个上下文）
	go func(context.Context, *sync.WaitGroup) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				fmt.Println("monitor1 listen ctx done")
			case <-time.After(4 * time.Second):
				fmt.Println("monitor being do cancel")
				cancel()
			}
		}
	}(ctx, wg)
	//被动关闭
	go func(context.Context, *sync.WaitGroup) {
		index := 0
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				fmt.Println("monitor2 listen ctx done")
				return
			case <-time.After(2 * time.Second):
				index += 1
				fmt.Println("index add 1")
			}
		}
	}(ctx, wg)
	//阻塞main
	wg.Wait()
}
