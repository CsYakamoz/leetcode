package golang

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func maximumUniqueSubarray(nums []int) int {
	result, begin, dict, length :=
		0, 0, make(map[int]int), len(nums)
	sum, prefixSum := 0, make([]int, length+1)

	for end := 0; end < length; end++ {
		val := nums[end]
		sum += val
		prefixSum[end+1] = sum
		idx, exist := dict[val]

		if exist {
			result = max(prefixSum[end]-prefixSum[begin], result)
			begin = max(idx+1, begin)
		}

		dict[val] = end
	}

	return max(result, sum-prefixSum[begin])
}
