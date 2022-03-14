package golang

import "fmt"

func executeInstructions(n int, startPos []int, s string) []int {
	length, result := len(s), []int{}
	dict := map[string]int{}

	for i := length - 1; i >= 0; i-- {
		instructions, pos := 0, make([]int, 2, 2)
		copy(pos, startPos)

		for j := i; j < length; j++ {
			direction := string(s[j])
			if count, exist := dict[genKey(j, pos)]; exist {
				instructions += count
				break
			}

			if done, idx, v := executing(n, pos, direction); done {
				pos[idx] += v
				instructions++
			} else {
				break
			}
		}

		key := genKey(i, startPos)
		dict[key] = instructions

		result = append(result, instructions)
	}

	reverse(result)
	return result
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func executing(n int, pos []int, direction string) (bool, int, int) {
	if direction == "U" && pos[0] > 0 {
		return true, 0, -1
	} else if direction == "D" && pos[0] < n-1 {
		return true, 0, 1
	} else if direction == "L" && pos[1] > 0 {
		return true, 1, -1
	} else if direction == "R" && pos[1] < n-1 {
		return true, 1, 1
	}

	return false, 0, 0
}

func genKey(idx int, pos []int) string {
	return fmt.Sprintf("%d-%d-%d", idx, pos[0], pos[1])
}
