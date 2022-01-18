package example

import "fmt"

/*

	切片(Slice)
	切片是一个拥有相同类型元素的可变长度的序列
	它是基于数组类型做的一层封装
	它非常灵活, 支持自动扩容

	切片是一个引用类型, 它的内部结构包含 地址,长度,容量
	切片一般用于快速的操作一块数据集合

*/

// F061 切片的定义
func F061() {

	// 它与数组最直观的区别在于定义的时候不需要指定长度
	var a []string
	var b []int
	var c = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 1. 基于数组定义切片
	d := [5]int{55, 56, 57, 58, 59}
	e := d[1:4]           // 从数组 d 的第一个元素一直切到第四个元素, 然后作为切片赋值给 e, 包左不包右
	fmt.Println(e)        // [56 57 58]
	fmt.Printf("%T\n", d) // [5]int , 数组类型, 带有长度标识
	fmt.Printf("%T\n", e) // []int , 切片类型, 不带有长度标识

	// 2. 切片再次切片
	// f := e[0:len(e)]	// 从 e 的第一个元素一直到最后
	// f := e[0:]	// 从 e 的第一个元素开始一直到最后
	f := e[:]      // 从 e 中切出全部作为一个新的切片
	fmt.Println(f) // 以上三种方式是等效的

	// 3. make 函数构造切片
	// 有三个参数
	// 第一个参数: 切片的类型
	// 第二个参数: 切片的长度, 长度是切片中已有元素的个数, 如 int 类型的切片, 如果默认长度为 0, 在默认情况下会有五个 0 进行占位, 可以通过 len() 来获取长度
	// 第三个参数: 切片的容量, 容量是不扩容情况下能存入的元素的个数, 当长度超过容量的一定的比例(阈值), 则会自动扩容, 可以通过 cap() 来获取容量
	g := make([]int, 5, 10)
	fmt.Println(g)        // [0 0 0 0 0]
	fmt.Printf("%T\n", g) // []int

}

// F062 获取切片的长度
func F062() {
	a := make([]int, 5, 10)
	fmt.Println(len(a)) // 5
}

// F063 获取切片的容量
func F063() {
	a := make([]int, 5, 10)
	fmt.Println(cap(a)) // 10
}

// F064 切片的比较
func F064() {

	// 切片之间不能比较, 且不能够使用 == 操作符来判断两个切片是否含有全部相等元素
	// 切片唯一合法的比较操作是和 nil 比较
	// 一个 nil 值的切片并没有底层数组, 一个 nil 值的切片的长度和容量是 0
	// 但是我们不能说一个长度和容量都是 0 的切片一定是 nil

	var s1 []int // len(s1)=0;cap(s1)=0;s1==nil
	fmt.Println(len(s1), cap(s1), s1)
	if s1 == nil {
		fmt.Println("s1 是 nil")
	}
	s2 := []int{} // len(s2)=0;cap(s2)=0;s2!=nil
	fmt.Println(len(s2), cap(s2), s2)
	if s2 == nil {
		fmt.Println("s2 是 nil")
	}
	s3 := make([]int, 0) // len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println(len(s3), cap(s3), s3)
	if s3 == nil {
		fmt.Println("s3 是 nil")
	}

	// 所以要判断一个切片是否是空的, 要使用 len(s) == 0 来判断, 不应该使用 s == nil 来判断
}

// F065 切片的复制拷贝
func F065() {

	// 切片为引用传递

	a := make([]int, 3) // [0, 0, 0]
	b := a              // 此时 b 和 a 公用同一个底层数组
	b[0] = 100

	fmt.Println(a)
	fmt.Println(b)

}

// F066 切片的遍历
func F066() {

	// 1. 根据索引遍历
	a := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println()

	// 2. 使用 for range 遍历
	for index, value := range a {
		fmt.Println(index, value)
	}

}

// F067 切片的扩容
func F067() {

	// 切片要初始化之后才能使用
	// 尽量在切片初始化的时候就能考虑好其容量, 尽量避免切片的自动扩容

	var a []int // 此时它并没有申请内存

	// append() 函数可能会返回原来的切片, 也可能会返回一个扩容之后的切片, 所以需要用一个变量来接收它
	a = append(a, 10) // 一次往 a 中追加一个
	a = append(a, 1, 2, 3, 4, 5)

	a1 := []int{11, 12, 13, 14, 15}
	a = append(a, a1...) // 将 a1 追加到 a 中

	fmt.Println(a)

	var b []int
	for i := 0; i < 10; i++ {
		b = append(b, i)
		fmt.Printf("%v len:%d cap: %d ptr:%p\n", b, len(b), cap(b), b)
	}

	// [0] len:1 cap: 1 ptr:0xc00009e260
	// [0 1] len:2 cap: 2 ptr:0xc00009e270
	// [0 1 2] len:3 cap: 4 ptr:0xc00009a0c0
	// [0 1 2 3] len:4 cap: 4 ptr:0xc00009a0c0
	// [0 1 2 3 4] len:5 cap: 8 ptr:0xc0000ba0c0
	// [0 1 2 3 4 5] len:6 cap: 8 ptr:0xc0000ba0c0
	// [0 1 2 3 4 5 6] len:7 cap: 8 ptr:0xc0000ba0c0
	// [0 1 2 3 4 5 6 7] len:8 cap: 8 ptr:0xc0000ba0c0
	// [0 1 2 3 4 5 6 7 8] len:9 cap: 16 ptr:0xc0000ca180
	// [0 1 2 3 4 5 6 7 8 9] len:10 cap: 16 ptr:0xc0000ca180

}

// F068 切片的复制
func F068() {

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 5)
	c := b // 引用传递, 此时虽然 b 还没有值, 但是实际上是将 b 的引用指向了 c

	copy(b, a) // 第一个参数为 目标切片, 第二个参数为原切片, 将 a 切片的内容复制到 b 切片中, 两个切片分属于两个引用

	b[0] = 100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// F069 切片删除元素
func F069() {

	a := []string{"北京", "上海", "广州", "深圳"}
	a = append(a[0:2], a[3:]...) // 删除下标为 2 的那个元素, 实际上就是将原始的切片要保留的两部分组成一个新的切片
	fmt.Println(a)

}
