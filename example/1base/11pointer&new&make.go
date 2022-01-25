package _base

import "fmt"

/*

	指针

	Go 语言中的函数传参都是值拷贝
	当我们想要修改某个变量的时候, 我们可以创建一个指向该变量地址的指针变量
	传递数据使用指针, 而无须拷贝数据
	类型指针不能进行偏移和运算
	Go 语言中的指针操作非常简单, 只有两个符号: &(取地址) 和 *(根据地址取值)


	new&make
	new 是一个内置的函数, 函数签名: func new(Type) *Type
			Type 表示类型, new函数只接受一个参数, 这个参数是一个类型
			*Type 表示类型指针, new 函数返回一个指向该类型内存地址的指针
	make 也是用于内存分配的, 区别于 new, 它只用于 slice, map以及 chan 的内存创建, 而且它返回的类型就是这三个类型本身, 而不是他们的指针类型
	因为这三种类型就是引用类型, 所以就没有必要返回他们的指针了
	make 函数的签名函数为: func make(t Type, size ...IntegerType) Type
	make 函数是无可替代的, 我们在使用 slice, map 以及 channel 的时候, 都需要使用 make 进行初始化, 然后才可以对他们进行操作

	new 和 make 的区别
	1. 二者都是用来做内存分配的
	2. make 只用于 slice, map 以及 channel 的初始化, 返回的还是这三个引用类型本身
	3. 而 new 用于类型的内存分配, 并且内存对应的值为类型零值, 返回的是指向类型的指针


*/

// F111 指针地址和指针类型
// 每个变量在运行时都拥有一个地址, 这个地址代表变量在内存中的位置.
// Go 语言中使用 & 字符放在变量前面对变量进行 "取地址" 操作.
// Go 语言中的值类型(int, float, bool, string, array, struct) 都有相应的指针类型, 如: *int, *int64, *string 等
func F111() {

	a := 1                       // a 是一个 int 类型的变量, 其值为 1
	b := &a                      // 通过 & 将 a 的指针赋值给 b, 此时 b 是 a 的指针, 类型为指针类型
	fmt.Printf("%v %p\n", a, &a) // 0xc00000a300 其值是一个十六进制的值
	fmt.Println(*b)              // 通过 * 将 b 这个指针指向的变量的值拿出来

	// 	ptr := &v		// v 的类型为 T
	// v: 代表被取地址的变量, 类型为 T
	// ptr: 用于接收地址的变量, ptr 的类型就是 *T, 称为 T 的指针类型. * 代表指针

	// 对变量进行取地址 (&) 操作, 可以获得这个变量的指针变量
	// 指针变量的值是指针地址
	// 对指针变量进行取值 (*) 操作, 可以获得指针变量指向的原变量的值
}

func d(x int) {
	x = 100
}

func e(y *int) {
	*y = 100
}

// F112 值传递和指针传递
func F112() {
	a := 1
	d(a)
	fmt.Println(a)
	e(&a)
	fmt.Println(a)
}

// F113 new&make
func F113() {

	var a *int      // 声明一个 *int 类型的指针
	a = new(int)    // 为 a 进行初始化
	fmt.Println(*a) // 初始值为 0
	*a = 100        // 赋值
	fmt.Println(*a) // 取值

	var b map[string]int     // 声明一个 map 类型的变量 b
	b = make(map[string]int) // 使用 make 为 b 进行初始化
	fmt.Println(b)           // 是一个空 map, 而不是 nil
	b["a"] = 100             // 赋值
	fmt.Println(b)           //

}
