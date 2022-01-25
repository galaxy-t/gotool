package _advanced

import (
	"fmt"
	"reflect"
)

/*

	反射
	反射是指程序运行期对程序本身进行访问和修改的能力.

	// 结构体反射
	任意值通过 reflect.TypeOf() 获得反射对象信息后, 如果它的类型是结构体, 可以通过反射值对象
	(reflect.Type) 的 NumField() 和 Field() 等方法获得结构体成员的详细信息.

	反射是一个强大并富有表现力的工具, 能让我们写出更灵活的代码, 但是反射不应该被滥用, 原因如下:
		1. 基于反射的代码是极其脆弱的, 反射中的类型错误会在真正运行的时候才会引发 panic, 那很可能会是在代码写完的很长时间之后
		2. 大量使用反射的代码通常难以理解
		3. 反射的性能低下, 基于反射实现的代码通常会比正常代码运行速度慢一到两个数量级


*/

type Dog21 struct {
}

// F191 使用 refelct.TypeOf 得到变量的类型
func F191() {

	x := 111
	t := reflect.TypeOf(x)
	fmt.Println(t.Name()) // int   Name 是类型信息
	fmt.Println(t.Kind()) // int	 Kind 是种类信息

	fmt.Println()

	var d Dog21
	t1 := reflect.TypeOf(d)
	fmt.Println(t1.Name()) // Dog21
	fmt.Println(t1.Kind()) // struct

	// 注: Go 语言中的反射中, 像 数组,切片,Map,指针 等类型的变量, 他们的 .Name() 都是返回 空

}

// F192 通过 reflect.ValueOf() 获取变量的值
func F192() {

	x := 222
	v := reflect.ValueOf(x)
	fmt.Println(v)        // 得到值反射变量
	fmt.Println(v.Kind()) // 得到值类型
	fmt.Println(v.Int())  // 根据类型得到具体的值, 得到 int 类型的值

	fmt.Println()

	// 使用 switch 判断多种类型
	switch v.Kind() {
	case reflect.String:
		fmt.Println(v.String())
	case reflect.Float32:
		fmt.Println(float32(v.Float()))
	case reflect.Int32:
		fmt.Println(int32(v.Int()))
	case reflect.Int:
		fmt.Println(v.Int())

	}

}

// F193 设置某个变量的值
func F193() {

	x := 111
	fmt.Println(x)
	v := reflect.ValueOf(&x)

	// Elem() 用来根据指针取对应的值
	v.Elem().SetInt(222)
	fmt.Println(x)

}

// F194 反射值检查及获取结构体中的属性的反射值
func F194() {

	x := true
	v := reflect.ValueOf(x)
	fmt.Println(v.IsZero())  // 判断 v 持有的值是否为该类型的零值(初始值/默认值)
	fmt.Println(v.IsValid()) // 检查 v 是否持有一个值, 如果 v 是 Value 的零值会返回 false, 此时 v 除了 isValid,String,Kind 之外的方法都会导致 panic
	// fmt.Println(v.IsNil())		    // 检查 v 持有的值是否为 nil, v 持有的值的分类必须是 通道,函数,接口,映射,指针,切片 之一, 否则该函数会导致 panic

	s := struct{}{}                               // 声明一个空的结构体
	v1 := reflect.ValueOf(s)                      // 得到 s 的反射值
	fmt.Println(v1.FieldByName("name").IsValid()) // 检查属性 name 是否存在

}

type Student21 struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	age   int
}

// F195 获取结构体中的属性信息
func F195() {

	s1 := Student21{"张三", 12, 13}
	t1 := reflect.TypeOf(s1)

	fieldCount := t1.NumField() // 3  通过反射得到结构体有多少个属性, 包括私有的

	// 遍历结构体的所有变量
	for i := 0; i < fieldCount; i++ {
		sf := t1.Field(i) // 根据下标获取属性
		name := sf.Name   // 获取字段名
		type1 := sf.Type  // 获取类型
		tag := sf.Tag     // 获取 tag

		fmt.Printf("name: %v  type: %v   tag: %v \n", name, type1, tag)
		fmt.Println(tag.Get("json")) // 获取某一个 tag 值
	}
}

// F196 通过反射调用结构体的方法
// 感觉 Go 语言的编程模式, 这种反射方法的情况可能很少
func F196() {
}
