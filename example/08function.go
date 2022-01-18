package example

import "fmt"

/*

	函数

*/

// F081 函数的声明
// 不需要参数也没有返回值的函数
func F081() {
	fmt.Println("Hello world")
}

// F082 需要一个参数没有返回值的函数
func F082(a int) {

}

// F083 需要多个参数没有返回值的函数
func F083(a, b int, c string) {

}

// F084 需要一个返回值的函数
func F084() int {
	return 1
}

// F085 需要多个返回值的函数
func F085() (int, int, string) {
	return 1, 2, "sdfsd"
}

// F086 多个被命名过的返回值
// 若返回值已经被命名, 那么函数内部可以不再需要声明该变量, 最后的 return 也不再需要特别声明要返回哪个变量
func F086() (key int, value string) {
	key = 1
	value = "sdfsd"
	return
}

// F087 接收可变参数
// 此时参数 a 为一个 切片
func F087(a ...int) {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

// F088 同时拥有固定参数和可变参数
// 一个函数只允许拥有一个可变参数, 可变参数必须要放到固定参数后面
func F088(a string, b ...int) {

}

// F089 defer 的使用
// 延迟执行
// 在方法即将返回之前执行 defer 语句
// 最先被声明的 defer 最后被执行, 最后被声明的 defer 最先被执行
func F089() {

	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end ...")

}
