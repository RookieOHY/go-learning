package syncronized

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*

	数据竞争解决方案：
		unsafe 案例
			不安全例子：累加操作
			解决方式：
				sync/atomic
				sync/Mutex 类比java的lock
				sync/RMutex 类比ReentrantReadWriteLock
		生产者、消费者模型
			例子：容量满了，生产者停止生产；容量空了，消费者停止消费
			本质：协程之间的通信。是同步问题
			解决方式：
				sync/Cond

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

//使用atomic保证累加操作安全
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

//使用互斥锁实现加减互斥
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

//只有一个协程可以执行成功 其余的都会失败 多用初始化操作（如初始化数据库连接）==> Do方法只会被执行一次
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

//阻塞式的生产者消费者模型案例
const min = 0
const max = 3

type iceCube int
type cup struct {
	iceCubes []iceCube
}

//开启两个协程 分别作为是生产者 和 消费者 （协程的调度顺序是不确定的）
func ProduceAndConsumerSimulationWithSyncCond() {
	stopCh := make(chan struct{})

	lc := new(sync.Mutex)
	cond := sync.NewCond(lc)

	cup := cup{
		iceCubes: make([]iceCube, 3, 3),
	}

	// consumer
	go func() {
		for {
			cond.L.Lock()
			for len(cup.iceCubes) == min {
				cond.Wait()
			}
			// 删除头部的冰块
			cup.iceCubes = cup.iceCubes[1:]
			fmt.Println("consume 1 iceCube, left iceCubes ->  ", len(cup.iceCubes))
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// producer
	go func() {
		for {
			cond.L.Lock()
			for len(cup.iceCubes) == max {
				cond.Wait()
			}
			// 杯子中新添加进一个冰块.
			cup.iceCubes = append(cup.iceCubes, 1)
			fmt.Println("producer 1 iceCube, left iceCubes ", len(cup.iceCubes))
			cond.Signal()
			cond.L.Unlock()
		}

	}()

	for {
		select {
		case <-stopCh:
			return
		default:
		}
	}
}
