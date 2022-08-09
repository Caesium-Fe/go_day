package sort_algorithm

import (
	"fmt"
	"math/rand"
)

/*
归排思想：
	每次取一半长度的元素排序，将排序后的数组返回
	将左右两半有序的数组，再进行排序
*/

func mergeSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	num := length / 2
	left := mergeSort(data[:num])
	right := mergeSort(data[num:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

func MergeSort() {
	s := make([]int, 0, 16)
	for i := 0; i < 16; i++ {
		s = append(s, rand.Intn(100))
	}
	fmt.Println(s)
	s = mergeSort(s)
	fmt.Println(s)
}
