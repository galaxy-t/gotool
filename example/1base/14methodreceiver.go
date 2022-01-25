package _base

import (
	"fmt"
	"strconv"
)

/*

	方法和接收者
	Go 语言中的 方法(Method) 是一种作用于特定类型变量的函数
	这种特定类型变量叫做 接收者(Receiver)
	接收者的概念就类似于其他语言中的 this 或者 self

	方法的定义格式如下:
	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	}
	其中:
		接收者变量: 接收者中的参数变量名在命名时, 官方建议使用接收者类型名的第一个小写字母, 而不是 self,this 之类的命名.
				   例如 Person 类型的接收者变量应该命名为 p, Connector 则应该命名为 c
		接收者类型: 接收者类型和参数类型, 可以是指针类型和非指针类型

*/

// Car 车辆
type Car struct {
	color string // 车辆颜色
	speed int    // 速度
}

// NewCar 车辆的构造函数
func NewCar(color string, speed int) *Car {
	return &Car{color, speed}
}

// Run 车辆运行 - 指针接收者
// 因为方法接收者是一个指针类型, 所以我们在方法内部可以任意修改接收者的成员变量, 在方法结束后修改是有效的
func (c *Car) Run() {
	c.color = "绿色"
	fmt.Println(c.color + "---" + strconv.Itoa(c.speed))
}

func F141() {
	car := NewCar("红色", 28)
	fmt.Printf("%T\n", car)
	car.Run()
	fmt.Printf("%#v\n", car)
	// 语法糖, 实际上 NewCar 无论返回的是 值类型还是指针类型, 都不影响 car.Run() 的调用
}

// Run1 值接收者
// 当方法作用于值类型接收者时, Go 语言会在代码运行时将接收者的值赋值一份.
// 在值类型接收者的方法中可以获取接收者的成员值, 但修改操作只是针对副本, 无法修改接收者变量本身.
func (c Car) Run1() {
	c.color = "白色"
	fmt.Println(c.color + "---" + strconv.Itoa(c.speed))
}

func F142() {
	car := NewCar("红色", 28)
	fmt.Printf("%T\n", car)
	car.Run1()
	fmt.Printf("%#v\n", car)
	// 语法糖, 实际上 NewCar 无论返回的是 值类型还是指针类型, 都不影响 car.Run() 的调用
}

// 什么时候应该使用指针类型
// 1. 需要修改接收者中的值
// 2. 接收者是拷贝代价比较大的对象
// 3. 保持一致性, 如果又某个方法使用了指针接收者, 那么其他的方法也应该使用指针接收者

// 任意类型添加方法
// 在 Go 语言中, 接收者的类型可以是任何类型, 不仅仅是结构体, 任何类型都可以拥有方法,
// 例: 我们基于内置的 int 类型使用 type 关键字可以定义新的自定义类型, 然后为我们的自定义类型添加方法
// 注: 非本地类型不能定义方法, 也就是我们不能给别的包的类型定义方法

type MyInt int

func (m MyInt) haha() {
	fmt.Println("haha")
}
func F143() {
	var m1 MyInt = 100
	m1.haha()
}
