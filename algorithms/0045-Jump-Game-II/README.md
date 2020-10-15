## [45. Jump Game II](https://leetcode.com/problems/jump-game-ii/#/description)

### Description

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

For example:
Given array A = `[2,3,1,1,4]`

The minimum number of jumps to reach the last index is `2`. (Jump `1` step from index 0 to 1, then `3` steps to the last index.)

**Note:**
You can assume that you can always reach the last index.

**Difficult:** `Hard`

**Tags:** `Array` `Greedy`

### Solution One

对于第 i 个元素

若 `i + nums[i]`大于等于`nums.size() - 1`，则表示可以直接跳到末尾

若小于，则需要 `(i, i + nums[i] ]`区间（左开右闭）中找到一个元素跳得最远的（贪心）

```c++
class Solution {
public:
    int jump(vector<int>& nums) {
        int i = 0;
        int length = 0;
        while (i < nums.size() - 1)
        {
            int jumpCount = nums[i];
            // nextPos points to the next place where we jump
            int nextPos = i;

            // maxIndex points to the maximum index where nextPos can jump
            int maxIndex = i + nums[i];

            // if we can jump to the last index
            if (maxIndex >= nums.size()-1)
            {
                ++length;
                break;
            }
            while (--jumpCount != -1)
            {
                ++i;
                if (i + nums[i] > maxIndex)
                {
                    maxIndex = i + nums[i];
                    nextPos = i;
                }
            }
            i = nextPos;
            ++length;
        }
        return length;
    }
};
```
