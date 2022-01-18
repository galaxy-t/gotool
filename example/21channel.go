package example

import (
	"fmt"
	"time"
)

/*

	channel
	channel 是 goroutine 之间的连接, 可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制
	channel 是一种特殊的类型, 类似于队列, 遵循先入先出(FIFO)的规则, 保证收发数据的顺序
	每一个 channel 都要求指定元素类型

	实现
	var 变量 chan 元素类型

	实现之后必须使用 make 函数进行初始化(总共只有 map,slice,channel 三种必须要使用make进行初始化)
	make(chan 元素类型,[缓冲大小])

	操作
	通道有 发送(send),接收(receive)和关闭(close) 三种操作
	发送和接收都是用 <- 符号
	关闭
	我们可以通过内置的 close 函数来关闭通道
	close(通道变量名)
	对于关闭通道需要注意的事情是, 只有在通知接收方 goroutine 所有的数据都发送完毕的时候才需要关闭通道,
	通道是可以被垃圾回收机制回收的, 它和关闭文件是不一样的, 在结束操作之后关闭文件是必须的, 但关闭通道不是必须的

	关闭后的通道有以下特点
		1. 对一个关闭的通道再发送值会导致 panic
		2. 对一个关闭的通道将进行接收会一直获取值直到通道为空
		3. 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值
		4. 关闭一个已经关闭的通道会导致 panic

	注: 向一个通道发送值, 发送完了立刻关闭该通道, 另一个 goroutine 在读取该通道值的时候是可以继续取值的, 另一个 goroutine 取值方式如下:
		tmp, ok := <-ch1
		第一个返回值是读取到的值, 第二个返回值是个布尔类型, 当通道中的值读取完毕并且该通道已经被关闭时, 会返回 false, 用于标识通讯结束

	无缓冲区通道(同步通道)
	即创建通道的时候不设置缓冲区大小, 如下:
	ch := make(chan int)
	此时该通道是阻塞的, 即向该通道发送值后就会被阻塞, 直到有另一个 goroutine 读取该通道才会被真正的写入通道

	带缓冲区通道(异步通道)
	即创建通道的时候设置缓冲区大小, 如下:
	ch := make(chan int, 1)
	向通道发送值的时候直接放到通道的缓冲区, 不需要考虑接收方是否收到该值, 是异步的

	多路复用
	用于同时读取多个通道
	select {
	case <-ch1:
		...
	case data := <-ch2:
		...
	case ch3 <- data:
		...
	default:
		默认操作
	}
	select 的使用类似于 switch 语句, 它有一些 case 分支和一个默认的分支. 每个 case 对应一个通道的通信(接收或发送)过程.
	select 会一直等待, 直到某个 case 的通信操作完成时, 就会执行 case 分支对应的语句
	使用 select 语句能提高代码的可读性
		1. 可处理一个或多个 channel 的 发送/接收 操作
		2. 如果多个 case 同时满足, select 会随机选择一个
		3. 对于没有 case 的 select{} 会一直等待, 可用于阻塞 main() 函数


*/

// F211 创建 channel
func F211() {

	var ch1 chan int   // 声明一个传递整型的通道
	var ch2 chan bool  // 声明一个传递布尔类型的通道
	var ch3 chan []int // 声明一个 int切片的通道

	fmt.Println(ch1) // nil
	fmt.Println(ch2) // nil
	fmt.Println(ch3) // nil

	ch1 = make(chan int)   // 实例化
	ch2 = make(chan bool)  // 实例化
	ch3 = make(chan []int) // 实例化

	fmt.Println(ch1) // nil
	fmt.Println(ch2) // nil
	fmt.Println(ch3) // nil
}

// F212 channel 操作
func F212() {

	// ch := make(chan int)		// 无缓冲区通道

	ch := make(chan int, 1) // 定义一个通道, 设置缓冲区为 1

	ch <- 10             // 向通道中发送一个 10
	fmt.Println(len(ch)) // 获取通道中元素的数量
	fmt.Println(cap(ch)) // 获取通道的容量

	i := <-ch // 从通道中接收值并赋值给变量
	// <- ch						// 从通道中接收值, 忽略结果

	fmt.Println(i)

	close(ch) // 关闭通道
}

// 1. 生成 0-100的数字发送到 ch1
// 在 chan 和 int 的中间添加 <- 符号, 用于标识这是一个单向只写通道, 在该函数中, 只允许往 ch 中发送值, 不允许读取 ch 中的值
func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// 2. 从 ch1 中取出数据计算它的平方, 把结果发送到 ch2 中
// 在 通道名 和 chan 中间添加 <- 符号, 用于标识这是一个只读单向通道, 在该函数中, 只允许从 ch1 中读取值, 不允许往 ch1 中写入值
func f2(ch1 <-chan int, ch2 chan<- int) {

	// 1. 无限循环从通道中取值
	for {
		tmp, ok := <-ch1
		fmt.Println("************", tmp, ok)
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}

	close(ch2)
}

// F213 channel 使用
// 两个 goroutine
//    1. 生成 0-100的数字发送到 ch1
//    2. 从 ch1 中取出数据计算它的平方, 把结果发送到 ch2 中
func F213() {

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	// 2. 使用 for range 来从通道取值
	for ret := range ch2 {
		fmt.Println(ret)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		fmt.Println("start worker:", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
	}

}

// F214 WorkPool
func F214() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启 3 个 goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 发送 5 个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	// 输出结果

	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}

// F215 多路复用
func F215() {

	// ch := make(chan int, 1)	// 通道缓冲区大小为 1, 那么此时该函数的结果应该是 0 2 4 6 8
	ch := make(chan int, 10)	// 通道缓冲区大小为 10, 通道是可以存在多个值的, 但是 select 在符合多个 case 的时候只会随机选择一个, 那么结果就变的不再确定

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("......")
		}
	}

}