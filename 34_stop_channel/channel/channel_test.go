package channel

import (
	"fmt"
	"testing"
	"time"
)

/*
	stop函数的测试
*/
func TestStopUtil(t *testing.T) {
	//构建管道
	//chan struct{}:不会存储数据，扮演信号量的角色，控制程序的启停
	stopCh := make(chan struct{})
	//程序代码
	fn := func() {
		fmt.Println("xxx program run ~")
	}
	//设定定时逻辑：3s之后程序代码停止的动作
	//开辟一个协程：3s之后关闭管道
	go func(chan struct{}) {
		select {
		case <-time.After(3 * time.Second):
			close(stopCh)
			fmt.Println("stopCh closed: program will terminate")
		}
	}(stopCh)
	//调用停止函数
	StopUtil(stopCh, fn)
}
