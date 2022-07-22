package day8

import (
	"fmt"
)

func Test1() {
	const a = 1
	b := float64(1.12) + a
	fmt.Printf("%T\n", b)
	fmt.Println(b)
	c := int64(1) + a
	fmt.Printf("%T\n", c)
	fmt.Println(c)
	d := byte(1) + a
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", d)
	fmt.Println(d)
}
