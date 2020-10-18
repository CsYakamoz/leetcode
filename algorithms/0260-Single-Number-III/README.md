## [260. Single Number III](https://leetcode.com/problems/single-number-iii/description/)

### Description

Given an array of numbers `nums`, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

For example:

Given `nums = [1, 2, 1, 3, 2, 5]`, return `[3, 5]`.

**Note**:

1. The order of the result is not important. So in the above example, `[5, 3]` is also correct.
2. Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

**Difficulty:** `Medium`

**Tags:** `Bit Manipulation`

### Solution One

[感受异或的神奇](https://www.lijinma.com/blog/2014/05/29/amazing-xor/)

```c++
class Solution {
public:
    vector<int> singleNumber(vector<int>& nums) {
        int res = 0;
        for (auto i : nums)
            res ^= i;

        int bits = res & ~(res - 1);

        int a = 0;
        for (auto i : nums)
            if (i & bits)
                a ^= i;
        return { a, res ^ a };
    }
};
```
