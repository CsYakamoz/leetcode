## [1013. Partition Array Into Three Parts With Equal Sum](https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/)

### Description

Given an array `A` of integers, return `true` if and only if we can partition the array into three **non-empty** parts with equal sums.

Formally, we can partition the array if we can find indexes `i+1 < j` with `(A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])`

**Example 1:**

```
Input: A = [0,2,1,-6,6,-7,9,1,2,0,1]
Output: true
Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
```

**Example 2:**

```
Input: A = [0,2,1,-6,6,7,9,-1,2,0,1]
Output: false
```

**Example 3:**

```
Input: A = [3,3,6,5,-2,2,5,1,-9,4]
Output: true
Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
```

**Constraints:**

- `3 <= A.length <= 50000`
- `-10^4 <= A[i] <= 10^4`

**Difficulty:** `Easy`
**Tags:** `Array`

### Solution One

1. 遍历一遍 `A`, 得到 `sumList`, 其元素 `sumList[i]` 定义为 `A[0] + ... + A[i]`;
2. 判断数组 `A` 所有元素之和是否为 `3` 的倍数，不是则返回 `false`, 否则进行下一步；
3. 在 `sumList` 中寻找值为 `partSum` 和 `2 * partSum` 的元素，且其下标符合相关条件；

Time: O(n); Space: O(n)

```javascript
/**
 * @param {number[]} A
 * @return {boolean}
 */
const canThreePartsEqualSum = function (A) {
  let totalSum = 0;
  const sumList = [];

  for (const val of A) {
    totalSum += val;
    sumList.push(totalSum);
  }

  if (totalSum % 3 !== 0) {
    return false;
  }

  const partSum = totalSum / 3;

  const firstIdx = sumList.indexOf(partSum);
  if (firstIdx === -1 || firstIdx >= A.length - 2) {
    return false;
  }

  const secondIdx = sumList.indexOf(2 * partSum, firstIdx + 1);
  if (secondIdx === -1 || secondIdx >= A.length - 1) {
    return false;
  }

  return totalSum - sumList[secondIdx] === partSum;
};
```
