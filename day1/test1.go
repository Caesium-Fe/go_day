package day1

import (
	"fmt"
)

func Test1() {
	var i1 int64 = 9999999
	var i2 int8 = int32(i1)
	fmt.Println(i1, i2)
}
