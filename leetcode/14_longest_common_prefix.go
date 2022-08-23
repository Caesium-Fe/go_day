package leetcode

func LongestCommonPrefix(strs []string) string {
	var common string
	var long int
	for k, str := range strs {
		if k == 0 {
			common = str
			continue
		} else {
			if len(common) > len(str) {
				long = len(str)
			} else {
				long = len(common)
			}
			if long == 0 {
				return ""
			}
			for i := 0; i < long; i++ {
				if str[i] == common[i] {
					// common = str[:i]
					if i == long-1 {
						common = common[:long]
					}
					continue
				} else {
					if i == 0 {
						return ""
					}
					common = common[:i]
					break
				}
			}
		}
	}
	return common
}
