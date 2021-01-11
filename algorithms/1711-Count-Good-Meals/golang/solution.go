package golang

import (
	"sort"
)

func countPairs(deliciousness []int) int {
	sort.Ints(deliciousness)

	dict := map[int]int{}
	ans, modulo := 0, int(1e9+7)

	for _, val := range deliciousness {
		anotherVal := findAnotherVal(val)
		count, exist := dict[anotherVal]
		if exist {
			ans = (ans + count) % modulo
		}

		if val != 0 && val&(val-1) == 0 {
			count, exist := dict[val]
			if exist {
				ans = (ans + count) % modulo
			}
		}

		dict[val]++
	}

	return ans
}

func findAnotherVal(val int) int {
	init := 1

	for init < val {
		init <<= 1
	}

	return init - val
}
