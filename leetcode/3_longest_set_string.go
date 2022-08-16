package leetcode

import "strings"

func lengthOfLongestSubstring(s string) int {
	// a := make([]string, 128, 128)
	d, max := 0, 0
	t := ""
	for _, i := range s {
		// 查找字串中的index
		if j := strings.Index(t, string(i)); j != -1 {
			t = t[j+1:]
		}
		t += string(i)
		d = len(t)
		if d > max {
			max = d
		}
	}
	return max
}

//优化前后代码执行时间下降，但内存消耗任然居高，下一次优化内存分配
