## [457. Circular Array Loop](https://leetcode.com/problems/circular-array-loop/)

### Description

<p>You are given a <b>circular</b> array <code>nums</code> of positive and negative integers. If a number <i>k</i> at an index is positive, then move forward <i>k</i> steps. Conversely, if it&#39;s negative (-<i>k</i>), move backward <i>k</i>&nbsp;steps. Since the array is circular, you may assume that the last element&#39;s next element is the first element, and the first element&#39;s previous element is the last element.</p>

<p>Determine if there is a loop (or a cycle) in <code>nums</code>. A cycle must start and end at the same index and the cycle&#39;s length &gt; 1. Furthermore, movements in a cycle must all follow a single direction. In other words, a cycle must not consist of both forward and backward movements.</p>

<p>&nbsp;</p>

<p><b>Example 1:</b></p>

<pre>
<b>Input:</b> [2,-1,1,2,2]
<b>Output:</b> true
<b>Explanation:</b> There is a cycle, from index 0 -&gt; 2 -&gt; 3 -&gt; 0. The cycle&#39;s length is 3.
</pre>

<p><b>Example 2:</b></p>

<pre>
<b>Input:</b> [-1,2]
<b>Output:</b> false
<b>Explanation:</b> The movement from index 1 -&gt; 1 -&gt; 1 ... is not a cycle, because the cycle&#39;s length is 1. By definition the cycle&#39;s length must be greater than 1.
</pre>

<p><b>Example 3:</b></p>

<pre>
<b>Input:</b> [-2,1,-1,-2,-2]
<b>Output:</b> false
<b>Explanation:</b> The movement from index 1 -&gt; 2 -&gt; 1 -&gt; ... is not a cycle, because movement from index 1 -&gt; 2 is a forward movement, but movement from index 2 -&gt; 1 is a backward movement. All movements in a cycle must follow a single direction.</pre>

<p>&nbsp;</p>

<p><b>Note:</b></p>

<ol>
	<li>-1000 &le;&nbsp;nums[i] &le;&nbsp;1000</li>
	<li>nums[i] &ne;&nbsp;0</li>
	<li>1 &le;&nbsp;nums.length &le; 5000</li>
</ol>

<p>&nbsp;</p>

<p><b>Follow up:</b></p>

<p>Could you solve it in <b>O(n)</b> time complexity and&nbsp;<strong>O(1)</strong> extra space complexity?</p>

**Difficulty:** `Medium`

**Tags:** `Array` `Two Pointers`

### Solution One

此处记录已遍历索引的方法：索引值加 1000000, 但这是由于题目中提到的限制 (-1000 ≤ nums[i] ≤ 1000 && nums[i] ≠ 0)

否则，一般情况下，个人使用布尔类型数组来存储，命名为 `visited`, 未访问过则为 false, 否则为 true.

思路：

1. 每次内循环开始前，先记录当前索引 (i) 的方向
2. 在第一个内循环中，如果发现当前索引 (idx) 已访问过，或者方向不正确，则结束循环
3. 走到第二个内循环时，有以下几种情况：
   1. 方向不一致，不满足要求，循环过程中依旧判断方向即可；
   2. 遇到之前（非本次循环）已遍历过的索引，不满足要求，依旧循环过程中判断方向即可；
   3. 遇到之前（本次循环）已遍历过的索引，意味着是单方向的循环，但需要检查长度是否大于 1;

```golang
var mark int = 1000000

// -1000 ≤ nums[i] ≤ 1000 && nums[i] ≠ 0
func isInCorrectRange(num int) bool {
	return -1000 <= num && num <= 1000
}

func getValAfterMark(num int) int {
	return num + mark
}

func getOriginalVal(num int) int {
	if isInCorrectRange(num) {
		return num
	}

	return num - mark
}

func getNextIdx(movement int, idx int, length int) int {
	trulyMov := movement % length
	return (trulyMov + idx + length) % length
}

func circularArrayLoop(nums []int) bool {
	length := len(nums)

	for i := 0; i < length; i++ {
		if !isInCorrectRange(nums[i]) {
			continue
		}

		direction, idx := nums[i] > 0, i

		for isInCorrectRange(nums[idx]) {
			movement := nums[idx]
			if (movement > 0) != direction {
				break
			}

			nums[idx] = getValAfterMark(movement)
			idx = getNextIdx(movement, idx, length)
		}

		curr, loopLength := idx, 0
		for {
			movement := getOriginalVal(nums[curr])
			if movement > 0 != direction || movement%length == 0 {
				break
			}
			curr = getNextIdx(movement, curr, length)
			loopLength++

			if curr == idx {
				break
			}
		}

		if curr == idx && loopLength > 1 {
			return true
		}
	}

	return false
}
```
