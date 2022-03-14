package golang

import (
	"math"
)

func calcTrip(dict *map[int]int, time int64) int64 {
	res := int64(0)
	for k, v := range *dict {
		res += time / int64(v) * int64(k)
	}

	return res
}

func minimumTime(time []int, totalTrips int) int64 {
	d, maxCost, dict := 0.0, 0, map[int]int{}

	for _, v := range time {
		dict[v]++
		if v > maxCost {
			maxCost = v
		}
	}

	for k, v := range dict {
		d += float64(v) / float64(k)
	}

	left, right := int64(math.Ceil(float64(totalTrips)/d)), int64((^uint64(0) >> 1))
	for left < right {
		mid := left + (right-left)/2
		trip := calcTrip(&dict, mid)
		if trip >= int64(totalTrips) {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}
