package golang

func carPooling(trips [][]int, capacity int) bool {
	// 0 <= trips[i][1] < trips[i][2] <= 1000
	minTrip, maxTrip := 1001, -1
	locationDict := map[int]int{}

	for _, trip := range trips {
		if trip[1] < minTrip {
			minTrip = trip[1]
		}
		if trip[2] > maxTrip {
			maxTrip = trip[2]
		}

		locationDict[trip[1]] += trip[0]
		locationDict[trip[2]] -= trip[0]
	}

	for passengers, i := 0, minTrip; i <= maxTrip; i++ {
		passengers += locationDict[i]
		if passengers > capacity {
			return false
		}
	}

	return true
}
