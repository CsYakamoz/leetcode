package golang

import (
	"strconv"
)

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		char := num[i : i+1]
		unit, _ := strconv.Atoi(char)

		if unit%2 == 0 {
			continue
		}

		return num[:i+1]
	}

	return ""
}
