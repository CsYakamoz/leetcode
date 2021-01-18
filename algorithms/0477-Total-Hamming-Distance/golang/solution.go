package golang

import (
	"strconv"
)

func totalHammingDistance(nums []int) int {
	ans, length, dict := 0, len(nums), map[int]int{}

	for _, val := range nums {
		str := strconv.FormatInt(int64(val), 2)
		length := len(str)

		for i := length - 1; i >= 0; i-- {
			if str[i] == '0' {
				continue
			}
			dict[length-i]++
		}
	}

	for _, count := range dict {
		ans += count * (length - count)
	}

	return ans
}
