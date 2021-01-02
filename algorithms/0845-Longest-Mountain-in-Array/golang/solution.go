package golang

func findLeftBottom(arr []int, length, begin int) int {
	for i := begin + 1; i < length; i++ {
		if arr[i] > arr[i-1] {
			return i - 1
		}
	}

	return length
}

func findTop(arr []int, length, begin int) int {
	for i := begin + 1; i < length; i++ {
		if arr[i] <= arr[i-1] {
			return i - 1
		}
	}

	return length
}

func findRightBottom(arr []int, length, begin int) int {
	for i := begin + 1; i < length; i++ {
		if arr[i] >= arr[i-1] {
			return i - 1
		}
	}

	return length - 1
}

func longestMountain(arr []int) int {
	length := len(arr)

	// subarray.length should be >= 3
	if length < 3 {
		return 0
	}

	largestMountain, point := 0, 0
	for point < length {
		left := findLeftBottom(arr, length, point)
		if left >= length-2 {
			return largestMountain
		}

		top := findTop(arr, length, left+1)
		if top >= length-1 {
			return largestMountain
		}

		right := findRightBottom(arr, length, top)
		if right > top {
			subarrayLength := right - left + 1
			if subarrayLength > largestMountain {
				largestMountain = subarrayLength
			}
		}

		point = right
	}

	return largestMountain
}
