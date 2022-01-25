package _base

import "fmt"

/*

	结构体
	是一种自定义数据类型, 可以封装多个基本数据类型, 类似于 Java 中的 POJO

	使用 type 和 struct 关键字来定义结构体
	type 类型名 struct {
		字段名 字段类型
		字段名 字段类型
		...
	}
	类型名: 标识自定义结构体的名称, 在同一包内不能重复
	字段名: 表示结构体字段名, 结构体中的字段名必须唯一
	字段类型: 表示结构体字段的具体类型

*/

// 定义一个名称为 user 的结构体
// 尽量将同一类型的字段挨着排列, 参考 Go 语言中的内存对齐: https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
type user struct {
	name    string
	address string
	age     int
	sex     int8
}

// NewUser 结构体的构造函数
// Go 语言的结构体没有构造函数, 我们可以自己实现
// 因为 struct 是值类型(值传递), 如果结构体比较复杂的话, 值拷贝的性能开销会比较大, 所以建议构造函数返回的是结构体指针类型
func NewUser(name string, address string, age int, sex int8) *user {
	return &user{name, address, age, sex}
}

// F131 结构体实例化 - 基本实例化
// 只有当结构体实例化时, 才会真正地分配内存. 也就是必须实例化后才能使用结构体字段.
// 结构体本身也是一种类型, 我们可以像声明内置类型一样使用 var 关键字声明结构体类型.
func F131() {

	var u user // 此时已经进行实例化, p=edu04.user{name:"", address:"", age:0, sex:0}
	fmt.Printf("p=%#v\n", u)
	u.name = "张三"
	u.address = "青岛"
	u.age = 18
	u.sex = 1
	fmt.Printf("p=%#v\n", u) // p=edu04.user{name:"张三", address:"青岛", age:18, sex:1}

	fmt.Println(u.name)
	fmt.Println(u.address)
	fmt.Println(u.age)
	fmt.Println(u.sex)
}

// F132 结构体实例化 - 匿名结构体
// 在定义一些临时数据结构等场景下还可以使用匿名结构体
func F132() {

	var person struct {
		name    string
		married bool
	}
	fmt.Printf("p=%#v\n", person) // p=struct { name string; married bool }{name:"", married:false}
	person.name = "李四"
	person.married = true
	fmt.Printf("p=%#v\n", person) // p=struct { name string; married bool }{name:"李四", married:true}

}

// F133 创建指针类型结构体
// 我们还可以通过 new 关键字对结构体进行实例化, 得到的是结构体的地址
func F133() {

	var u = new(user)
	fmt.Printf("%T\n", u)  // *edu04.user 此时 u 是一个指针类型
	fmt.Printf("%#v\n", u) // &edu04.user{name:"", address:"", age:0, sex:0}

	// 以下写法其实是一样的
	(*u).sex = 12
	u.name = "sdalf" // 语法糖, 可以直接通过指针来操作其属性

}

// F134 取结构体的地址进行实例化
func F134() {

	u := user{}
	fmt.Printf("%T\n", u)  // edu04.user
	fmt.Printf("%#v\n", u) // edu04.user{name:"", address:"", age:0, sex:0}

	u1 := &user{}
	fmt.Printf("%T\n", u1)  // *edu04.user
	fmt.Printf("%#v\n", u1) // &edu04.user{name:"", address:"", age:0, sex:0}
}

// F135 结构体的初始化 - 键值对初始化
func F135() {
	u := user{
		name: "张三",
		// address: "青岛",  允许不给某一个属性初始化
		age: 18,
		sex: 1,
	}
	fmt.Printf("%T\n", u)  // edu04.user
	fmt.Printf("%#v\n", u) // edu04.user{name:"张三", address:"青岛", age:18, sex:1}

	u2 := &user{
		name:    "张三",
		address: "青岛",
		age:     18,
		sex:     1,
	}
	fmt.Printf("%T\n", u2)  // *edu04.user
	fmt.Printf("%#v\n", u2) // &edu04.user{name:"张三", address:"青岛", age:18, sex:1}
}

// F136 结构体的初始化 - 值列表进行初始化
func F136() {

	// 必须初始化结构体的所有字段
	// 初始值的填充顺序与字段在结构体中的声明顺序一致
	// 该方式不能和键值初始化方式混用

	u := user{"张三", "青岛", 18, 1}
	fmt.Printf("%T\n", u)  // edu04.user
	fmt.Printf("%#v\n", u) // edu04.user{name:"张三", address:"青岛", age:18, sex:1}

	u2 := &user{"张三", "青岛", 18, 1}
	fmt.Printf("%T\n", u2)  // *edu04.user
	fmt.Printf("%#v\n", u2) // &edu04.user{name:"张三", address:"青岛", age:18, sex:1}
}
