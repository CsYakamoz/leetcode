## [441. Arranging Coins](https://leetcode.com/problems/arranging-coins/description/)

### Description

You have a total of _n_ coins that you want to form in a staircase shape, where every _k_-th row must have exactly _k_ coins.

Given _n_, find the total number of **full** staircase rows that can be formed.

_n_ is a non-negative integer and fits within the range of a 32-bit signed integer.

**Example 1:**

```
n = 5

The coins can form the following rows:
¤
¤ ¤
¤ ¤

Because the 3rd row is incomplete, we return 2.

```

**Example 2:**

```
n = 8

The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

Because the 4th row is incomplete, we return 3.
```

**Difficult:** `Easy`

**Tags:** `Binary Search` `Math`

### Solution One

对于公式 `(1 + n) * n / 2`，当 n 达到一定程度时，其和会溢出，导致 WA

> Status: Wrong Answer
>
> **13 / 1336** test cases passed.
>
> Input: 1804289383
>
> Output: 1804289383
>
> Expected: 60070

```c++
class Solution {
public:
    int arrangeCoins(int n) {
        int left = 1;
        int right = n;
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            // 1 + ... + n = (1 + n) * n / 2
            int total = (1 + mid) * mid / 2;
            if (total == n)
            {
                right = mid;
                break;
            }
            else if (total < n)
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

### Solution Two

尝试将 total 类型改为 `long long` 或者 `long double`，但依然会溢出，于是想办法从公式本身出发修改

```c++
class Solution {
public:
    int arrangeCoins(int n) {
        // sum = 1 + ... + n = (1 + n) * n / 2
        // sum = (n^2 + n) / 2
        // 2sum = (n^2 + n)
        // n^2 + n + (1/4) = 2sum + (1/4)
        // ( n + (1/2) )^2 = 2sum + (1/4)
        // n + (1/2) = sqrt( 2sum + (1/4) )
        // n = sqrt( 2sum + (1/4) ) - (1/2)
        double x = (double)n * 2 + 0.25;
        return sqrt(x) - 0.5;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    int arrangeCoins(int n) {
        return (int)(2*sqrt(0.0625+n/2.0)-0.5);
    }
};
```
