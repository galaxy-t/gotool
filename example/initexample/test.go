package test

import "fmt"

// Name 公共变量, 可以被其他包调用
var Name string = "张三"

// 私有变量, 仅当钱包调用
var name string = "zhangsan"

// 在包导入的时候自动执行
// 该函数没有参数也没有返回值
func init()  {
	fmt.Println("test init()......")
	fmt.Println(name)
}


// Hello 公共函数, 可以被其他包调用
func Hello()  {
	fmt.Println("Hello")
}

// 私有函数, 仅当前包使用
func hello()  {
	fmt.Println("hello")
}