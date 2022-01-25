package _base

import "fmt"

/*

	类型别名和自定义类型
	在 Go 语言中有一些基本的数据类型, 如 string,int,float,bool 等数据类型, go 语言中可以使用 type 关键字来定义自定义类型

*/

// F121 自定义类型 是定义了一个全新的类型, 我们可以基于内置的基本类型定义, 也可以通过 struct 定义
func F121() {

	// 将 MyInt 定义为 int 类型
	// 通过 Type 关键字的定义, MyInt 就是一种新的类型, 它具有 int 的特性.
	type MyInt int

	var a MyInt = 1
	fmt.Printf("type: %T value: %v\n", a, a)

}

// F122 类型别名
// 类型别名是 Go1.9 版本添加的新功能
// 类型别名规定: TypeAlias 只是 Type 的别名. 本质上 TypeAlias 与 Type 是同一个类型.
func F122() {
	// 给 uint8 和 uint32 分别设置不同的类型别名
	type NewInt = uint8
	var b NewInt = 12
	fmt.Printf("type: %T value: %v\n", b, b)
}
