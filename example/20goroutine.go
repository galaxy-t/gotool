package example

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*

	并发: 同一时间段内执行多个任务(在用微信和两个朋友聊天)
	并行: 同一时刻执行多个任务(两个人和自己的一个朋友聊天)

	Go 语言的并发通过 goroutine 实现. goroutine 类似于线程, 属于用户态的线程, 我们可以根据需要创建成千上万个 goroutine 并发工作.
	goroutine 是由 Go 语言的运行时 {runtime} 调度完成, 而线程是由操作系统调度完成.

	Go 语言还提供了 channel 在多个 goroutine 间进行通信. goroutine 和 channel 是 Go 语言秉承 CSP(communicating sequential processes) 并发模式的重要实现基础

	1. 一个操作系统线程对应用户态多个 goroutine
	2. go 程序可以同时使用多个操作系统线程. 即 Go 程序中的所有的 goroutine 可以分不到不同的 CPU 核心中去执行
	3. goroutine 和 OS 线程是 多对多的关系, 即 m 个 goroutine 分不到 n 个操作系统线程上

*/

// 定义一个线程计数器
// 注: 不要尝试把创建好的线程计数器当作参数传入到并发方法中, 这样会造成死锁
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("Hello world!", i)
	wg.Done() // 方法执行完了, 线程计数器减一
}

// F201 实现并发
func F201() {

	wg.Add(20000) // 线程计数器, 设置为 1

	// 只需要在要调用的函数前面加上 go 关键字, 一个 goroutine 必须对应一个函数, 可以创建多个 goroutine 去执行相同的函数
	for i := 0; i < 10000; i++ {
		go hello(i)
	}

	// 并发执行匿名函数
	// 闭包的时候参数应该是传入的, 而不是直接调用外部参数
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("这里是匿名函数", i)
			wg.Done()
		}(i)
	}

	fmt.Println("Hello F31") // 这一句话的打印不一定会在上一句的后面

	// time.Sleep(time.Second * 3)		// 让程序等待三秒
	wg.Wait() // 等待线程计数器归零

}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func F202() {

	// 只占用一个 CPU 核心
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)

	// 后果是先执行完一个函数, 再执行完另一个函数, 因为只有一个核心
}
