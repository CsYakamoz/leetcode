package golang

import (
	"math"
)

func beautySum(s string) int {
	var dict [26]int

	ans := 0
	prefix := [][26]int{}

	base := int('a')
	for _, c := range s {
		idx := int(c) - base
		dict[idx]++

		ans += calcBeauty(&dict)

		prefix = append(prefix, dict)
	}

	length := len(s)
	for i := 2; i < length; i++ {
		for j := 0; j < length-i; j++ {
			min, max := math.MaxInt32, math.MinInt32

			for k := 0; k < 26; k++ {
				count := prefix[j+i][k] - prefix[j][k]
				if count == 0 {
					continue
				}

				if count < min {
					min = count
				}

				if count > max {
					max = count
				}
			}

			ans += max - min
		}
	}

	return ans
}

func calcBeauty(dict *[26]int) int {
	min, max := math.MaxInt32, math.MinInt32

	for _, val := range *dict {
		if val == 0 {
			continue
		}

		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	return max - min
}
