## [1573. Number of Ways to Split a String](https://leetcode.com/problems/number-of-ways-to-split-a-string/)

### Description

Given a binary string `s` (a string consisting only of '0's and '1's), we can split `s` into 3 **non-empty** strings s1, s2, s3 (s1+ s2+ s3 = s).

Return the number of ways `s` can be split such that the number of characters '1' is the same in s1, s2, and s3.

Since the answer may be too large, return it modulo 10^9 + 7.

**Example 1:**

```
Input: s = "10101"
Output: 4
Explanation: There are four ways to split s in 3 parts where each part contain the same number of letters '1'.
"1|010|1"
"1|01|01"
"10|10|1"
"10|1|01"
```

**Example 2:**

```
Input: s = "1001"
Output: 0
```

**Example 3:**

```
Input: s = "0000"
Output: 3
Explanation: There are three ways to split s in 3 parts.
"0|0|00"
"0|00|0"
"00|0|0"
```

**Example 4:**

```
Input: s = "100100010100110"
Output: 12
```

**Constraints:**

- `3 <= s.length <= 10^5`
- `s[i]` is `'0'` or `'1'`.

**Difficulty:** `Medium`

**Tags:** `String`

### Solution One

```javascript
/**
 * @param {string} s
 * @return {number}
 */
const numWays = function (s) {
  const oneIdxRecord = [];
  const modulo = 10 ** 9 + 7;
  for (let i = 0, len = s.length; i < len; i++) {
    if (s[i] === '0') {
      continue;
    }

    oneIdxRecord.push(i);
  }

  if (oneIdxRecord.length % 3 !== 0) {
    return 0;
  }

  if (oneIdxRecord.length === 0) {
    return (((s.length - 1) * (s.length - 2)) / 2) % modulo;
  }

  const partCount = oneIdxRecord.length / 3;
  const firstPartEndingIdx = oneIdxRecord[partCount - 1];
  const secondPartStartingIdx = oneIdxRecord[partCount];
  const secondPartEndingIdx = oneIdxRecord[partCount * 2 - 1];
  const thirdPartStartingIdx = oneIdxRecord[partCount * 2];

  const zeroCountBtwFirstAndSecond =
    secondPartStartingIdx - firstPartEndingIdx - 1;
  const zeroCountBtwSecondeAndThird =
    thirdPartStartingIdx - secondPartEndingIdx - 1;

  return (
    ((zeroCountBtwFirstAndSecond + 1) * (zeroCountBtwSecondeAndThird + 1)) %
    modulo
  );
};
```
