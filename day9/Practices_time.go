package day9

import (
	"fmt"
	"time"
)

// 获取当前时间
// 格式化输出为2017/06/19 20:30:05格式
func Format_time() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}

// 编写程序统计一段代码的执行耗时时间
// 单位精确到微秒
func get_now_time() int64 {
	return time.Now().UnixMilli()
}

func Use_time() {
	ago := get_now_time()
	time.Sleep(time.Second * 3)
	now := get_now_time()
	duration := now - ago
	// duration := now.Sub(ago)
	fmt.Println(duration)
}
