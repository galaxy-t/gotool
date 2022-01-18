package example

import "fmt"

/*
	map
	引用类型, 默认为 nil , 和切片一样, 必须初始化才能使用
*/

// F071 定义
func F071() {

	a := map[string]int{} // 该种方式必须得使用花括号 {} 进行初始化
	fmt.Println(a)

	a1 := map[string]int{ // 声明一个 map 同时完成初始化, 并添加键值对
		"hello": 1,
		"world": 2,
	}
	fmt.Println(a1)

	var b map[string]int // 只声明了一个 map类型, 但是没有进行初始化, a 的初始值就是 nil
	fmt.Println(b == nil)

	// map 的初始化
	// 第一个参数为 类型, 第二个参数为初始化容量
	b = make(map[string]int, 10)
	fmt.Println(b == nil)

	b["hello"] = 1
	b["world"] = 2

	fmt.Println(b)
	fmt.Printf("type: %T\n", b)

}

// F072 判断键值对是否存在
func F072() {
	scoreMap := make(map[string]int, 8)
	scoreMap["zhangsan"] = 100
	scoreMap["lisi"] = 200

	value, ok := scoreMap["wanger"]
	if ok {
		fmt.Println("存在--", value)
	} else {
		fmt.Println("不存在")
	}
}

// F073 map 的遍历
func F073() {

	scoreMap := make(map[string]int, 8)
	scoreMap["zhangsan"] = 100
	scoreMap["lisi"] = 200

	// 遍历的顺序和添加的顺序是无关的

	// 1. 遍历键值对
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 2. 只遍历 key
	for k := range scoreMap {
		fmt.Println(k)
	}

	// 3. 按照某个固定顺序遍历 map
	// 3.1 先通过 for 循环取出所有的 key 存放到 切片 中
	// 3.2 对 key 进行排序	sort.Strings()
	// 3.3 按照排序后的 key 对 map 进行取值

}

// F074 删除键值对
func F074() {

	scoreMap := make(map[string]int, 8)
	scoreMap["zhangsan"] = 100
	scoreMap["lisi"] = 200

	delete(scoreMap, "zhangsan")

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}
