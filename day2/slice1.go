package day2

import (
	"fmt"
)

var months = [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

func Test1() {
	// months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	// fmt.Printf("%T", months)
	// var months = get_month()
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2) //["April" "May" "June"]
	fmt.Println(summer)
}

func Test2() {
	//使用一个值接收range, 则得到的是切片的下标
	for i := range months {
		fmt.Println(i) //返回下标 0 1 ... 12
	}
	//使用两个值接收range，则得到的是下标和对应的值
	for i, v := range months {
		fmt.Println(i, v) //返回下标0 1 ... 12 和 值 "" "January" ... "December"
	}
}

func Test3() {
	var a []int
	a = append(a, 1)                 // 追加一个元素
	a = append(a, 1, 2, 3)           // 追加多个元素
	a = append(a, []int{1, 2, 3}...) // 追加一个切片，注意追加切片时后面要加...
	fmt.Println(a)
	a = append([]int{0}, a...)       // 在开头添加一个元素
	a = append([]int{1, 2, 3}, a...) // 在开头添加一个切片
	fmt.Println(a)
	a = append(a[:1], append([]int{9}, a[1:]...)...)       //在第i个位置上插入x
	a = append(a[:1], append([]int{6, 8, 7}, a[1:]...)...) //在第i个位置上插入切片
	fmt.Println(a)
}
