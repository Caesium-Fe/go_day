package leetcode

import "fmt"

// 超出时间限制  原因是没有break跳出死循环
func LongestPalindrome(s string) string {
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	// t += "#"
	// return t
	// 存储最长回文字符串
	ans := ""
	for i := 0; i < len(t); i++ {
		j := 1
		now := ""
		for {
			// if s[i] == "#"{
			//     break
			// }
			right := j + i
			left := i - j
			if (left >= 0) && (right < len(t)) {
				// if (string(t[left]) == "#") || (string(t[right]) == "#") {
				// 	j++
				// 	continue
				// }
				if t[left] == t[right] {
					j++
					now = t[left : right+1]
					fmt.Println(now)
					if len(now) > len(ans) {
						ans = now
					}
				} else {
					break
				}
			} else {
				break
			}
		}
	}
	now := ""
	for i := 0; i < len(ans); i++ {
		if ans[i] != '#' {
			now += string(ans[i])
		}
	}
	return now
}
