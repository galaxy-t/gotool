package _advanced

import (
	"fmt"
	"sync"
)

/*

	同步与锁

	互斥锁 sync.Mutex
	是一种常用的控制共享资源访问的方法, 它能够保证同时只有一个 goroutine 可以访问共享资源.
	使用互斥锁能够保证同一时间有且只有一个 groutine 进入临界区, 其它的 goroutine 则在等待锁;
	当互斥锁释放后, 等待的 groutine 才可以获取锁进入临界区, 多个goroutine 同时等待一个锁时候, 唤醒的策略是随机的.

	读写互斥锁 sync.RWMutex
	互斥锁是完全互斥的, 但是在很多情况下, 并发的读取一个资源不涉及资源修改的时候是没有必要枷锁的.
	读写锁,分为两种: 读锁和写锁
	当一个 goroutine 获取读锁之后, 其它的 goroutine 如果获取读锁会继续获得锁, 如果是获取写锁就会等待;
	当一个 goroutine 获取写锁之后, 其它的 goroutine 无论获取读锁还是写锁都会等待.

	注: 读写锁适合读多写少的场景, 如果读和写的操作差别不大, 读写锁的优势就发挥不出来.

	线程计数器 sync.WaitGroup

	在高并发的场景下只执行一次 sync.Once

	并发安全的Map, sync.Map

	原子操作
	代码中的加锁操作因为设计内核态的上下文切换会比较耗时, 代价比较高.
	针对基础数据类型我们还可以使用原子操作来保证并发安全.
	因为原子操作是 Go 语言提供的方法, 它在用户态就可以完成, 因此性能比夹所操作更好.
	Go 语言中的原子操作由内置的 sync/atomic 提供.

*/

var (
	x       int64
	wg1     sync.WaitGroup
	lock    sync.Mutex // 互斥锁
	addOnce sync.Once  // 只加载一次
)

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 针对 x 的操作进行加锁
		x = x + 1
		lock.Unlock() // 针对 x 的操作进行释放锁
	}
	wg1.Done()
}

func F221() {

	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	fmt.Println(x)

}

func add1() {
	x++
}

func F222() {
	for i := 0; i < 10; i++ {
		addOnce.Do(add1)
	}
	fmt.Println(x) // x 的最终结果为 1, 因为 addOnce.Do 只会被执行一次
}

var m sync.Map

func F223() {

	for i := 0; i < 20; i++ {
		wg1.Add(1)
		go func(i int) {
			m.Store(i, i+100)
			v, _ := m.Load(i)
			fmt.Println(v)
			wg1.Done()
		}(i)
	}
	wg1.Wait()
}
