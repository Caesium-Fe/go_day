package day11

import "fmt"

const (
	a = iota
	b
)
const (
	// iota在遇到const关键字会被重置为0
	// const里的技巧：当第一个常量被赋值后
	// 后续常量如果没有赋值，那么值同上一个一致
	name = "name"
	c    = iota
	d
)

func Iota_learn() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

//当类型定义了String方法
// 那么在使用fmt输出时就会自动调用String
type Direction int

const (
	//  iota只需要给一个const赋值
	// 那么剩下的就会依次赋值
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

// 有个例外 就是当类型为指针类型时
// 并不会自动调用，得手动
// func (d *Direction) String() string {
// 	return [...]string{"North", "East", "South", "West"}[*d]
// }

func Iota_test1() {
	fmt.Println(South)
}
