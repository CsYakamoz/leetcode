## [136. Single Number](https://leetcode.com/problems/single-number/#/description)

### Description

Given an array of integers, every element appears _twice_ except for one. Find that single one.

**Note:**
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

**Difficulty:** `Easy`

**Tags:** `Hash Table` `Bit Manipulation`

### Solution One

```c++
class Solution {
public:
    int singleNumber(vector<int>& nums) {
        unordered_map<int, int> nums_count;
        int result = 0;
        for (size_t i = 0; i < nums.size(); i++)
        {
            ++nums_count[nums[i]];
        }
        auto iter = nums_count.cbegin();
        while (iter != nums_count.cend())
        {
            if (iter->second == 1)
            {
                result = iter->first;
                break;
            }
            ++iter;
        }
        return result;
    }
};
```

### Solution Two - In Top Solutions

两个相同的数进行异或操作后，值为 0

即 `A XOR A = 0`

```c++
class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int result = 0;
        for (const auto &i : nums)
            result ^= i;
        return result;
    }
};
```
