## [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/description/)

### Description

Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

**Difficulty:** `Easy`

**Tags:** `Array` `Hash Table`

### Solution One

```c++
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        set<int> s;
        for (auto i : nums)
        {
            if (s.find(i) != s.end())
            {
                return true;
            }
            else
            {
                s.insert(i);
            }
        }
        return false;
    }
};
```

### Solution Two - In Top Solutions

第一层循环确定`nums`中元素的范围

第二层循环，首先将该位置`nums[i]-min]`置为 `true`，若再次出现，即表示出现重复的元素

```c++
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        int min = INT_MAX;
        int max = INT_MIN;

        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] > max) {
                max = nums[i];
            }

            if (nums[i] < min) {
                min = nums[i];
            }
        }

        vector<bool> exists(max - min + 1, false);

        for (int i = 0; i < nums.size(); ++i) {
            if (exists[nums[i] - min]) {
                return true;
            }
            else {
                exists[nums[i] - min] = true;
            }
        }

        return false;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        for (int i = 1; i < nums.size(); i++) {
            if (nums[i] == nums[i-1]) {
                return true;
            }
        }

        return false;
    }
};
```
