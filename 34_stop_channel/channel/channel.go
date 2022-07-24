package channel

import (
	"fmt"
	"time"
)

/*
	定义函数：
		监听管道channel的值，值变化，停止程序运行
*/
// StopUtil 参数：管道；函数
func StopUtil(stopCh chan struct{}, fn func()) {
	//监听
	for {
		select {
		//结束运行程序（空chan不用于存储值，何时可以读到值，该chan被close时触发）
		case <-stopCh:
			return
		//1s之后执行的逻辑
		case <-time.After(1 * time.Second):
			fmt.Println("xxx program begin running ")
			//执行函数（具体的程序代码）
			fn()
		}
	}
}
