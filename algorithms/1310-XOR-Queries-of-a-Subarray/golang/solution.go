package golang

func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, 0, len(arr))

	currentXorSum := 0
	for _, num := range arr {
		currentXorSum ^= num
		prefix = append(prefix, currentXorSum)
	}

	result := make([]int, 0, len(queries))

	for _, q := range queries {
		l, r := q[0], q[1]

		xor := prefix[r]
		if l != 0 {
			xor ^= prefix[l-1]
		}

		result = append(result, xor)
	}

	return result
}
