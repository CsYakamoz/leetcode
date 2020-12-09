## [762. Prime Number of Set Bits in Binary Representation](https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/)

### Description

<p>
Given two integers <code>L</code> and <code>R</code>, find the count of numbers in the range <code>[L, R]</code> (inclusive) having a prime number of set bits in their binary representation.
</p><p>
(Recall that the number of set bits an integer has is the number of <code>1</code>s present when written in binary.  For example, <code>21</code> written in binary is <code>10101</code> which has 3 set bits.  Also, 1 is not a prime.)
</p><p>

<p><b>Example 1:</b><br /><pre>
<b>Input:</b> L = 6, R = 10
<b>Output:</b> 4
<b>Explanation:</b>
6 -> 110 (2 set bits, 2 is prime)
7 -> 111 (3 set bits, 3 is prime)
9 -> 1001 (2 set bits , 2 is prime)
10->1010 (2 set bits , 2 is prime)
</pre></p>

<p><b>Example 2:</b><br /><pre>
<b>Input:</b> L = 10, R = 15
<b>Output:</b> 5
<b>Explanation:</b>
10 -> 1010 (2 set bits, 2 is prime)
11 -> 1011 (3 set bits, 3 is prime)
12 -> 1100 (2 set bits, 2 is prime)
13 -> 1101 (3 set bits, 3 is prime)
14 -> 1110 (3 set bits, 3 is prime)
15 -> 1111 (4 set bits, 4 is not prime)
</pre></p>

<p><b>Note:</b><br><ol>
<li><code>L, R</code> will be integers <code>L <= R</code> in the range <code>[1, 10^6]</code>.</li>
<li><code>R - L</code> will be at most 10000.</li>
</ol></p>

**Difficulty:** `Easy`

**Tags:** `Bit Manipulation`

### Solution One

```golang
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
```

### Solution Two - In Top Solutions

TODO: explain this

```golang
func countPrimeSetBits(L int, R int) int {
	rev := 0
	for L <= R {
		rev += 665772 >> bits.OnesCount64(uint64(L)) & 1
		L++
	}
	return rev
}
```
