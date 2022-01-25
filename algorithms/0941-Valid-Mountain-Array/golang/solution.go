package golang

func validMountainArray(arr []int) bool {
	length := len(arr)

	// required: arr.length >= 3, and strictly increasing at first
	if length < 3 || arr[0] >= arr[1] {
		return false
	}

	dir := true // true -> inc, false -> dec
	for i := 2; i < length; i++ {
		// required: strictly increasing or decreasing
		if arr[i] == arr[i-1] {
			return false
		}

		if arr[i] > arr[i-1] {
			if !dir {
				return false
			}
		} else {
			if dir {
				dir = false
			}
		}
	}

	return !dir
}
