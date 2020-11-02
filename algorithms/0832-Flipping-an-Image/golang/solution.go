package golang

func flipAndInvertImage(A [][]int) [][]int {
	length := len(A)
	result := make([][]int, 0, length)

	for _, val := range A {
		row := make([]int, 0, length)

		for i := 0; i < length; i++ {
			row = append(row, 1-val[length-1-i])
		}

		result = append(result, row)
	}

	return result
}

