package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	// 纯数学方法
	var rev int
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev

	// 转成字符串方法
	s := string(x)
	r := []rune(s)
	if string(r[0]) == '-' {
		for i, j := 1, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
	} else {
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
	}
	s = string(r)
	j, err := strconv.Atoi(s)
	if err {
		fmt.Println(err)
	}
	if j < math.MinInt32/10 || j > math.MaxInt32/10 {
		return j
	} else {
		return 0
	}
}
