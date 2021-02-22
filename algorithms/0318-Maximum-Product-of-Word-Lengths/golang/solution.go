package golang

import (
// "fmt"
// "strconv"
)

type unit struct {
	bit        int
	wordLength int
}

func maxProduct(words []string) int {
	ans, length := 0, len(words)
	list := make([]unit, length, length)

	for i, word := range words {
		bit, wordLength := wordToBits(word), len(word)

		for j := 0; j < i; j++ {
			if bit&list[j].bit == 0 {
				ans = max(ans, wordLength*list[j].wordLength)
			}
		}

		list[i].bit = bit
		list[i].wordLength = wordLength
	}

	return ans
}

func wordToBits(word string) int {
	ans := 0

	for _, ch := range word {
		ans |= 1 << (int(ch) - 97)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
