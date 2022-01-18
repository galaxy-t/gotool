package example

import "fmt"

/*

	数组

*/

// F051 数组的定义
func F051() {

	// 数组的定义
	var a [3]int
	var b [4]int
	fmt.Println(a) // [0 0 0]
	fmt.Println(b) // [0 0 0 0]

	// 数组的初始化
	// 1. 定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)
	fmt.Println(cityArray[0])
	fmt.Println(cityArray[3])

	// 2. 编译器推导数组的长度, 数组的长度一定要是一个确定的值
	var boolArr = []bool{true, false, false, true, false}
	fmt.Println(boolArr)
	var boolArr1 = [...]bool{true, false, false, true, false}
	fmt.Println(boolArr1)

	// 3. 使用索引值的方式进行初始化
	var langArr = [...]string{1: "Golang", 3: "Python", 7: "Java"}
	fmt.Println(langArr)
	fmt.Printf("%T\n", langArr) // 打印数组类型  [8]string

}

// F052 数组的遍历
func F052() {

	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}

	// 1. 使用 for 循环遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	// 2. for range 遍历
	for index, value := range cityArray {
		fmt.Println(index, value)
	}

	// 3. 数组的传递为值传递
	x := [3]int{1, 2, 3}
	f12(x)
	fmt.Println(x)

}

func f12(a [3]int) {
	a[0] = 100
}