package day12

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func Global_varible_assignment() {
	// 全局变量不能用 := 赋值，这种赋值属于临时变量
	// p, err := foo()
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
