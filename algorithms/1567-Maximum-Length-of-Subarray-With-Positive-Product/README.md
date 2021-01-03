## [1567. Maximum Length of Subarray With Positive Product](https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/)

### Description

<p>Given an array of integers&nbsp;<code>nums</code>, find&nbsp;the maximum length of a subarray where the product of all its elements is positive.</p>

<p>A subarray of an array is a consecutive sequence of zero or more values taken out of that array.</p>

<p>Return&nbsp;<em>the maximum length of a subarray with positive product</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,-2,-3,4]
<strong>Output:</strong> 4
<strong>Explanation: </strong>The array nums already has a positive product of 24.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,1,-2,-3,-4]
<strong>Output:</strong> 3
<strong>Explanation: </strong>The longest subarray with positive product is [1,-2,-3] which has a product of 6.
Notice that we cannot include 0 in the subarray since that&#39;ll make the product 0 which is not positive.</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [-1,-2,-3,0,1]
<strong>Output:</strong> 2
<strong>Explanation: </strong>The longest subarray with positive product is [-1,-2] or [-2,-3].
</pre>

<p><strong>Example 4:</strong></p>

<pre>
<strong>Input:</strong> nums = [-1,2]
<strong>Output:</strong> 1
</pre>

<p><strong>Example 5:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,2,3,5,-6,4,0,10]
<strong>Output:</strong> 4
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10^5</code></li>
	<li><code>-10^9 &lt;= nums[i]&nbsp;&lt;= 10^9</code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Greedy`

### Solution One

```go
func getMaxLen(nums []int) int {
	prevZeroIdx := -1
	longest, length := 0, len(nums)
	negNumIdxRecord := make([]int, 0)

	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			longest = max(
				longest,
				calcMaxLength(negNumIdxRecord, prevZeroIdx, i),
			)

			prevZeroIdx = i
			negNumIdxRecord = negNumIdxRecord[0:0]
		} else if nums[i] < 0 {
			negNumIdxRecord = append(negNumIdxRecord, i)
		}
	}

	longest = max(
		longest,
		calcMaxLength(negNumIdxRecord, prevZeroIdx, length),
	)

	return longest

}

func calcMaxLength(negNumIdxRecord []int, prevZeroIdx, i int) int {
	recordLength := len(negNumIdxRecord)

	if recordLength%2 == 0 {
		return i - prevZeroIdx - 1
	} else {
		first, last :=
			negNumIdxRecord[0], negNumIdxRecord[recordLength-1]
		return max(last-prevZeroIdx-1, i-first-1)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
```
