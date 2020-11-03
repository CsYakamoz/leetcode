package golang

import (
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	dict := make(map[int]int)
	for _, val := range arr {
		dict[val]++
	}

	countArr := getSortedCountArr(dict)
	length := len(countArr)

	for _, count := range countArr {
		if k < count {
			continue
		}

		k -= count
		length--
	}

	return length
}

func getSortedCountArr(dict map[int]int) []int {
	result := make([]int, 0, len(dict))

	for _, count := range dict {
		result = append(result, count)
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i] < result[j]
	})

	return result
}
