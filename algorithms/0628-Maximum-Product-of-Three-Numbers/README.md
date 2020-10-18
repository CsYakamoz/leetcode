## [628. Maximum Product of Three Numbers](https://leetcode.com/problems/maximum-product-of-three-numbers/#/description)

### Description

```
Input: [1,2,3,4]
Output: 24

```

**Note:**

1. The length of the given array will be in range [3,10^4^] and all elements are in the range [-1000, 1000].
2. Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.

**Difficulty:** `Easy`

**Tags:** `Array` `Math`

### Solution One

```c++
class Solution {
public:
    int maximumProduct(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        return max(nums[0] * nums[1] * nums[nums.size() - 1],
            nums[nums.size() - 1] * nums[nums.size() - 2] * nums[nums.size() - 3]);
    }
};
```
