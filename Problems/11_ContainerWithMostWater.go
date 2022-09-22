package problems

// Accepted
func MaxArea(height []int) int {
	result := 0
	j := len(height) - 1

	for i := 0; i < len(height) && i != j; {
		shorter, side := getShorterOne(height[i], height[j])
		area := shorter * (j - i)

		if area > result {
			result = area
		}

		if side == "L" {
			i++
		} else {
			j--
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
