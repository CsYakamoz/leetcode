package golang

func checkIfExist(arr []int) bool {
	dict := make(map[int]bool)

	for _, v := range arr {
		_, e1 := dict[v*2]
		if e1 {
			return true
		}

		if v%2 == 0 {
			_, e2 := dict[v/2]
			if e2 {
				return true
			}
		}

		dict[v] = true
	}

	return false
}
