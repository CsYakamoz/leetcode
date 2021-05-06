package golang

import (
	"sort"
)

type maxFreqUnit struct {
	num   int
	count int
}

func maxFrequency(nums []int, k int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num]++
	}

	list := make([]maxFreqUnit, 0, len(dict))
	for k, v := range dict {
		list = append(list, maxFreqUnit{
			num:   k,
			count: v,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].num < list[j].num
	})

	ans := 0
	for i := 0; i < len(list); i++ {
		remainedK, maxFreq := k, list[i].count

		for j := i - 1; j >= 0; j-- {
			diff := list[i].num - list[j].num

			if remainedK < diff {
				break
			} else {
				if remainedK >= list[j].count*diff {
					maxFreq += list[j].count
					remainedK -= list[j].count * diff
				} else {
					maxFreq += remainedK / diff
					break
				}
			}
		}

		ans = max(ans, maxFreq)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
