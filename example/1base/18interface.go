package _base

import "fmt"

/*

	接口(interface) 定义了一个对象的行为规范, 只定义规范不是先, 由具体的对象来实现规范的细节.

	在 Go 语言中接口(interface) 是一种类型, 一种抽象的类型.

	定义:
	每个接口由数个方法组成, 接口的定义格式如下:
	type 接口类型名 interface {
		方法名1( 参数列表1 ) 返回值列表1
		方法名2( 参数列表2 ) 返回值列表2
	}
	其中:
		1. 接口名: 使用 type 将接口定义为自定义的类型名. Go 语言的接口在命名时, 一般会在单词后面添加 er,
			如有写操作的接口叫 Writer, 有字符串功能的接口叫 Stringer 等, 接口名最好要能突出该接口的类型含义
		2. 方法名: 当方法名首字母是大写且这个接口类型名首字母也是大写时, 这个方法可以被接口所在的包(package)之外的代码访问
		3. 参数列表,返回值列表: 参数列表和返回值列表中的参数变量名可以省略

	实现接口的条件:
	一个对象只要全部实现了接口中的方法, 那么就实现了这个接口

	接口的嵌套

	空接口
	空接口是指没有定义任何方法的接口, 因此任何类型都实现了空接口.
	空接口类型的变量可以存储任意类型的变量
	空接口一般不需要提前定义

	空接口的应用
	1. 空接口类型作为函数的参数, 类似于 java 中的 Object
	2. 空接口可以作为 map 的 value

	接口断言
	想要判断空接口中的值的时候可以使用类型断言, 其语法如下
	x.(T)
	其中:
		1. x: 表示类型为 interface{} 的变量
		2. T: 表示断言, x 可能是的类型
	该语法返回两个参数, 第一个参数是 x 转化成为 T 类型后的变量, 第二个值是一个布尔值, 若为 true 则表示断言成功,
	若为 false 则表示断言失败




*/

type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("啊啊啊~")
}

// 定义一个接口
// 接口不管是什么类型, 只管要实现什么方法
// 定义一个类型, 一个抽象的类型, 只要实现了 say() 这个方法的类型都可以成为 sayer 类型
type sayer interface {
	say()
}

// 打的函数
func da(arg sayer) {
	arg.say() // 不管是传入什么参数, 都要调用它的 say 方法
}

func F181() {

	c1 := cat{}
	da(c1)

	d1 := dog{}
	da(d1)

	p1 := person{"张三"}
	da(p1)

	var s sayer // 声明了一个 sayer 接口的变量
	c2 := cat{} // 声明了一个 cat 类型的变量
	s = c2      // 将 c2 赋值给 s 是不会报错的
	fmt.Println(s)

}

type mover interface {
	move()
}

type person1 struct {
	name string
	age  int
}

func (p person1) move() {
	fmt.Printf("%s在跑...", p.name)
}

// F182 使用值接收者实现接口: 类型的值和类型的指针都能够保存到接口变量中
func F182() {
	var m mover
	p1 := person1{"张三", 18}  // 值类型
	p2 := &person1{"李四", 18} // 引用类型
	m = p1
	m = p2
	m.move()

}

type eater interface {
	eat()
}

type person2 struct {
	name string
	age  int
}

func (p *person2) eat() {
	fmt.Printf("%s在吃...", p.name)
}

// F183 使用指针接收者实现接口: 类型的值不能保存到接口变量中, 类型的指针可以保存到接口变量中
func F183() {

	var e eater
	// p1 := person2{"张三", 18}		// 值类型
	p2 := &person2{"李四", 18} // 指针类型

	// e = p1	// 使用指针类型来接收接口的方法, 那么这个对象的值类型就不能够被保存到接口变量中
	e = p2
	e.eat()
}

// 接口的嵌套, 相当于实现了 sayer和eater 的所有的接口方法
type animal interface {
	sayer
	eater
}

// F184 空接口, 任何类型的变量都是空接口的实例
func F184() {
	var x interface{}       // 定义一个空接口变量 x
	p1 := person2{"张三", 18} // 定义一个 person2 的变量
	x = p1                  // 将 p1 赋值给 x
	fmt.Println(x)
	x = 1 // 将一个 int 赋值给 x
	fmt.Println(x)
	x = "哈哈" // 将一个字符串赋值给 x
	fmt.Println(x)
}

// F185 接口断言
func F185() {

	var x interface{}

	x = 1
	ret, ok := x.(int)
	fmt.Println(ret, ok)

	ret1, ok1 := x.(string) // 若 ok1 为 false 时候, 则 ret1 则为一个 string 类型的 零值, 即 空字符串
	fmt.Println(ret1, ok1)

	// 使用 switch 进行类型断言
	// v 是转换后的值
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型, value: %v\n", v)
	case int:
		fmt.Printf("是整型, value: %v\n", v)
	}

}
