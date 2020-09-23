## [665. Non-decreasing Array](https://leetcode.com/problems/non-decreasing-array/description/)

### Description

Given an array with `n` integers, your task is to check if it could become non-decreasing by modifying **at most** `1` element.

We define an array is non-decreasing if `array[i] <= array[i + 1]` holds for every `i` (1 <= i < n).

**Example 1:**

```
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.

```

**Example 2:**

```
Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.

```

**Note:** The `n` belongs to [1, 10,000].

**Difficult:** `Easy`

**Tags:** `Array`

### Solution One

```c++
class Solution {
public:
    bool checkPossibility(vector<int>& nums) {
        if (nums.size() < 2)
            return true;

        bool changed = false;
        if (nums[0] > nums[1])
        {
            nums[0] = nums[1] - 1;
            changed = true;
        }

        for (size_t i = 2; i < nums.size(); i++)
        {
            if (nums[i - 1] > nums[i])
            {
                if (changed == false)
                {
                    if (nums[i - 1] == nums[i - 2])
                        nums[i] = nums[i - 1];
                    else
                    {
                        if (nums[i] >= nums[i - 2])
                            nums[i - 1] = nums[i - 2];
                        else
                            nums[i] = nums[i - 1];
                    }
                    changed = true;
                }
                else
                    return false;
            }
        }

        return true;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    bool checkPossibility(vector<int>& nums) {
        int n = nums.size();
        int count = 0;
        for (int i = 0; i < n-1; i++) {
            if (nums[i] > nums[i+1]) {
                if (count++ > 0) {
                    return false;
                }
                if( i==0 || nums[i-1]<=nums[i+1]) {
                    nums[i] = nums[i+1];
                } else {
                    // 2    3    1    6   7
                    //      i    i+1
                    nums[i+1] = nums[i];
                }
            }
        }
        return true;
    }
};
```
