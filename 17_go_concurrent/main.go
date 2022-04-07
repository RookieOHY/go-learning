package main

import (
	"fmt"
	"time"
)

/*17.go的一些并发场景demo*/
/*①网络请求的超时*/
func main() {
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
