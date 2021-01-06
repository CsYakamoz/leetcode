## [1658. Minimum Operations to Reduce X to Zero](https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/)

### Description

<p>You are given an integer array <code>nums</code> and an integer <code>x</code>. In one operation, you can either remove the leftmost or the rightmost element from the array <code>nums</code> and subtract its value from <code>x</code>. Note that this <strong>modifies</strong> the array for future operations.</p>

<p>Return <em>the <strong>minimum number</strong> of operations to reduce </em><code>x</code> <em>to <strong>exactly</strong></em> <code>0</code> <em>if it&#39;s possible</em><em>, otherwise, return </em><code>-1</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,1,4,2,3], x = 5
<strong>Output:</strong> 2
<strong>Explanation:</strong> The optimal solution is to remove the last two elements to reduce x to zero.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [5,6,7,8,9], x = 4
<strong>Output:</strong> -1
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,2,20,1,1,3], x = 10
<strong>Output:</strong> 5
<strong>Explanation:</strong> The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= x &lt;= 10<sup>9</sup></code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Two Pointers` `Binary Search` `Greedy`

### Solution One

```go
func minOperations(nums []int, x int) int {
	// 1 <= nums.length <= 10^5
	ans, length, dict :=
		int(10e5+1), len(nums), map[int]int{0: -1}

	for i, sum := 0, 0; i < length; i++ {
		sum += nums[i]
		dict[sum] = i

		if sum == x {
			ans = i + 1
		}

		if sum >= x {
			break
		}
	}

	for i, sum := length-1, 0; i >= 0; i-- {
		sum += nums[i]

		anotherVal := x - sum
		idx, exist := dict[anotherVal]
		if exist && idx < i {
			ans = min(ans, length-i+idx+1)
		}

		if sum >= x {
			break
		}
	}

	if ans == 10e5+1 {
		return -1
	} else {
		return ans
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
```

### Solution Two - In Top Solutions

如同 `Hint` 所说, 我们可以反其道而行, 不是寻找最小的前后缀数组, 而是寻找最大的子数组

具体可看如下链接:

[[Java] Detailed Explanation - O(N) Prefix Sum/Map - Longest Target Sub-Array](<https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/discuss/935935/Java-Detailed-Explanation-O(N)-Prefix-SumMap-Longest-Target-Sub-Array>)

[C# Sliding window O(n) Time O(1) Space](<https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/discuss/935974/C-Sliding-window-O(n)-Time-O(1)-Space>)
