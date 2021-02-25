package golang

func scoreOfParentheses(S string) int {
	scoreStack := []int{}
	unclosedStack := []int{}

	for _, ch := range S {
		if ch == '(' {
			scoreStack = append(scoreStack, 1)
			unclosedStack = append(unclosedStack, len(scoreStack)-1)
			continue
		}

		scoreStackLength, unclosedStackLength :=
			len(scoreStack), len(unclosedStack)

		idx := unclosedStack[unclosedStackLength-1]
		if idx != scoreStackLength-1 {
			sum := 0
			for i := idx + 1; i < scoreStackLength; i++ {
				sum += scoreStack[i]
			}
			scoreStack[idx] = 2 * sum
		}

		unclosedStack = unclosedStack[:unclosedStackLength-1]
		scoreStack = scoreStack[:idx+1]
	}

	ans := 0
	for _, score := range scoreStack {
		ans += score
	}

	return ans
}
