package example

import "fmt"

/*

	变量作用域

	1. 内部可以访问到外部的变量
	2. 外部不能访问到内部的变量
*/

// 定义全局变量
var num = 10

// F091 调用变量
func F091() {

	num := 100

	// 函数内部访问变量
	// 优先访问函数内部的变量
	// 找不到内部的变量再去找外部的变量
	fmt.Println("全局变量", num)

}

func F092()  {
	// 函数是可以被作为参数
	abcd := F21
	fmt.Printf("%T\n", abcd)		// func()
	abcd()
}

// F093 匿名函数
func F093()  {

	// 1. 匿名函数
	f1 := func(a,b int) string {
		return "sdfsd"
	}
	result1 := f1(1,2)
	fmt.Println(result1)

	// 2. 定义并立即执行的匿名函数
	result2 := func(a,b int) string {
		fmt.Println("sdfsd")
		return "111111"
	}(1,2)

	fmt.Println(result2)
}

// F094 定义一个函数, 它的返回值是一个函数
func F094(name string) func() {
	// name := "zhangsan"
	return func() {
		fmt.Println("Hello", name)	// 引用了函数外层变量
	}
}

// F095 闭包
// 闭包 = 函数 + 函数对外层变量的引用
func F095() {
	r := F094("zhangsan")	// 此时 r 就是一个闭包, 闭包内部引用的变量不是自己声明的, 而是外部声明后传入的
	r()		// 相当于执行了 F24() 函数内部的匿名函数
}