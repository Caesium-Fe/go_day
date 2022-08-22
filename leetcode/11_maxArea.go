package leetcode

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	var area int
	var max int
	for l < r {
		// 面积 = 以最短的高度 * 两木板的距离
		if height[l] > height[r] {
			area = height[r] * (r - l)
			r-- // 1、此处的-- 是精髓
		} else {
			area = height[l] * (r - l)
			l++ // 2、此处的++ 是精髓
		}
		if max < area {
			max = area
		}
	}
	return max
}
