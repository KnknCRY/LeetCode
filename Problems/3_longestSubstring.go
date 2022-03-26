package problems

import (
	"strings"
)

// Accepted
func LengthOfLongestSubstring(s string) int {
	var resultStr string = ""
	var max int = 0
	for i := range s {
		resultStr = ""
		for j := i; j < len(s); j++ {
			// 如果現在走到的char不在resultStr內，就加進去resultStr
			if !strings.Contains(resultStr, string(s[j])) {
				resultStr += string(s[j])
			} else {
				break
			}
		}
		if len(resultStr) > max {
			max = len(resultStr)
		}
	}
	return max
}
