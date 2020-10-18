## [673. Number of Longest Increasing Subsequence](https://leetcode.com/problems/number-of-longest-increasing-subsequence/)

### Description

Given an unsorted array of integers, find the number of longest increasing subsequence.

**Example 1:**

```
Input: [1,3,5,4,7]
Output: 2
Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
```

**Example 2:**

```
Input: [2,2,2,2,2]
Output: 5
Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences' length is 1, so output 5.
```

**Note:** Length of the given array will be not exceed 2000 and the answer is guaranteed to be fit in 32-bit signed int.

**Difficulty:** `Medium`

**Tags:** `Dynamic Programming`

### Solution One

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
const findNumberOfLIS = function (nums) {
  if (nums.length === 0) {
    return 0;
  }

  let numOfLIS = 1; // 最长子序列的数量
  let lengthOfLIS = 1; // 最长子序列的长度

  const numOfItem = [1]; // 以 item[i] 结尾的最长子序列的数量
  const lengthOfItem = [1]; // 以 item[i] 结尾的最长子序列的长度

  for (let i = 1; i < nums.length; i++) {
    let num = 1;
    let length = 1;

    for (let j = 0; j < i; j++) {
      if (nums[i] > nums[j]) {
        if (lengthOfItem[j] + 1 === length) {
          num += numOfItem[j];
        } else if (lengthOfItem[j] + 1 > length) {
          num = numOfItem[j];
          length = lengthOfItem[j] + 1;
        }
      }
    }

    numOfItem.push(num);
    lengthOfItem.push(length);

    if (length === lengthOfLIS) {
      numOfLIS += num;
    } else if (length > lengthOfLIS) {
      numOfLIS = num;
      lengthOfLIS = length;
    }
  }

  return numOfLIS;
};
```

### Solutions

[673. Number of Longest Increasing Subsequence - Solution](https://leetcode.com/problems/number-of-longest-increasing-subsequence/solution/)
