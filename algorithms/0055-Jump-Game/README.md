## [55. Jump Game](https://leetcode.com/problems/jump-game/#/description)

### Description

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

For example:
A = `[2,3,1,1,4]`, return `true`.

A = `[3,2,1,0,4]`, return `false`.

**Difficulty:** `Medium`

**Tags:** `Array` `Greedy`

### Solution One

对于第 i 个元素，若能在`(i, i + nums[i] ]`区间（左开右闭）中找到一个能跳出该区间的元素

则以此元素为第 i 个元素，重新执行以上操作

否则，不能跳到结尾

```c++
class Solution {
public:
    bool canJump(vector<int>& nums) {
        int i = 0;
        int length = nums.size();
        // if nums = {0, ... }
        if (nums[0] == 0 && length != 1)
        {
            return false;
        }
        while (i < length)
        {
            int jumpCount = nums[i];
            while (++i < length && --jumpCount != -1)
            {
                if (jumpCount < nums[i])
                {
                    break;
                }
            }
            if (jumpCount == -1 && i != length)
            {
                return false;
            }
        }
        return true;
    }
};
```
