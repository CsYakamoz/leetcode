package golang

func minOperations(nums []int, x int) int {
	// 1 <= nums.length <= 10^5
	ans, length := int(10e5+1), len(nums)

	prefixSum, prefixSumLength := make([]int, length, length), 0
	for i, sum := 0, 0; i < length; i++ {
		sum += nums[i]
		prefixSum[i] = sum
		prefixSumLength++

		if sum == x {
			ans = i + 1
		}

		if sum >= x {
			break
		}
	}

	for i, sum := length-1, 0; i >= 0; i-- {
		sum += nums[i]

		anotherVal := x - sum
		if anotherVal == 0 {
			ans = min(ans, length-i)
		} else {
			idx := binarySearch(
				prefixSum[0:min(prefixSumLength, i)],
				anotherVal,
			)
			if idx != -1 {
				ans = min(ans, length-i+idx+1)
			}
		}

		if sum >= x {
			break
		}
	}

	if ans == 10e5+1 {
		return -1
	} else {
		return ans
	}
}

func binarySearch(prefixSum []int, target int) int {
	length := len(prefixSum)
	begin, end := 0, length
	for begin < end {
		mid := (end-begin)/2 + begin
		if prefixSum[mid] >= target {
			end = mid
		} else {
			begin = mid + 1
		}
	}

	if end < length && prefixSum[end] == target {
		return end
	} else {
		return -1
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
