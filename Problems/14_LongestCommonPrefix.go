package problems

import "fmt"

func LongestCommonPrefix(strs []string) string {
	shortestIdx := 0
	shortestLength := len(strs[0])
	var result string

	// 先找到最短的字串
	for i, item := range strs {
		if len(item) < shortestLength {
			shortestLength = len(item)
			shortestIdx = i
		}
	}

	fmt.Println("idx", shortestIdx, "length", shortestLength)

	// 跑過最短的字串
	var accepted bool = true
	for i := 0; i < shortestLength; i++ {
		fmt.Println("最小字串的第", i, "個位置:", string(strs[shortestIdx][i]))

		// 跑過輸入的array，並比較array內所有字串的第i個位置
		for j, str := range strs {
			if j == shortestIdx {
				continue
			}
			// 如果子串的i != 最短字串的i，代表i不符合最短共同字串
			// i要繼續往下走，並把result重置
			fmt.Println("開始跟", str, "比對")
			if str[i] != strs[shortestIdx][i] {
				fmt.Println("失敗")
				accepted = false
				break
			}
			fmt.Println("接受")
		}

		// 如果i跟所有array內的字串比對，都有對到
		// 那麼就可以繼續往下比，並把i存進result
		if accepted {
			result += string(strs[shortestIdx][i])
		}
	}

	return result
}
