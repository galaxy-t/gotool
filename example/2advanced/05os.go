package _advanced

import (
	"fmt"
	"io/ioutil"
	"os"
)

// F051 直接读取文件中的全部内容
func F051() {
	data, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		fmt.Println("读取文件异常", err)
	}
	fmt.Println(string(data))
}

// F052 OS 包中的常用函数
func F052() {


	// 查询内核提供的主机名
	name, err := os.Hostname()
	fmt.Println(name)
	fmt.Println(err)

	// Getpagesize返回底层的系统内存页的尺寸
	pagesize := os.Getpagesize()
	fmt.Println(pagesize)

	// 返回所有的环境变量
	envArr := os.Environ()
	for index, value := range envArr {
		fmt.Println(index, value)
	}

	// 检索并返回名为 key 的环境变量的值。如果不存在该环境变量则会返回空字符串。
	evn := os.Getenv("TEMP")
	fmt.Println(evn)

	// 设置名为 key 的环境变量，如果出错会返回该错误。
	// os.Setenv()

	// Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。
	// os.Exit(0)

	// Clearenv删除所有环境变量。
	// os.Clearenv()

	// Getuid 函数可以返回调用者的用户 ID。
	uid := os.Getuid()
	fmt.Println(uid)

	// Geteuid返回调用者的有效用户ID。
	euid := os.Geteuid()
	fmt.Println(euid)

	// Getgid 函数可以返回调用者的组 ID。
	gid := os.Getgid()
	fmt.Println(gid)

	// Getegid返回调用者的有效组ID。
	egid := os.Getegid()
	fmt.Println(egid)

	// Getpid 函数可以返回调用者所在进程的进程 ID。
	pid := os.Getpid()
	fmt.Println(pid)

	// Getppid返回调用者所在进程的父进程的进程ID。
	ppid := os.Getppid()
	fmt.Println(ppid)

	// Getwd 函数可以返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd 会返回其中一个。
	dirPath, err := os.Getwd()
	fmt.Println(dirPath, err)


	// Mkdir 函数可以使用指定的权限和名称创建一个目录。如果出错，会返回 *PathError 底层类型的错误。
	// err := os.Mkdir("c://abcd",0777)

	// MkdirAll 函数可以使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回 nil，否则返回错误。权限位 perm 会应用在每一个被该函数创建的目录上。如果 path 指定了一个已经存在的目录，MkdirAll 不做任何操作并返回 nil。
	// err := os.MkdirAll("c://a/b/c/d", 0777)

	// Remove 函数会删除 name 指定的文件或目录。如果出错，会返回 *PathError 底层类型的错误。
	// err := os.Remove("c://abcd")

	// RemoveAll 函数跟 Remove 用法一样，区别是会递归的删除所有子目录和文件。
	// err := os.RemoveAll("c://a/b/c/d")

}
