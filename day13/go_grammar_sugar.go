package day13

import "fmt"

// 下面这段代码输出什么？

// Go 提供的语法糖...
// 可以将 slice 传进可变函数
// 不会创建新的切片
func change(s ...int) {
	s = append(s, 3)
}

func Sugar_grammar() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}
