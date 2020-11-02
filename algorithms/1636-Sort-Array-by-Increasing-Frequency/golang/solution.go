package golang

import (
	"sort"
)

func frequencySort(nums []int) []int {
	dict := make(map[int]int)

	for _, val := range nums {
		count, ok := dict[val]
		if ok {
			dict[val] = count + 1
		} else {
			dict[val] = 1
		}
	}

	keys := getMapSortedKeys(dict)
	result := make([]int, 0, len(nums))

	for _, k := range keys {
		val := dict[k]
		list := make([]int, 0, val)
		for i := 0; i < val; i++ {
			list = append(list, k)
		}

		result = append(result, list...)
	}

	return result
}

func getMapSortedKeys(dict map[int]int) []int {
	keys := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i int, j int) bool {
		iVal := dict[keys[i]]
		jVal := dict[keys[j]]

		if iVal != jVal {
			return iVal < jVal
		} else {
			return keys[i] > keys[j]
		}
	})

	return keys
}
