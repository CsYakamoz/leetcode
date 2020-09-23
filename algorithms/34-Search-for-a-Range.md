## [34. Search for a Range](https://leetcode.com/problems/search-for-a-range/description/)

### Description

Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of _O_(log _n_).

If the target is not found in the array, return `[-1, -1]`.

For example,
Given `[5, 7, 7, 8, 8, 10]` and target value 8,
return `[3, 4]`.

**Difficult:** `Medium`

**Tags:** `Binary Search` `Array`

### Solution One

```c++
class Solution {
public:
    vector<int > searchRange(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size();
        // Search the first target
        while (left < right)
        {
            int mid = (right - left) / 2 + left;
            if (nums[mid] >= target)
            {
                right = mid;
            }
            else
            {
                left = mid + 1;
            }
        }
        if (right == nums.size() || nums[right] != target)
        {   // the target is not found in the array
            return { -1, -1 };
        }
        vector<int> range{ right };
        left = right;
        right = nums.size();
        // Search the element after the last target
        while (left < right)
        {
            int mid = (right - left) / 2 + left;
            if (nums[mid] > target)
            {
                right = mid;
            }
            else
            {
                left = mid + 1;
            }
        }
        // right points to the element after the last target
        range.push_back(right - 1);
        return range;
    }
};
```

---

```c++
class Solution {
public:
    vector<int > searchRange(vector<int>& nums, int target) {
        int left1 = 0, left2 = 0;
        int right1 = nums.size(), right2 = nums.size();
        while (left1 < right1 || left2 < right2)
        {
            if (left1 < right1)
            {   // Search the first target
                int mid1 = (right1 - left1) / 2 + left1;
                if (nums[mid1] >= target)
                    right1 = mid1;
                else
                    left1 = mid1 + 1;
            }

            if (left2 < right2)
            {   // Search the last target
                int mid2 = (right2 - left2) / 2 + left2;
                if (nums[mid2] > target)
                    right2 = mid2;
                else
                    left2 = mid2 + 1;
            }
        }
        if (right1 == nums.size() || nums[right1] != target)
        {
            return { -1,-1 };
        }
        else
        {
            return { right1, right2 - 1 };
        }
    }
};
```
