package _base

import "fmt"

/*

	结构体的匿名字段

*/

// Person 结构体的匿名字段
// 字段只有类型, 没有名字
// 匿名字段同一种类型只能出现一次
type Person struct {
	string
	int
}

func F151() {

	// 因为匿名字段是没有名字的, 所以初始化赋值要求按照顺序进行赋值
	p1 := Person{"张三", 18}
	// 取值需要使用匿名字段的类型来进行取值
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	// 赋值也需要使用类型来进行赋值
	p1.int = 1
}

// 嵌套结构体

type Address struct {
	Province string
	City     string
}
type Person1 struct {
	Name    string
	Address Address
	// Address	// 此处可以直接使用匿名结构体
}

// F152 使用嵌套结构体
func F152() {

	p1 := Person1{
		Name: "张三",
		Address: Address{
			Province: "山东",
			City:     "青岛",
		},
	}

	fmt.Println(p1.Name, p1.Address.Province, p1.Address.City)
	// 使用匿名结构体可以直接访问其属性
	// 如果嵌套匿名结构体包含多个同名的字段, 那么就必须使用指定其类型来调用其字段
	// fmt.Println(p1.Name, p1.Province, p1.City)

}

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%s会动~\n", a.name)
}

type Dog struct {
	Feet   int8
	Animal // 匿名嵌套, 而且嵌套的是一个结构体指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

// F153 结构体的继承
func F153() {

	d1 := &Dog{
		Feet: 4,
		Animal: Animal{
			name: "饭团",
		},
	}
	// 此时 Dog 类型的变量是可以调用其匿名结构体的方法
	d1.wang()
	d1.move()
}
