package day3

import (
	"fmt"
)

func Dif_var_make() {
	// 以下只是声明m变量，但并没有分配内存
	//顾不可以直接进行赋值操作
	// var m map[string]int
	// m["a"] = 1
	//由于没有 "b" 元素，v会返回值类型对应的零值
	// if v := m["b"]; v != nil {
	// 	fmt.Println(v)
	// }

	m := make(map[string]int)
	m["a"] = 1
	//由于没有 "b" 元素，ok就为false
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}
}
