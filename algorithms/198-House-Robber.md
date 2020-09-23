## [198. House Robber](https://leetcode.com/problems/house-robber/description/)

### Description

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and **it will automatically contact the police if two adjacent houses were broken into on the same night**.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight **without alerting the police**.

**Difficult:** `Easy`

**Tags:** `Dynamic Programming`

### Solution One - In Top Solutions

[Java O(n) solution, space O(1)](https://discuss.leetcode.com/topic/11082/java-o-n-solution-space-o-1)

```c++
class Solution {
public:
    int rob(vector<int>& nums)
    {
        int y = 0;
        int n = 0;
        for (auto money : nums) {
            int tmp = n;
            n = max(y, n);
            y = tmp + money;
        }
        return max(y, n);
    }
};
```
