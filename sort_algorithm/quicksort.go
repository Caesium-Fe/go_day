package sort_algorithm

import (
	"fmt"
	"math/rand"
)

/*
快排思想：
	找出中位数，然后分成前后两份（前小后大，反之亦可）
	前后两份也按这个思想找
	退出条件为只有一个元素
*/

func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	base := data[0]
	l, r := 0, len(data)-1
	for i := 1; i <= r; {
		if data[i] > base {
			data[i], data[r] = data[r], data[i]
			r--
		} else {
			data[i], data[l] = data[l], data[i]
			l++
			i++
		}
	}
	quickSort(data[:l])
	quickSort(data[l+1:])
}

func QuickSort() {
	s := make([]int, 0, 16)
	for i := 0; i < 16; i++ {
		s = append(s, rand.Intn(100))
	}
	fmt.Println(s)
	quickSort(s)
	fmt.Println(s)
}
