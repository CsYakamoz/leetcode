package golang

import (
	"math"
)

func sumFourDivisors(nums []int) int {
	ans := 0

	for _, val := range nums {
		ans += fourDivisorsSumOr0(val)
	}

	return ans
}

func fourDivisorsSumOr0(num int) int {
	count, sum, root := 0, 0, int(math.Sqrt(float64(num)))

	for i := 2; i <= root; i++ {
		if num%i == 0 {
			if i*i == num {
				count += 1
				sum += i
			} else {
				count += 2
				sum += i + num/i
			}
		}
	}

	if count != 2 {
		return 0
	} else {
		return sum + 1 + num
	}
}
