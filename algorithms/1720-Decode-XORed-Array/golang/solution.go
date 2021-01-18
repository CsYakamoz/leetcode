package golang

func decode(encoded []int, first int) []int {
	length := len(encoded)
	ans := make([]int, length+1, length+1)
	ans[0] = first

	for i := 1; i <= length; i++ {
		ans[i] = ans[i-1] ^ encoded[i-1]
	}

	return ans
}
