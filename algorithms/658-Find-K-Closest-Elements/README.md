## [658. Find K Closest Elements](https://leetcode.com/problems/find-k-closest-elements/description/)

### Description

Given a sorted array, two integers `k` and `x`, find the `k` closest elements to `x` in the array. The result should also be sorted in ascending order. If there is a tie, the smaller elements are always preferred.

**Example 1:**

```
Input: [1,2,3,4,5], k=4, x=3
Output: [1,2,3,4]
```

**Example 2:**

```
Input: [1,2,3,4,5], k=4, x=-1
Output: [1,2,3,4]
```

**Note:**

1. The value k is positive and will always be smaller than the length of the sorted array.
2. Length of the given array is positive and will not exceed 104
3. Absolute value of elements in the array and x will not exceed 104

**Difficult:** `Medium`

**Tags:** `Binary Search`

### Solution One

```c++
class Solution {
public:
    vector<int> findClosestElements(vector<int>& arr, int k, int x) {
        if (k == arr.size()) return arr;

        int length = arr.size();
        int left = 0;
        int right = length;
        // Find the first element that is greater than or equal to x
        while (left < right)
        {
            int mid = (right - left) / 2 + left;
            if (arr[mid] >= x)
            {
                right = mid;
            }
            else
            {
                left = mid + 1;
            }
        }
        // Right points to first element that is greater than or equal to x
        // If it does not exist, right is length
        deque<int> res;
        int i = right - 1;
        int j = right;
        while (k && j < length && arr[j] == x)
        {
            res.push_back(arr[j++]);
            k--;
        }
        while (k--)
        {
            int diff1 = INT_MAX;
            int diff2 = INT_MAX;
            if (i >= 0)
            {
                diff1 = x - arr[i];
            }
            if (j < length)
            {
                diff2 = arr[j] - x;
            }
            if (diff1 <= diff2)
            {
                res.push_front(arr[i--]);
            }
            else
            {
                res.push_back(arr[j++]);
            }
        }
        return vector<int>(res.begin(), res.end());
    }
};
```
