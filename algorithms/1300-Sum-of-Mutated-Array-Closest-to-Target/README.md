## [1300. Sum of Mutated Array Closest to Target](https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/)

### Description

<p>Given an integer array&nbsp;<code>arr</code> and a target value <code>target</code>, return&nbsp;the integer&nbsp;<code>value</code>&nbsp;such that when we change all the integers&nbsp;larger than <code>value</code>&nbsp;in the given array to be equal to&nbsp;<code>value</code>,&nbsp;the sum of the array gets&nbsp;as close as possible (in absolute difference) to&nbsp;<code>target</code>.</p>

<p>In case of a tie, return the minimum such integer.</p>

<p>Notice that the answer is not neccesarilly a number from <code>arr</code>.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> arr = [4,9,3], target = 10
<strong>Output:</strong> 3
<strong>Explanation:</strong> When using 3 arr converts to [3, 3, 3] which sums 9 and that&#39;s the optimal answer.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> arr = [2,3,5], target = 10
<strong>Output:</strong> 5
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> arr = [60864,25176,27249,21296,20204], target = 56803
<strong>Output:</strong> 11361
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
  <li><code>1 &lt;= arr.length &lt;= 10^4</code></li>
  <li><code>1 &lt;= arr[i], target &lt;= 10^5</code></li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Array` `Binary Search`

### Solution One

1. 将 `arr` 按照从小到大排序

2. 遍历 `arr`, 找到第一个 `changedSum` 大于等于 `target` 的下标

   `changedSum` 定义为：

   下标在 `[0, idx]` 范围内的整数保持不变；

   下标在 `(idx, len(arr))` 范围内的整数改成 `changedVal`;

   此时 `sum(arr)` 等于 `arr[0] + ... + arr[i] + changedVal * (len(arr) - i - 1)`;

   注意：步骤二中的 `changedVal` 等于 `arr[i]`;

3. 从第二步可得到 `idx`, 根据其值可分成三种情况：

   1. idx == len(arr): 原数组的总和依旧小于等于 `target`, 故直接返回最大值即可；

   2. idx == 0: 即使最小值也大于等于 `target`, 那么整个数组的元素都改成 `target / len(arr)` 或者 `(target / len(arr)) + 1` 即可；

   3. 0 < idx < len(arr): 证明当前数组中，下标为 `idx - 1` 和 `idx` 的 `changedSum` 最接近 `target`

      由于题目说答案并非须是数组元素，因此还需从 `arr[idx - 1]` 到 `arr[idx]` 中找到最接近的值；

      这里使用二分法寻找范围中第一个 `changedSum` (`changedVal` 为 `mid`) 大于等于 `target` 的值，此时得到 `end`;

      再分别计算 `end-1` 和 `end` 的 `changedSum`, 从而得到谁更接近 `target`;

```go
import (
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	length := len(arr)

	idx, prevSum := getIdxPrevSum(target, length, arr)

	if idx == length {
		return arr[length-1]
	} else if idx == 0 {
		mod := target % length
		quotient := target / length

		if mod <= (length / 2) {
			return quotient
		} else {
			return quotient + 1
		}
	} else {
		begin, end := arr[idx-1], arr[idx]
		for begin < end {
			mid := (end-begin)/2 + begin
			midSum := getChangedSum(idx-1, length, mid, prevSum)

			if midSum >= target {
				end = mid
			} else {
				begin = mid + 1
			}
		}

		x, y :=
			getChangedSum(idx-1, length, end-1, prevSum),
			getChangedSum(idx-1, length, end, prevSum)

		return helper(x, y, target, end)
	}
}

func helper(x int, y int, target int, end int) int {
	if (target - x) <= (y - target) {
		return end - 1
	} else {
		return end
	}
}

func getChangedSum(idx int, length int, changedVal int, idxSum int) int {
	return idxSum + changedVal*(length-idx-1)
}

func getIdxPrevSum(target int, length int, arr []int) (int, int) {
	prevSum, idxSum := 0, 0
	for idx := 0; idx < length; idx++ {
		val := arr[idx]
		idxSum += val

		if idx != 0 {
			prevSum += arr[idx-1]
		}

		if getChangedSum(idx, length, val, idxSum) >= target {
			return idx, prevSum
		}
	}

	return length, 0
}
```
