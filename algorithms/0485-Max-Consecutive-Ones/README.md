## [485. Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/#/description)

### Description

Given a binary array, find the maximum number of consecutive 1s in this array.

**Example 1:**

```
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.

```

**Note:**

- The input array will only contain `0` and `1`.
- The length of input array is a positive integer and will not exceed 10,000

**Difficulty:** `Easy`

**Tags:** `Array`

### Solution One

```c++
class Solution {
public:
    int findMaxConsecutiveOnes(vector<int>& nums) {
        size_t i = 0;
        int res = 0;
        while (i < nums.size()) {
            int count = 0;
            while (i < nums.size() && nums[i]) {
                count++;
                i++;
            }
            res = count > res ? count : res;
            while (i < nums.size() && !nums[i]) {
                i++;	// Skip zeros
            }
        }
        return res;
    }
};
```

### Solution Two

```c++
class Solution {
public:
    int findMaxConsecutiveOnes(vector<int>& nums) {
        size_t i = 0;
        int res = 0;
        while (i < nums.size()) {
            if (i + res >= nums.size()) {
                break;
            }
            int count = 0;
            while (i < nums.size() && nums[i]) {
                count++;
                i++;
            }
            res = count > res ? count : res;

            while (i < nums.size() && !nums[i]) {
                i++;	// Skip zeros
            }
        }
        return res;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    int findMaxConsecutiveOnes(vector<int>& nums) {
        int cnt = 0, ans = 0;
        for(int i = 0; i < nums.size(); ++ i) {
            if(nums[i])
                ++cnt;
            else {
                ans = max(ans, cnt);
                cnt = 0;
            }
        }
        ans = max(ans, cnt);
        return ans;
    }
};

```
