## [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/description/)

### Description

Given two sorted integer arrays _nums1_ and _nums2_, merge _nums2_ into _nums1_ as one sorted array.

**Note:**
You may assume that _nums1_ has enough space (size that is greater or equal to _m_ + _n_) to hold additional elements from _nums2_. The number of elements initialized in _nums1_ and _nums2_ are _m_ and _n_ respectively.

**Difficult:** `Easy`

**Tags:** `Array` `Two Pointers`

### Solution One - In Top Solutions

[This is my AC code, may help you](https://discuss.leetcode.com/topic/2461/this-is-my-ac-code-may-help-you)

```c++
class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int i = m - 1;
        int j = n - 1;
        int back = m + n - 1;
        while (i >= 0 && j >= 0)
        {
            if (nums1[i] > nums2[j])
            {
                nums1[back--] = nums1[i--];
            }
            else
            {
                nums1[back--] = nums2[j--];
            }
        }
        while (j >= 0)
        {
            nums1[back--] = nums2[j--];
        }
    }
};
```
