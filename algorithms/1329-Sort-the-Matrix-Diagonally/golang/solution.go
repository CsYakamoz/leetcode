package golang

func diagonalSort(mat [][]int) [][]int {
	rowLen, colLen := len(mat), len(mat[0])

	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			row, col := i-1, j-1
			for row >= 0 && col >= 0 && mat[row][col] > mat[row+1][col+1] {
				mat[row][col], mat[row+1][col+1] =
					mat[row+1][col+1], mat[row][col]

				row--
				col--
			}
		}
	}

	return mat
}

