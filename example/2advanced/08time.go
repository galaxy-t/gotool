package _advanced

import (
	"fmt"
	"time"
)

func F081() {

	// 获取系统当前时间, 默认时区为本地时区
	now := time.Now()

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 时
	minute := now.Minute() // 分
	second := now.Second() // 秒

	fmt.Println(now)
	fmt.Println(year, month, day, hour, minute, second)

	timestamp1 := now.Unix()     // 毫秒时间戳, 自 1970年1月1日至当前的总毫秒数
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Println(timestamp1, timestamp2)

	time1 := time.Unix(timestamp1, 0)	// 将时间戳转换为时间格式
	fmt.Println(time1)

	// 是 time 包定义的一个类型, 用于表示两个时间点之间的时间, 以纳秒为单位
	// 最长时间段大约 290 年
	var a time.Duration = time.Minute * 10	// 表示十分钟
	fmt.Println(a)

	// 为 time1 增加三个小时, 如果要减少则可以传入负数
	time1.Add(time.Hour * 3)

	// 求两个时间之间的差
	// now.Sub(time1)

	// 定时器, 本质上是一个通道
	/*ticker := time.Tick(time.Second * 10)
	for i := range ticker {
		fmt.Println(i)	// i 为发生的当前时间
	}*/

	// 日期格式化成字符串, 使用 2006-01-02 15-03-04 代替 yyyy-MM-dd HH:mm:ss
	ftime := now.Format("2006-01-02 15-03-04")
	fmt.Println(ftime)

	// 根据字符串格式化时间, 默认时区为 UTC
	time2, err  := time.Parse("2006-01-02 15:04:05","2021-11-03 23:06:01")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time2)

	time3,_ := time.ParseInLocation("2006-01-02 15:04:05","2021-11-03 23:06:01",time.Local)
	fmt.Println(time3)


}
