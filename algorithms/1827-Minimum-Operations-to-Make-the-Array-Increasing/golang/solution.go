package golang

func minOperations(nums []int) int {
	ans, prev, length := 0, nums[0], len(nums)

	for i := 1; i < length; i++ {
		if prev >= nums[i] {
			ans += prev - nums[i] + 1
			prev = prev + 1
		} else {
			prev = nums[i]
		}
	}

	return ans
}
