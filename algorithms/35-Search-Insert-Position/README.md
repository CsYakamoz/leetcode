## [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/#/description)

### Description

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Here are few examples.
`[1,3,5,6]`, 5 → 2
`[1,3,5,6]`, 2 → 1
`[1,3,5,6]`, 7 → 4
`[1,3,5,6]`, 0 → 0

**Difficult:** `Easy`

**Tags:** `Array` `Binary Search`

### Solution One

在数组中查找第一个大于等于`target`的元素

若找到，返回该元素的索引，否则，返回数组长度

```c++
class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        for (size_t i = 0; i < nums.size(); i++)
        {
            if (target <= nums[i])
            {
                return i;
            }
        }
        return nums.size();
    }
};
```

### Solution Two

`Binary Search`

使用二分搜索法

```c++
class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int front = 0;
        int back = nums.size();
        while (front < back)
        {
            int mid = (front + back) / 2;
            if (nums[mid] == target)
            {
                return mid;
            }
            else if (nums[mid] < target)
            {
                front = mid + 1;
            }
            else
            {
                back = mid;
            }
        }
        return front;
    }
};
```
