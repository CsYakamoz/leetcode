## [1481. Least Number of Unique Integers after K Removals](https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/)

### Description

<p>Given an array of integers&nbsp;<code>arr</code>&nbsp;and an integer <code>k</code>.&nbsp;Find the <em>least number of unique integers</em>&nbsp;after removing <strong>exactly</strong> <code>k</code> elements<b>.</b></p>

<ol>
</ol>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input: </strong>arr = [5,5,4], k = 1
<strong>Output: </strong>1
<strong>Explanation</strong>: Remove the single 4, only 5 is left.
</pre>

<strong>Example 2:</strong>

<pre>
<strong>Input: </strong>arr = [4,3,1,1,3,3,2], k = 3
<strong>Output: </strong>2
<strong>Explanation</strong>: Remove 4, 2 and either one of the two 1s or three 3s. 1 and 3 will be left.</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= arr.length&nbsp;&lt;= 10^5</code></li>
	<li><code>1 &lt;= arr[i] &lt;= 10^9</code></li>
	<li><code>0 &lt;= k&nbsp;&lt;= arr.length</code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Array` `Sort`

### Solution One

思路:

1. 以值为 key, 值出现的次数为 value 建立 map;
2. 获取 value 的数组, 并按从小到达排序;
3. 排在前面的便是要移除的值所出现的次数;

```golang
import (
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	dict := make(map[int]int)
	for _, val := range arr {
		dict[val]++
	}

	countArr := getSortedCountArr(dict)
	length := len(countArr)

	for _, count := range countArr {
		if k < count {
			continue
		}

		k -= count
		length--
	}

	return length
}

func getSortedCountArr(dict map[int]int) []int {
	result := make([]int, 0, len(dict))

	for _, count := range dict {
		result = append(result, count)
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i] < result[j]
	})

	return result
}
```
