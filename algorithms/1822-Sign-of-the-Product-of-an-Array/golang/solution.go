package golang

func arraySign(nums []int) int {
	negativeCount := 0

	for _, val := range nums {
		if val == 0 {
			return 0
		}
		if val < 0 {
			negativeCount++
		}
	}

	if negativeCount%2 == 0 {
		return 1
	} else {
		return -1
	}
}
