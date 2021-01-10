package golang

func waysToSplit(nums []int) int {
	modulo := int(1e9 + 7)
	sum, length := 0, len(nums)
	prefixSum := make([]int, length, length)

	for i, val := range nums {
		sum += val
		prefixSum[i] = sum
	}

	if sum == 0 {
		return int(float64(1+length-2)*float64(length-2)/2) % modulo
	}

	ans := 0

	for i := 0; i < length-2; i++ {
		left := prefixSum[i]

		maxMid := (sum - left) / 2
		if maxMid < left {
			break
		}

		minIdx := binarySearch(
			i+1,
			length,
			prefixSum,
			2*left,
		)

		maxIdx := binarySearch(
			i+1,
			length,
			prefixSum,
			left+maxMid+1,
		)

		ans += maxIdx - minIdx
		ans %= modulo
	}

	return ans
}

func binarySearch(begin, end int, prefixSum []int, target int) int {
	for begin < end {
		mid := (end-begin)/2 + begin

		if prefixSum[mid] >= target {
			end = mid
		} else {
			begin = mid + 1
		}
	}

	return end
}
