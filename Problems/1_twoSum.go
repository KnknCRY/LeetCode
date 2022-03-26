package problems

import "fmt"

// Accepted
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			fmt.Println(idx, ok)
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{0, 0}
}

// func twoSum(nums []int, target int) (ans []int) {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			sum := nums[i] + nums[j]
// 			if sum == target {
// 				ans = append(ans, i, j)
// 				return
// 			}
// 		}
// 	}
// 	return
// }
