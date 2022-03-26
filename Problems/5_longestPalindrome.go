package problems

// 幹！Time Limit Exceeded
func LongestPalindrome(s string) string {
	palindrome := make(map[string]string)
	result := ""

	for i := range s {
		for j := i + 1; j <= len(s); j++ {
			//如果現在切出來的substring不在map內
			if _, exist := palindrome[s[i:j]]; !exist {
				temp := s[i:j]
				tempReverse := reverse(temp)

				// 如果把現在切出來的substring是回文
				if temp == tempReverse {
					//就把他與他的回文存進去
					palindrome[temp] = tempReverse

					if len(result) < len(temp) {
						result = temp
					}
				}
			}
		}
	}
	return result
}

func reverse(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	return string(rune)
}
