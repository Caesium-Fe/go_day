package day5

import "fmt"

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func Test1() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}
