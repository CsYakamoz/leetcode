## [1636. Sort Array by Increasing Frequency](https://leetcode.com/problems/sort-array-by-increasing-frequency/)

### Description

<p>Given an array of integers <code>nums</code>, sort the array in <strong>increasing</strong> order based on the frequency of the values. If multiple values have the same frequency, sort them in <strong>decreasing</strong> order.</p>

<p>Return the <em>sorted array</em>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,1,2,2,2,3]
<strong>Output:</strong> [3,1,1,2,2,2]
<strong>Explanation:</strong> &#39;3&#39; has a frequency of 1, &#39;1&#39; has a frequency of 2, and &#39;2&#39; has a frequency of 3.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,3,1,3,2]
<strong>Output:</strong> [1,3,3,2,2]
<strong>Explanation:</strong> &#39;2&#39; and &#39;3&#39; both have a frequency of 2, so they are sorted in decreasing order.
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [-1,1,-6,4,5,-6,1,4,1]
<strong>Output:</strong> [5,-1,4,4,-6,-6,1,1,1]</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>-100 &lt;= nums[i] &lt;= 100</code></li>
</ul>

**Difficulty:** `Easy`

**Tags:** `Array` `Sort`

### Solution One

```go
import (
	"sort"
)

func frequencySort(nums []int) []int {
	dict := make(map[int]int)

	for _, val := range nums {
		count, ok := dict[val]
		if ok {
			dict[val] = count + 1
		} else {
			dict[val] = 1
		}
	}

	keys := getMapSortedKeys(dict)
	result := make([]int, 0, len(nums))

	for _, k := range keys {
		val := dict[k]
		list := make([]int, 0, val)
		for i := 0; i < val; i++ {
			list = append(list, k)
		}

		result = append(result, list...)
	}

	return result
}

func getMapSortedKeys(dict map[int]int) []int {
	keys := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i int, j int) bool {
		iVal := dict[keys[i]]
		jVal := dict[keys[j]]

		if iVal != jVal {
			return iVal < jVal
		} else {
			return keys[i] > keys[j]
		}
	})

	return keys
}
```
