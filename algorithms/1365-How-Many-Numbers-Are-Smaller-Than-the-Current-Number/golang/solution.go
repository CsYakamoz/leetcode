package golang

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	length, dict := len(nums), map[int]int{}

	sortedNums := make([]int, length, length)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	i, j := 0, 1
	for ; j < length; j++ {
		if sortedNums[j] == sortedNums[i] {
			continue
		}

		dict[sortedNums[i]] = i
		i = j
	}
	dict[sortedNums[i]] = i

	ans := make([]int, length, length)
	for idx, n := range nums {
		ans[idx] = dict[n]
	}

	return ans
}
