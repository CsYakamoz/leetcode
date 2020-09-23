## [240. Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/description/)

### Description

Write an efficient algorithm that searches for a value in an *m* x *n* matrix. This matrix has the following properties:

- Integers in each row are sorted in ascending from left to right.
- Integers in each column are sorted in ascending from top to bottom.

For example,

Consider the following matrix:

```
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

```

Given **target** = `5`, return `true`.

Given **target** = `20`, return `false`.



**Difficult:** `Medium`

**Tags:** `Binary Search` `Divide and Conquer`



### Solution One - In Top Solutions

[My concise O(m+n) Java solution](https://discuss.leetcode.com/topic/20064/my-concise-o-m-n-java-solution)

> We start search the matrix from top right corner, initialize the current position to top right corner, if the target is greater than the value in current position, then the target can not be in entire row of current position because the row is sorted, if the target is less than the value in current position, then the target can not in the entire column because the column is sorted too. We can rule out one row or one column each time, so the time complexity is O(m+n).

```c++
class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        if (matrix.empty() || matrix[0].empty())
        {
            return false;
        }
        int m = matrix.size();
        int n = matrix[0].size();
        int i = 0;
        int j = n - 1;
        while (i < m && j >= 0)
        {
            if (matrix[i][j] == target)
            {
                return true;
            }
            else if (matrix[i][j] > target) 
            {
                j--;
            }
            else
            {
                i++;
            }
        }
        return false;
    }
};
```



### Solution Two - In Top Solutions

[C++ two solutions (O(m+n), O(mlogn))](https://discuss.leetcode.com/topic/19487/c-two-solutions-o-m-n-o-mlogn)

```c++
 bool searchMatrix(vector<vector<int>>& matrix, int target) {
return searchMatrix(matrix, target, 0, matrix.size() - 1);
 }

 bool searchMatrix(vector<vector<int>>& matrix, int target, int top, int bottom) {
if (top > bottom)
	return false;

int mid = top + (bottom - top) / 2;
if (matrix[mid].front() <= target && target <= matrix[mid].back())
	if (searchVector(matrix[mid], target)) return true;

if (searchMatrix(matrix, target, top, mid - 1)) return true;
if (searchMatrix(matrix, target, mid + 1, bottom)) return true;

return false;
 }

 bool searchVector(vector<int>& v, int target) {
int left = 0, right = v.size() - 1;

while (left <= right) {
	int mid = left + (right - left) / 2;
	if (v[mid] == target)
		return true;
	if (v[mid] < target)
		left = mid + 1;
	else
		right = mid - 1;
}

return false;
 }
```



### Solution Three - In Top Solutions

[九章算法 | Google 面试题：Search a 2D Matrix II](https://zhuanlan.zhihu.com/p/29555088)

