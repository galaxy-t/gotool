package main

//import (
//	log "github.com/sirupsen/logrus"
//	"os"
//)
//
///*
//
//	日志框架 logrus 的使用
//
//*/
//
//func main()  {
//
//
//	// 设置日志格式位 json 格式
//	log.SetFormatter(&log.TextFormatter{})
//
//	// 设置将日志输出到标准输出(默认的输出位 stderr, 标准错误)
//	// 日志消息输出可以是任意的 io.Writer 类型
//	log.SetOutput(os.Stdout)
//
//	// 打印日志所产生的位置
//	// 效果: func=main.main file="D:/dev/go/gotool/example/4frame/02logrus.go:32"
//	log.SetReportCaller(true)
//
//	// 设置日志输出级别
//	log.SetLevel(log.InfoLevel)
//
//
//	log.WithFields(log.Fields{
//		"animal": "walrus",
//	}).Info("A walrus appears")
//
//	log.Info("skldjfkldsf")
//
//	log.Debug("debug")
//
//	log.Error("abcdefg", "hahahaha")
//
//
//
//}
