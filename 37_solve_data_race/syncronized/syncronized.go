package syncronized

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*

	数据竞争解决方案：
		unsafe 案例
		解决的案例：
			sync/atomic
			sync/Mutex 类比java的lock
			sync/RMutex 类比ReentrantReadWriteLock

*/
//非原子的加法操作
func add(w *sync.WaitGroup, num *int) {
	defer w.Done()
	*num = *num + 1
}

//unsafe 案例
func UnsafeAdd() {
	var n int = 0
	wg := new(sync.WaitGroup)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add(wg, &n)
	}
	//阻塞
	wg.Wait()
	fmt.Println(n)
}

//使用atomic
func SafeAdd() {
	var n int32 = 0
	wg := new(sync.WaitGroup)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&n, 1)
		}()
	}
	//阻塞
	wg.Wait()
	fmt.Println(n)
}

//使用互斥锁
