## [1200. Minimum Absolute Difference](https://leetcode.com/problems/minimum-absolute-difference/)

### Description

Given an array of **distinct** integers `arr`, find all pairs of elements with the minimum absolute difference of any two elements.

Return a list of pairs in ascending order(with respect to pairs), each pair `[a, b]` follows

- `a, b` are from `arr`
- `a < b`
- `b - a` equals to the minimum absolute difference of any two elements in `arr`

**Example 1:**

```
Input: arr = [4,2,1,3]
Output: [[1,2],[2,3],[3,4]]
Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
```

**Example 2:**

```
Input: arr = [1,3,6,10,15]
Output: [[1,3]]
```

**Example 3:**

```
Input: arr = [3,8,-10,23,19,-4,-14,27]
Output: [[-14,-10],[19,23],[23,27]]
```

**Constraints:**

- `2 <= arr.length <= 10^5`
- `-10^6 <= arr[i] <= 10^6`

**Difficulty:** `Easy`

**Tags:** `Array`

### Solution One

```javascript
/**
 * @param {number[]} arr
 * @return {number[][]}
 */
const minimumAbsDifference = function (arr) {
  arr.sort((a, b) => a - b);

  const dict = new Map();
  let minDiff = Number.MAX_SAFE_INTEGER;
  for (let i = 1, len = arr.length; i < len; i++) {
    const diff = arr[i] - arr[i - 1];
    minDiff = Math.min(minDiff, diff);

    if (!dict.has(diff)) {
      dict.set(diff, []);
    }

    const list = dict.get(diff);
    list.push([arr[i - 1], arr[i]]);
  }

  return dict.get(minDiff) || [];
};
```
