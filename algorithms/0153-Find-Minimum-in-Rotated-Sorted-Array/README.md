## [153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/)

### Description

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., `0 1 2 4 5 6 7` might become `4 5 6 7 0 1 2`).

Find the minimum element.

You may assume no duplicate exists in the array.

**Difficulty:** `Medium`

**Tags:** `Array` `Binary Search`

### Solution One

思路来源于 [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/description/) 中的 Top Solutions [Concise O(log N) Binary search solution](https://discuss.leetcode.com/topic/3538/concise-o-log-n-binary-search-solution)

如果 `nums[mid] < nums[right]` 成立，则意味着 [mid, right] 都是有序的，最小值在 [left, mid] 中

否则，意味着最小值在 [mid + 1, right] 中

```c++
class Solution {
public:
    int findMin(vector<int>& nums) {
        int left = 0;
        int right = nums.size() - 1;
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            if (nums[mid] < nums[right])
            {
                right = mid;
            }
            else
            {
                left = mid + 1;
            }
        }
        return nums[right];
    }
};
```
