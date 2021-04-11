## [1390. Four Divisors](https://leetcode.com/problems/four-divisors/)

### Description

<p>Given an integer array <code>nums</code>, return the sum of divisors of the integers in that array that have exactly four divisors.</p>

<p>If there is no such integer in the array, return <code>0</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [21,4,7]
<strong>Output:</strong> 32
<b>Explanation:</b>
21 has 4 divisors: 1, 3, 7, 21
4 has 3 divisors: 1, 2, 4
7 has 2 divisors: 1, 7
The answer is the sum of divisors of 21 only.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10^4</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10^5</code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Math`

### Solution One

```go
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
```
