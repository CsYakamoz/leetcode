## [278. First Bad Version](https://leetcode.com/problems/first-bad-version/#/description)

### Description

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API `bool isBadVersion(version)` which will return whether `version` is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.



**Difficult:** `Easy`

**Tags:** `Binary Search`



### Solution One

`Binary Search`

**Warning:** 对于第 11 行，一开始我是用`mid = (left + right) / 2`，但测试不通过

根据不通过的 Testcase，发现当`left`和`right`都很大时，`left + right`会溢出，所以改为`mid = (right - left) / 2 + left`

```c++
// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int left = 1;
        int right = n;
        while (left < right)
        {
            int mid = (right - left) / 2 + left;
            if (isBadVersion(mid))
            {
                right = mid;
            }
            else
            {
                left = mid + 1;
            }
        }
        return right;
    }
};
```



