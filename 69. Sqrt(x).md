## [69. Sqrt(x)](https://leetcode.com/problems/sqrtx/#/description)

### Description

Implement `int sqrt(int x)`.

Compute and return the square root of *x*.



**Difficult:** `Easy`

**Tags:** `Binary Search` `Math`



### Solution One

`Binary Search`

```c++
class Solution {
public:
    int mySqrt(int x) {
        long long left = 0;
        long long right = x;
        while (left <= right)
        {
            long long mid = (right - left) / 2 + left;
            long long val = mid * mid;
            if (val == x)
            {
                return mid;
            }
            else if (val < x)
            {
                left = mid + 1;
            }
            else
            {
                right = mid - 1;
            }
        }
        return right;
    }
};
```



### Solution Two - In Top Solutions

[Share my O(log n) Solution using bit manipulation](https://discuss.leetcode.com/topic/2671/share-my-o-log-n-solution-using-bit-manipulation)

[3-4 short lines, Integer Newton, Every Language](https://discuss.leetcode.com/topic/24532/3-4-short-lines-integer-newton-every-language)