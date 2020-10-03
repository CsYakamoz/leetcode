## [201. Bitwise AND of Numbers Range](https://leetcode.com/problems/bitwise-and-of-numbers-range/)

### Description

Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.

**Example 1:**

```
Input: [5,7]
Output: 4
```

**Example 2:**

```
Input: [0,1]
Output: 0
```

**Difficult:** `Medium`

**Tags:** `Bit Manipulation`

### Solution One

```javascript
/**
 * @param {number} diff
 * @return {number}
 */
const findTarget = (diff) => {
  let target = 1;

  while (target <= diff) {
    target *= 2;
  }

  return ~(target - 1);
};

/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
const rangeBitwiseAnd = function (m, n) {
  if (m === n) {
    return m;
  }

  return m & n & findTarget(n - m);
};
```
