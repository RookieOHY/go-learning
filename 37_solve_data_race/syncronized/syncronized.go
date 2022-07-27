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
func Mutex() {
	var number int = 0
	//新建锁
	mutex := new(sync.Mutex)
	w := new(sync.WaitGroup)
	w.Add(2)
	//只有两个协程
	go func() {
		defer w.Done()
		for i := 0; i < 1000; i++ {
			fmt.Println("t1 run add")
			mutex.Lock()
			number = number + 1
			mutex.Unlock()
		}
	}()

	go func() {
		defer w.Done()
		for i := 0; i < 1000; i++ {
			fmt.Println("t2 run minus")
			mutex.Lock()
			number = number - 1
			mutex.Unlock()
		}
	}()

	w.Wait()
	fmt.Println(number)
}

//只有一个协程可以执行成功 多用初始化操作（如初始化数据库连接）
func Once() {
	once := new(sync.Once)
	wg := new(sync.WaitGroup)

	wg.Add(5)

	for i := 0; i < 5; i++ {
		tmp := i
		go func() {
			defer wg.Done()
			//fmt.Println(tmp)
			once.Do(func() {
				fmt.Println(tmp)
			})
		}()
	}

	wg.Wait()
}
