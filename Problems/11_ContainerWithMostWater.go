package problems

import "fmt"

func MaxArea(height []int) int {
	result := 0
	j := len(height) - 1

	for i := 0; i < len(height) && i != j; i++ {
		fmt.Println(i, j, "兩邊", height[i], height[j])
		shorter, side := getShorterOne(height[i], height[j])
		area := shorter * (j - i)

		fmt.Println("較短的一邊", side, "數值為", shorter, "面積為", area)

		if area > result {
			fmt.Println("更大面積出現", area)
			result = area
		} else {
			fmt.Println("沒比較大")
			// 如果左邊比較短
			if side == "L" {
				fmt.Println("左邊比較短")
				// 左邊的下一個又比現在的左邊短，則跳過他
				if height[i+1] < shorter {
					fmt.Println("下一個左邊又更短，跳過")
					i++
				}
			} else {
				fmt.Println("右邊比較短")
				// 如果右邊比較短
				// 右邊的下一個比現在的右邊長，則j--
				if height[j-1] > shorter {
					fmt.Println("下一個右邊比較長，j--")
					i-- // 此時i應該不要動，所以先--，因為等等會被++
					j--
				} else {
					j--
				}
			}
		}
	}
	return result
}

func getShorterOne(a, b int) (int, string) {
	if a <= b {
		return a, "L"
	} else {
		return b, "R"
	}
}
