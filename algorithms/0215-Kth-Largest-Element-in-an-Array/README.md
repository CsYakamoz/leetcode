## [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/description/)

### Description

Find the **k**th largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

For example,
Given `[3,2,1,5,6,4]` and k = 2, return 5.

**Note: **
You may assume k is always valid, 1 ? k ? array's length.

**Difficulty:** `Easy`

**Tags:** `Heap` `Divide and Conquer`

### Solution One

```c++
class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        sort(nums.rbegin(), nums.rend());
        return nums[k - 1];
    }
};
```

### Solution Two

[Quickselect](https://en.wikipedia.org/wiki/Quickselect)

```c++
class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        return quickSelect(nums, 0, nums.size() - 1, nums.size() - k);
    }

    int quickSelect(vector<int> &nums, int left, int right, int k)
    {
        int pivot = nums[right];
        int pivotPos = right;
        for (int i = right - 1; i >= left; i--)
        {
            if (nums[i] > pivot)
            {
                pivotPos--;
                if (pivotPos != i)
                {
                    swap(nums[i], nums[pivotPos]);
                }
            }
        }
        swap(nums[pivotPos], nums[right]);

        if (pivotPos == k)
        {
            return nums[k];
        }
        else if (pivotPos > k)
        {
            return quickSelect(nums, left, pivotPos - 1, k);
        }
        else
        {
            return quickSelect(nums, pivotPos + 1, right, k);
        }
    }
};
```
