## [1827. Minimum Operations to Make the Array Increasing](https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/)

### Description

<p>You are given an integer array <code>nums</code> (<strong>0-indexed</strong>). In one operation, you can choose an element of the array and increment it by <code>1</code>.</p>

<ul>
	<li>For example, if <code>nums = [1,2,3]</code>, you can choose to increment <code>nums[1]</code> to make <code>nums = [1,<u><b>3</b></u>,3]</code>.</li>
</ul>

<p>Return <em>the <strong>minimum</strong> number of operations needed to make</em> <code>nums</code> <em><strong>strictly</strong> <strong>increasing</strong>.</em></p>

<p>An array <code>nums</code> is <strong>strictly increasing</strong> if <code>nums[i] &lt; nums[i+1]</code> for all <code>0 &lt;= i &lt; nums.length - 1</code>. An array of length <code>1</code> is trivially strictly increasing.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,1,1]
<strong>Output:</strong> 3
<strong>Explanation:</strong> You can do the following operations:
1) Increment nums[2], so nums becomes [1,1,<u><strong>2</strong></u>].
2) Increment nums[1], so nums becomes [1,<u><strong>2</strong></u>,2].
3) Increment nums[2], so nums becomes [1,2,<u><strong>3</strong></u>].
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,5,2,4,1]
<strong>Output:</strong> 14
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [8]
<strong>Output:</strong> 0
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 5000</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
</ul>

**Difficulty:** `Easy`

**Tags:** `Array` `Greedy`

### Solution One

```go
func minOperations(nums []int) int {
	ans, prev, length := 0, nums[0], len(nums)

	for i := 1; i < length; i++ {
		if prev >= nums[i] {
			ans += prev - nums[i] + 1
			prev = prev + 1
		} else {
			prev = nums[i]
		}
	}

	return ans
}
```