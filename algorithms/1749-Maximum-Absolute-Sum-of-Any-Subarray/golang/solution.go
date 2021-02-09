package golang

func maxAbsoluteSum(nums []int) int {
	length := len(nums)

	ans, prevMin, prevMax := 0, 0, 0
	for i, sum := 0, 0; i < length; i++ {
		sum += nums[i]
		ans = max(
			ans,
			max(abs(sum-prevMax), abs(sum-prevMin)),
		)

		if sum > prevMax {
			prevMax = sum
		}

		if sum < prevMin {
			prevMin = sum
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
