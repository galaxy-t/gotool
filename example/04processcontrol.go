package example

import "fmt"

/*

	流程控制

*/

// F041 if 判断
func F041() {

	// 1. 基本写法
	var score = 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 2. 特殊写法, score1 的声明放在了 if 语句里面, 该变量只允许在 if 代码块中生效
	if score1 := 65; score1 >= 90 {
		fmt.Println("A")
	} else if score1 > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}

// F042 for 循环
func F042() {

	// 1. 基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2. 省略初始语句, 但是必须保留初始语句后面的分号
	j := 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	// 3. 省略初始语句和结束语句
	x := 10
	for x > 0 {
		fmt.Println(x)
		x--
	}

	// 4. 无限循环
	for {
		fmt.Println("hello")
	}

	// 5. break, continue

	// 6. for range (键位循环)
}

// F043 switch
func F043() {

	// 一个 switch 只能有一个 default
	// 一个 case 有多个值, 则使用逗号 ,  分割
	// case 中可以使用表达式,

	finger := 3
	switch finger {
	case 1, 11, 12, 13:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效输入")
	}

}