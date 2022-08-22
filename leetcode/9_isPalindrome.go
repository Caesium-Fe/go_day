package leetcode

import "math"

func isPalindrome(x int) bool {
	com := x
	rev := 0
	for x > 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return false
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	if rev == com {
		return true
	} else {
		return false
	}
}
