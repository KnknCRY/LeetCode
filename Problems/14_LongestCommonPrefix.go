package problems

import "fmt"

func LongestCommonPrefix(strs []string) string {
	shortestIdx := 0
	shortestLength := len(strs[0])
	var result rune

	// 先找到最短的字串
	for i, item := range strs {
		if len(item) < shortestLength {
			shortestLength = len(item)
			shortestIdx = i
		}
	}

	fmt.Println("idx", shortestIdx, "length", shortestLength)

	// 用最短的字串跟其他每個比
	// 先跑過strs裡面的每一個字串
	for i := 0; i < len(strs); i++ {
		// 不用跟最短的那個比
		if i != shortestIdx {
			// 再來把每一個字串 跟 最短的字串比較，一個一個字元比
			for j, item := range strs[shortestIdx] {
				if strs[i][j] == strs[shortestIdx][j] {
					continue
				}
			}
		}
	}

	return ""
}
