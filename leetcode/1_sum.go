package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	// 有求值数组  有结果值   找两个数坐标
	var a []int
	for i, v1 := range nums {
		// 一个算一遍 不重复
		for j, v2 := range nums[i+1:] {
			if sum := v1 + v2; sum == target {
				// fmt.Println(i, i+j)
				a = append(a, i)
				a = append(a, i+j+1)
				return a
			}
		}
	}
	return nil
}

func TwoSum() {
	nums := [3]int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
