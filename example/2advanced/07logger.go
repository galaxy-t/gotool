package _advanced

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 该函数要测试的话需要修改为 init
func init1() {

	// 设置日志前缀
	// 以最后一次设置为准, 设置后全局通用
	log.SetPrefix("日志前缀 ")

	// 设置日志输出的各种属性
	// log.LstdFlags:  设置打印本地时区的日期和时间, 默认选择
	// log.Lmsgprefix: 设置将日志前缀打印到日志消息的前面
	// log.Llongfile:  设置打印完整的文件名和行号. D:/dev/go/gotool/glog/test.go:11
	// log.Lshortfile:  设置打印文件名和行号. test.go:11
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 设置日志输出方式位 log 文件
	logFile, err := os.OpenFile("glog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open log file failed, err: ", err)
		return
	}
	defer logFile.Close()

	// 将文件写入和控制台输出进行组合, 以做到日志不仅写入文件, 同时会写入到控制台
	mw := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(mw)

	// log.Print()
	// log.Printf("这是一条日志%s","哈哈哈")
	log.Println("这是一条日志")

	// 打印日志后会结束当前线程, 但不会退出程序(在web开发中)
	// log.Panic()
	// log.Panicf()
	log.Panicln("这是一条会引发 panic 的日志")

	// 打印日志后退出程序, 执行 os.Exit(1)
	// log.Fatal()
	// log.Fatalf()
	log.Fatalln("这是一条会引发 fatal 的日志")
}
