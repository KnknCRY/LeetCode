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

	fmt.Println("res", result)

	a := [][]int{}
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			y := map[int]int{}
			for _, value := range result[i] {
				y[value]++
			}
			for _, value := range result[j] {
				y[value]--
			}

			fmt.Println("比對idx", i, j)
			fmt.Println("比對val", result[i], result[j])

			var sum int
			for _, value := range y {
				sum += value
			}
			if sum != 0 {
				a = append(a)
				result = append(result[:j], result[j+1:]...)
			}
		}
	}
	return result
}
