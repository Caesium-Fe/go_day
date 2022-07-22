package learn

import (
	"fmt"
	"sort"
)

func Slice_index() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	fmt.Printf("%v", s[:4])
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

// 切片遍历
func Bianli_slice() {
	s := []int{1, 3, 5}

	//方式同数组一致
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}

// append()方法为切片添加元素
func Append_slice() {
	var s []int
	s = append(s, 1)       // [1]
	s = append(s, 2, 3, 4) // [1 2 3 4]
	s2 := []int{5, 6, 7}
	s = append(s, s2...) // [1 2 3 4 5 6 7]
}

//练习题1
func Test1_Append() {
	var a = make([]string, 5, 10)
	if a[0] == "0" {
		fmt.Println()
	}
	fmt.Printf("a:%v len(a):%v cap(a):%v\n", a, len(a), cap(a))
	for i := 0; i < 2; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Printf("a:%v len(a):%v cap(a):%v\n", a, len(a), cap(a))
}

//请使用内置的sort包对
//数组var a = [...]int{3, 7, 8, 9, 1}进行排序
func Test2_Sort() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	// sort.IntSlice(a[:])
	if a[0] == 0 {
		fmt.Println()
	}
	fmt.Printf("a:%v len(a):%v cap(a):%v\n", a, len(a), cap(a))
	/*for i := 0; i < 2; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Printf("a:%v len(a):%v cap(a):%v\n", a, len(a), cap(a))
	*/
}
