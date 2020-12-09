package golang

import (
	"math"
)

func countPrimeSetBits(L int, R int) int {
	result := 0
	dict := make(map[int]int)
	dict[1] = 0

	if L == 1 {
		L += 1
	}

	for i := L; i <= R; i++ {
		bit := countBit(i)

		v1, exist := dict[bit]
		if exist {
			result += v1
			continue
		}

		v2 := isPrime(bit)
		dict[bit] = v2
		result += v2
	}

	return result
}

func countBit(num int) int {
	count := 0
	for num != 0 {
		count += 1
		num = num & (num - 1)
	}

	return count
}

func isPrime(num int) int {
	endCondition := int(math.Sqrt(float64(num)))
	for i := 2; i <= endCondition; i++ {
		if num%i == 0 {
			return 0
		}
	}

	return 1
}
