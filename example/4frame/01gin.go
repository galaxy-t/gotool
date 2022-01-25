package main

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"time"
//)
//
///*
//
//	Gin 框架的使用
//
//*/
//
//type HelloVo struct {
//	Message string `json:"message"`
//}
//
//type User struct {
//	Name     string `json:"name" form:"name"` // tag: form 用于表示从 request 中的请求参数进行反射, url,form , 如果 url和form 同时存在重名参数, 则优先 form
//	Age      int    `json:"age" form:"age"`   // 如果请求是使用 json 格式进行传参, 那么 只需要写 这两个 tag 中的一个即可
//	Sex      string `json:"sex" form:"sex"`
//	Birthday string `json:"birthday" form:"birthday"`
//}
//
//func sayHello(context *gin.Context) {
//
//	// 获取 url 中传递的参数
//	// URL 获取中文时会自动转码, 不会被转译
//	name := context.Query("name")
//	fmt.Println(name)
//	age := context.Query("age")
//	fmt.Println(age)
//	sex := context.DefaultQuery("sex", "男") // 从 url 中获取参数, 若 key 不存在则返回设置的默认值
//	fmt.Println(sex)
//	birthday, ok := context.GetQuery("birthday")
//	if ok {
//		fmt.Println(birthday)
//	} else {
//		fmt.Println("请选择生日")
//	}
//	// 获取 POST 请求时 Form 提交的参数
//	name1 := context.PostForm("name")
//	fmt.Println("form 表单提交:", name1)
//	age1 := context.PostForm("age")
//	fmt.Println("form 表单提交:", age1)
//
//	// 根据请求中的 contentType 类型进行绑定参数
//	// 绑定前端传递的参数到结构体中, 结构体的字段必须使用 tag(form) 进行标注
//	// 可以用于绑定 get,post,json 等, 如果 get 传参和其它方式同时存在, 则会忽略 get 传参, 即使参数都在 url 中也不会被选择
//	var user User
//	err := context.ShouldBind(&user)
//	if err != nil {
//		context.String(http.StatusInternalServerError, err.Error())
//	}
//	fmt.Printf("%#v\n", user)
//
//	// 从请求中读取文件
//	// f,err := context.FormFile()
//	// context.SaveUploadedFile(f, "path/f.Filename")
//
//	// 响应
//	// 1. 响应一个 json 串
//	context.JSON(http.StatusOK, gin.H{
//		"hello": "hello",
//	})
//
//	// 2. 使用 map 组装一个 json
//	m := map[string]interface{}{
//		"hello": "hello111",
//	}
//	context.JSON(http.StatusOK, m)
//
//	// 3. 使用结构体封装返回一个 json
//	context.JSON(http.StatusOK, HelloVo{Message: "hello222"})
//
//	// 4. 响应一个字符串
//	context.String(http.StatusOK, "hello333%s", "4444")
//
//}
//
//// 获取路径通配参数
//func uriParams(context *gin.Context) {
//	pageSize := context.Param("pageSize")
//	pageIndex := context.Param("pageIndex")
//	fmt.Println(pageSize, pageIndex)
//
//	// 重定向, 内部,外部重定向均支持
//	// context.Redirect(http.StatusMovedPermanently, "url")
//	context.JSON(http.StatusOK, gin.H{
//		"message": "ok",
//	})
//}
//
//// 测试函数
//func test(context *gin.Context) {
//	fmt.Println("test========")
//	context.JSON(http.StatusOK, gin.H{
//		"message": "ok",
//	})
//}
//
//// 接口耗时管理处理程序
//func timeConsumingHandler(context *gin.Context) {
//	fmt.Println("timeConsumingHandler    start")
//	start := time.Now()	// 记录开始时间
//	context.Next() // 调用后续的处理程序
//	// context.Abort()	// 阻止调用后续的处理函数
//	cost := time.Since(start)	// 计算执行时间
//	fmt.Println("timeConsumingHandler    end", cost)
//
//}
//
//func main() {
//
//	// 注意, 如果需要在一个中间件中使用 goroutine, 需要使用 gin.Context 的拷贝, 不能使用原始的上下文
//
//	// 创建一个默认的路由引擎
//	// 默认使用 Logger和Recovery 中间件
//	// Logger 中间件将日志写入 gin.DefaultWriter, 即使配置了 GIN_MODE=relaese
//	// Recovery 中间件会 recover 任何 panic, 如果有 panic 的话, 会写入 500 响应码
//	r := gin.Default()
//
//	// 新建一个没有任何默认中间件的路由
//	// r := gin.New()
//
//	// 统一注册全局处理程序
//	r.Use(timeConsumingHandler)
//
//	// 普通路由
//	// 指定用户使用 GET 请求访问 /hello 时, 执行 sayHello 这个函数
//	r.POST("/hello", sayHello)
//	r.GET("/article/:pageSize/:pageIndex", uriParams)
//	r.GET("/test", timeConsumingHandler, test) // 多个执行程序按照传参顺序正序执行
//
//	// 使用该函数用来处理全部的请求类型(GET,POST等)
//	r.Any("/index", func(context *gin.Context) {
//		switch context.Request.Method {
//		case http.MethodGet:
//			fmt.Println("get")
//		case http.MethodPost:
//			fmt.Println("post")
//		default:
//			fmt.Println("default")
//		}
//	})
//
//	// 路由组
//	userGroup := r.Group("/user")
//	{ // 一个普通代码块, 仅用于区分, 无实际意义
//		userGroup.GET("", func(context *gin.Context) {
//			fmt.Println("userlksdjflsdjlk")
//		})
//		// 嵌套路由
//		xxGroup := userGroup.Group("xxx")
//		{
//			xxGroup.GET("", func(context *gin.Context) {
//			})
//		}
//	}
//
//	// 设置 404
//	r.NoRoute(func(context *gin.Context) {
//		context.JSON(http.StatusNotFound, gin.H{
//			"message": "未找到页面",
//		})
//	})
//
//	// 启动服务, 并指定端口为 9090(端口默认为8080)
//	r.Run(":9090")
//
//}
