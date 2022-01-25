package main

import (
	"fmt"
	"net/http"
)

/*

	Go 语言实现一个服务端

*/

// 创建一个普通的接口函数
func sayHello(writer http.ResponseWriter, request *http.Request) {
	// 响应一个返回值
	fmt.Fprintln(writer, "Hello world")
}

func main() {

	// 定义一个 url, 并且使这个 url 调用 sayHello 函数
	http.HandleFunc("/hello", sayHello)
	// 监听本机 8080 端口并启动服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http serve failed, err: %v\n", err)
	}

}
