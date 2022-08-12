package day13

import (
	"fmt"
)

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	// 这里defer f并没有初始化 所以f == nil 调用不会执行后面初始化的函数体
	// 移到初始化代码后执行 结果就为9
	// defer f()
	f = func() {
		r += 2
	}
	defer f()
	return n + 1
}

func Defer_func() {
	fmt.Println(f(3))
}
