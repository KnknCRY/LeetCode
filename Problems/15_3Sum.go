package problems

import "fmt"

func ThreeSum(nums []int) [][]int {
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					fmt.Println("找到", nums[i], nums[j], nums[k])
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	if len(result) == 0 {
		return result
	}

	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			y := map[int]int{}
			for _, value := range result[i] {
				y[value]++
			}
			for _, value := range result[j] {
				y[value]++
			}

			fmt.Println("比對idx", i, j)
			fmt.Println("比對val", result[i], result[j])
			fmt.Println("y:", y)

			var same bool = true
			for _, value := range y {
				if value%2 != 0 { // 如果不是偶數，代表value出現幾數次，那就一定不重複
					same = false
					break
				}
			}
			if same {
				fmt.Println("找到一樣，殺掉", result[j])
				result = append(result[:j], result[j+1:]...)
				fmt.Println("殺完後", result)
				j--
			}
		}
	}

	return result
}
