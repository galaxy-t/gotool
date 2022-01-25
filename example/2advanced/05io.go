package _advanced

import (
	"fmt"
	"io/ioutil"
)

// F051 直接读取文件中的全部内容
func F051() {
	data, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		fmt.Println("读取文件异常", err)
	}
	fmt.Println(string(data))
}
