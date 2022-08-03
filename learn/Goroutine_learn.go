package learn

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var wg1 sync.WaitGroup

// 尝试使用优雅的sync包来等待子线程的结束
func hello1() {
	fmt.Println("hello")
	wg1.Done() // 告知当前goroutine完成
}

func Demo_sync_use() {
	wg1.Add(1) // 登记1个goroutine
	go hello1()
	fmt.Println("你好")
	wg1.Wait() // 阻塞等待登记的goroutine完成
}

// 多个goroutine并发执行
func hello(i int) {
	defer wg1.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}

func More_goroutine() {
	for i := 0; i < 10; i++ {
		wg1.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg1.Wait() // 等待所有登记的goroutine都结束
}

func Lianxiti1() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
