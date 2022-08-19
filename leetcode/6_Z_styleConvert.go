package leetcode

func Convert(s string, numRows int) string {
	if len(s) <= numRows {
		return s
	} else if numRows == 1 {
		return s
	}
	ans := ""
	//first row and last row is same interval
	// interval := numRows * 2 - 2
	// i is axis of s
	for i := 0; i < numRows; i++ {
		row := string(s[i])
		// count num n,m
		m, n := 1, i
		// flag num j
		j := 0

		// each row interval
		interval := numRows*2 - 2
		interval_left := i * 2
		interval_right := interval - interval_left
		for {
			if i == 0 || i == numRows-1 {
				// interval := numRows*2 - 2
				end := i + m*interval
				m++
				if end < len(s) {
					row += string(s[end])
				} else {
					break
				}
			} else {
				// interval := numRows*2 - 2
				// interval_right := i * 2
				// interval_left := interval - interval_right
				// n := i + interval_left
				// end_r := i + n*interval_right
				// n++
				// if end_l >= len(s) || end_r >= len(s) {
				// 	break
				// } else {
				// 	if j == 0 {
				// 		row += string(s[end_r])
				// 		j = 1
				// 	} else {
				// 		row += string(s[end_l])
				// 		j = 0
				// 	}
				// }
				if j == 0 {
					// right
					n += interval_right
					if n < len(s) {
						row += string(s[n])
						j = 1
					} else {
						break
					}
				} else {
					//left
					n += interval_left
					if n < len(s) {
						row += string(s[n])
						j = 0
					} else {
						break
					}
				}
			}
		}
		ans += row
	}
	return ans
}
