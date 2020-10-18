## [367. Valid Perfect Square](https://leetcode.com/problems/valid-perfect-square/#/description)

### Description

Given a positive integer _num_, write a function which returns True if _num_ is a perfect square else False.

**Note:** **Do not** use any built-in library function such as `sqrt`.

**Example 1:**

```
Input: 16
Returns: True

```

**Example 2:**

```
Input: 14
Returns: False
```

**Difficulty:** `Easy`

**Tags:** `Binary Search` `Math`

### Solution One

一开始对于变量`left`, `right`, `mid`, `val`的类型都是 int，但是这样会不通过

因为若`mid`很大时，`mid * mid`会溢出，导致结果为负

故后面改为使用`long long`类型

**Update:** 如果不想使用`int`类型，第 10 行可以改为`if (mid = num / mid)`

```c++
class Solution {
public:
    bool isPerfectSquare(int num) {
        long long left = 0;
        long long right = num;
        while (left <= right)
        {
            long long mid = (right - left) / 2 + left;
            long long val = mid * mid;
            if (val == num)
            {
                return true;
            }
            else if (val < num)
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
