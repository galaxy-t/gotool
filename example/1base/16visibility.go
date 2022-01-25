package _base

import (
	"encoding/json"
	"fmt"
)

/*

	结构体字段的可见性
	结构体中字段大写开头表示可公开访问, 小写表示私有(仅在定义当前结构体的包中可访问)
	如果一个 Go 语言中定义的标识符是首字母大写的, 那么就是对外可见的(对所有包(package)可见)
	如果一个结构体中的字段名首字母是大写的, 那么该字符就是对外可见的

	结构体与 JSON 序列化
	结构体标签(Tag)
	Tag 是结构体的元信息, 可以在运行的时候通过反射的机制读取出来, Tag 在结构体字段的后方定义, 由一对反引号包裹起来, 如下:
	`key1:"value1" key2:"value2"`
	ID   int	`json:"id"`
	Title    string    `json:"title,omitempty"`

*/

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	age  int    // 如果结构体中一个属性首字母是小写的, 那么在转换 json 的时候是无法获取到的
}

// NewStudent Student 的构造函数
func NewStudent(id int, name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}

type Class struct {
	Title    string    `json:"title,omitempty"`
	Students []Student `json:"students,omitempty"`
}

func F161() {

	// 创建一个班级变量
	c1 := Class{
		Title:    "1班",
		Students: make([]Student, 0, 20),
	}
	// 往班级 c1 中添加学生
	for i := 0; i < 10; i++ {
		// 创建 10 个学生并添加到班级
		c1.Students = append(c1.Students, NewStudent(i, fmt.Sprintf("stu%02d", i)))
	}
	fmt.Printf("%#v\n", c1)

	// JSON 序列化: Go 语言中的数据 -> JSON 格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("JSON 转化失败: ", err)
		return
	}
	fmt.Printf("%T\n", data) // []uint8
	fmt.Printf("%s\n", data)

	// JSON 反序列化: JSON 格式的字符串 -> Go 语言中的数据
	var c2 Class
	err = json.Unmarshal(data, &c2)
	if err != nil {
		fmt.Println("JSON 反序列化失败: ", err)
		return
	}
	fmt.Printf("%#v\n", c2)

}
