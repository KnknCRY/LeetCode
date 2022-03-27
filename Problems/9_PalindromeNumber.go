package problems

import (
	"strconv"
)

// Accepted
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		s := strconv.Itoa(x)
		return isPalindrome(s)
	}
}

// func isPalindrome(x int) bool {
//   s := strconv.Itoa(x)
//   r := []rune(s)
//   for i, j := 0, len(r) - 1; i < j; i, j = i + 1, j - 1 {
//       if r[i] != r[j] {return false}
//   }
//   return true
// }
