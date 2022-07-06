// main
package main

import (
	"fmt"
	"go_day/day1"
	"go_day/day2"
	"go_day/learn"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
	// day1.Test1()
	day1.Test2()

	//print(len("abcdefghijk"))
	//learn.Recursion()
	learn.Mutex1()
	//str := "hello\n12"
	str := `hello\n12` + "asd"
	str1 := []byte(str)
	str1[0] = 'x'
	str = string(str1)
	fmt.Println(str)
	varsa := 12.6
	varsa2 := int(varsa)
	aa := fmt.Sprintf("%f", varsa)
	a, c := strconv.ParseFloat(aa, 32)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(varsa2)
	learn.Abspath()
	// learn.Day2()
	// day2.Test1()
	// day2.Test2()
	day2.Test3()
}
