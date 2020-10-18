## [231. Power of Two](https://leetcode.com/problems/power-of-two/#/description)

### Description

Given an integer, write a function to determine if it is a power of two.

**Difficulty:** `Easy`

**Tags:** `Math` `Bit Manipulation`

### Solution One

```c++
class Solution {
public:
    bool isPowerOfTwo(int n) {
        long long val = 1;
        while (val < n)
        {
            val *= 2;
        }
        return val == n;
    }
};
```

### Solution Two - In Top Solutions

[Using n&(n-1) trick](https://discuss.leetcode.com/topic/17857/using-n-n-1-trick)

```c++
class Solution {
public:
    bool isPowerOfTwo(int n) {
        if(n<=0) return false;
        return !(n&(n-1));
    }
};
```

### Solution Three - In Top Solutions

[4 different ways to solve -- Iterative / Recursive / Bit operation / Math](https://discuss.leetcode.com/topic/47195/4-different-ways-to-solve-iterative-recursive-bit-operation-math)
