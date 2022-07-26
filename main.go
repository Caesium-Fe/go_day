// main
package main

import (
	"fmt"
	"go_day/day1"
	"go_day/day2"
	"go_day/day3"
	"go_day/day4"
	"go_day/day5"
	"go_day/day6"
	"go_day/day7"
	"go_day/day8"
	"go_day/learn"
	"strconv"
)

func main() {
	day7.Test4()
	day7.Test3()
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
	learn.Map_all_test()
	// learn.Msgpack_demo()
	learn.GobDemo()
	day8.Test1()
	learn.Test2_Sort()
	learn.Test1_Append()
	learn.Slice_index()
	learn.Sumr()
	learn.Find_two_index()
	learn.Scan_test()
	fmt.Println("Hello World!")
	day7.Jiujiu()
	day7.Test1()
	day7.Test2()
	day6.Test1()
	// day1.Test1()
	day5.Test1()
	day4.Test1()
	day1.Test2()

	var d [2]string
	d[0] = "cmn"
	d[1] = "fkxkvlsd"
	fmt.Println(d)
	//make只能给slice，map，channel分配内存
	// d := make([3]string)
	// d[0] = "cmn"
	// fmt.Println(d)
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
	day3.Dif_var_make()
}
