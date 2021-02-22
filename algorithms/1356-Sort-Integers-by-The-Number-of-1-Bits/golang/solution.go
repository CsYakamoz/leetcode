package golang

import (
	"sort"
)

type unit struct {
	bits int
	n    int
}

func sortByBits(arr []int) []int {
	length := len(arr)
	dump := make([]unit, length, length)
	ans := make([]int, length, length)

	for idx, n := range arr {
		dump[idx].n = n
		dump[idx].bits = countBits(n)
	}

	sort.Slice(dump, func(i int, j int) bool {
		iVal, jVal := dump[i], dump[j]
		if iVal.bits != jVal.bits {
			return iVal.bits < jVal.bits
		}

		return iVal.n < jVal.n
	})

	for idx, item := range dump {
		ans[idx] = item.n
	}

	return ans
}

func countBits(n int) int {
	bits := 0

	for n != 0 {
		n &= (n - 1)
		bits++
	}

	return bits
}
