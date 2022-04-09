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
			fmt.Println("面積沒比較大")
		}
		if side == "L" {
			fmt.Println("左邊比較短，左邊往中間一步") // i++
		} else {
			fmt.Println("右邊比較短，右邊往中間一步") // j--，i--(因為i等等會被++)
			j--
			i--
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
