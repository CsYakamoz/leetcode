## [442. Find All Duplicates in an Array](https://leetcode.com/problems/find-all-duplicates-in-an-array/#/description)

### Description

Given an array of integers, 1 ≤ a[i] ≤ _n_ (_n_ = size of array), some elements appear **twice** and others appear **once**.

Find all the elements that appear **twice** in this array.

Could you do it without extra space and in O(_n_) runtime?

**Example:**

```
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
```

**Difficult:** `Medium`

**Tags:** `Array`

### Solution One

`Hash Table`

```c++
class Solution {
public:
    vector<int> findDuplicates(vector<int>& nums) {
        map<int, int> m;
        vector<int> result;
        for (size_t i = 0; i < nums.size(); i++)
        {
            ++m[nums[i]];
            if (m[nums[i]] == 2)
            {
                result.push_back(nums[i]);
            }
        }
        return result;
    }
};
```

### Solution Two

`Sort`

```c++
class Solution {
public:
    vector<int> findDuplicates(vector<int>& nums) {
        vector<int> result;
        sort(nums.begin(), nums.end());
        for (size_t i = 1; i < nums.size(); i++)
        {
            if (nums[i] == nums[i-1])
            {
                result.push_back(nums[i++]);
            }
        }
        return result;
    }
};
```

### Solution Three - In Top Solutions

由于元素值范围为 [1, n]\(n = size of array)，故用 `abs[nums[i]] - 1` 做下标

检查`nums[index]`是否为负

若为负，则代表，这是第二次出现，即需要添加到`result`

若为正，则代表，这是第一次出现，使其为负

```c++
class Solution {
public:
    vector<int> findDuplicates(vector<int>& nums) {
        vector<int> result;
        for (size_t i = 0; i < nums.size(); i++)
        {
            int index = abs(nums[i]) - 1;
            if (nums[index] < 0)
            {
                result.push_back(abs(nums[i]));
            }
            else
            {
                nums[index] = -nums[index];
            }
        }
        return result;
    }
};
```
