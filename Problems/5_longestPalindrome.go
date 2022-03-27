package problems

// Accepted
func LongestPalindrome(s string) string {
	palindrome := make(map[string]interface{})
	result := ""

	for i := range s {
		for j := i + 1; j <= len(s); j++ {
			// 先切出此次的sub
			sub := s[i:j]

			if _, exist := palindrome[sub]; !exist {
				if isPalindrome(sub) {
					palindrome[sub] = nil

					// 取最長的sub
					if len(result) < len(sub) {
						result = sub
					}
				}
			}
		}
	}
	return result
}

func isPalindrome(input string) bool {
	j := len(input) - 1

	for i := 0; i < len(input); i++ {
		if input[i] != input[j] {
			return false
		}
		j--
	}
	return true
}

// func longestPalindrome(s string) string {
// 	lenS := len(s)

// 	maxStr := s[:1]
// 	for i := 0; i < lenS; i++ {
// 		var str string
// 		for l, r := i, i; l >= 0 && r < lenS && s[l] == s[r]; l, r = l-1, r+1 {
//             str = s[l:r+1]
// 		}
//         if len(maxStr) < len(str) {
// 			maxStr = str
// 		}

//         for l, r := i, i+1; l >= 0 && r < lenS && s[l] == s[r]; l, r = l-1, r+1 {
//             str = s[l:r+1]
// 		}
// 		if len(maxStr) < len(str) {
// 			maxStr = str
// 		}
// 	}
// 	return maxStr
// }
