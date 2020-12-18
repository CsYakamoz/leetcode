package golang

func getSumAbsoluteDifferences(nums []int) []int {
	length := len(nums)
	prefixSum := make([]int, length)

	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
		prefixSum[i] = sum
	}

	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] += (nums[i] * (i + 1)) - prefixSum[i]
		result[i] += (sum - prefixSum[i]) - (length-(i+1))*nums[i]
	}

	return result
}
