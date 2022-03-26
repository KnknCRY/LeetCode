package problems

import "fmt"

// unsubmit
func ZigzagConversion(s string, rows int) string {
	arr := [][]string{}

	// 先讓arr的rows長出來
	for i := 0; i < rows; i++ {
		arr = append(arr, []string{})
	}

	// increse是true代表往下走，是false代表往上走
	var increse bool = true
	var i int = 0

	for idx := 0; idx < len(s); idx++ {
		fmt.Println("start", i)
		// 下方else if i == rows裡面的i-2結果，有可能會小於0
		// 或是回到第一row的時候，因為此時是increse = false 會i--變成-1
		// 所以要把i歸零後再把increse = true
		if i <= 0 {
			i = 0
			increse = true
		}

		// 看i走到哪，就appaend進去哪一個row
		if i < rows {
			arr[i] = append(arr[i], string(s[idx]))
			// fmt.Println("append", i)
		} else if i == rows {
			// 如果i走到最後一row
			i -= 2
			arr[i] = append(arr[i], string(s[idx]))
			// fmt.Println("append", i)
			increse = false
		}

		if increse {
			i++
		} else {
			i--
		}
	}
	return s
}
