## [74. Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/description/)

### Description

Write an efficient algorithm that searches for a value in an _m_ x _n_ matrix. This matrix has the following properties:

- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

For example,

Consider the following matrix:

```
[
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
```

Given **target** = `3`, return `true`.

**Difficult:** `Medium`

**Tags:** `Array` `Binary Search`

### Solution One

```c++
class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        if (matrix.empty() || matrix[0].empty())
        {
            return false;
        }
        int col = matrix[0].size();
        int left = 0;
        int right = matrix.size() * col - 1;
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            int val = matrix[mid / col][mid % col];
            if (val == target)
            {
                return true;
            }
            else if (val < target)
            {
                left = mid + 1;
            }
            else
            {
                right = mid - 1;
            }
        }
        return false;
    }
};
```
