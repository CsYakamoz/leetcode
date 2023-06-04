package golang

func validateStackSequences(pushed []int, popped []int) bool {
	length := len(pushed)
	stack := make([]int, 0, length)

	j := 0
	for i := 0; i < length; i++ {
		for len(stack) != 0 && popped[j] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			j++
		}

		stack = append(stack, pushed[i])
	}

	for ; len(stack) != 0 && j < length && popped[j] == stack[len(stack)-1]; j++ {
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0 && j == length
}
