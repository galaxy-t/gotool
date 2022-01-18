package example

import "fmt"

/*

	内置函数

	close			主要用来关闭 channel
	len				用来求长度, 比如 string, array, slice, map, channel
	new				用来分配内存, 主要用来分配值类型, 如: int, struct, 返回的是指针
	make			用来分配内存, 主要用来分配引用类型, 如: chan, map, slice
	append			用来追加元素到数组, slice 中
	panic&recover 	用来做错误处理

*/

func a10() {
	fmt.Println("func a")
}
func b10() {

	// 在 panic 之前进行 defer, 否则不生效
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("func b error")
		}
	}()

	panic("panic in b")
}
func c10() {
	fmt.Println("func c")
}

// F101 panic&recover
// Go 语言中目前没有异常机制, 但是使用 panic/recover 模式来处理错误, panic 可以再任何地方引发,
// 但 recover 只有在 defer 调用的函数中有效.
// 注: 1. recover() 必须搭配 defer 使用. 2. defer 一定要在可能引发 panic 的语句之前定义
func F101() {

	a10()
	b10() // 此处会抛出异常
	c10()

}
