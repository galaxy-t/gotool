package example

import (
	"fmt"
	"math"
	"strings"
)

/*

	基础数据类型

	整型

	uint8 	无符号 8 位整型 (0 到 255)
    uint16 	无符号 16 位整型 (0 到 65535)
    uint32 	无符号 32 位整型 (0 到 4294967295)
    uint64 	无符号 64 位整型 (0 到 18446744073709551615)
    int8 	有符号 8 位整型 (-128 到 127)
    int16 	有符号 16 位整型 (-32768 到 32767)
    int32 	有符号 32 位整型 (-2147483648 到 2147483647)
    int64 	有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

	uint	32位操作系统上就是 uint32, 64位操作系统上就是 uint64
	int		32位操作系统上就是 int32, 64位操作系统上就是 int64
	uintptr	无符号整型, 用于存放一个指针

*/
func F021()  {
	// 十进制打印为二进制
	n := 10
	fmt.Printf("%b\n", n)	// 二进制打印
	fmt.Printf("%d\n", n)	// 十进制打印

	// 八进制
	m := 075
	fmt.Printf("%o\n", m)	// 八进制打印
	fmt.Printf("%d\n", m)	// 十进制打印

	// 十六进制
	f := 0xff
	fmt.Println(f)					// 十进制打印
	fmt.Printf("%x\n",f)		// 十六进制打印

	// uint8
	var age uint8
	fmt.Println(age)

	// 浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	// 布尔值
	// 1. 布尔类型变量默认值为 false
	// 2. Go 语言中不允许将整型强制转换为布尔型
	// 3. 布尔型无法参与数值运算, 也无法与其他类型进行转换
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	// 字符串
	// 使用 UTF-8 编码
	// 字符串的值为 双引号 "" 中的内容
	s1 := "hello world"
	s2 := "你好"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("c:\\code\\go.exe")		// 打印 windows 平台下的一个路径 c:\code\go.exe
	// 多行字符串
	s3 := `多行字符串
		两个反引号之间的内容
		会原样输出
		\t
		\n
	`
	fmt.Println(s3)
	// 求字符串的长度
	s4 := "hello"
	fmt.Println(len(s4))
	s5:= "hello 你好"
	fmt.Println(len(s5))
	// 拼接字符串
	fmt.Println(s4 + s5)
	s6 := fmt.Sprintf("%s - %s", s4, s5)
	fmt.Println(s6)
	// 字符串的分割
	s7 := "how do you do"
	fmt.Println(strings.Split(s7, " "))
	fmt.Printf("%T\n", strings.Split(s7, " "))	// 打印字符串切割后返回的类型
	// 判断是否包含
	fmt.Println(strings.Contains(s7, "do"))
	// 判断前缀
	fmt.Println(strings.HasPrefix(s7, "how"))
	// 判断后缀
	fmt.Println(strings.HasSuffix(s7, "how"))
	// 判断子串出现的位置
	fmt.Println(strings.Index(s7, "do"))
	// 判断子串最后出现的位置
	fmt.Println(strings.LastIndex(s7, "do"))
	// join, 拼接数组成字符串
	s8 := []string{"how", "do", "you", "do"}
	fmt.Println(s8)
	s9 := strings.Join(s8, "-")
	fmt.Println(s9)


	// byte 和 rune 类型
	// 组成每个字符串的元素叫做 "字符", 可以通过遍历或者单个获取字符串元素获得字符
	// 字符用单引号 '' 包裹起来
	// byte uint8 的别名, 只能标识常见的 ASCII 码
	// rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T   c2:%T\n", c1, c2)		// 打印 c1 和 c2 的类型
	// 打印字符串中的每一个字符
	s10 := "hello你好"
	for i:=0;i<len(s10);i++ {
		fmt.Printf("%c\n",s10[i])
	}
	fmt.Println()
	for _,r :=range s10{
		fmt.Printf("%c\n", r)
	}


}