## [191. Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/#/description)

### Description

Write a function that takes an unsigned integer and returns the number of ’1' bits it has (also known as the [Hamming weight](http://en.wikipedia.org/wiki/Hamming_weight)).

For example, the 32-bit integer ’11' has binary representation `00000000000000000000000000001011`, so the function should return 3.

**Difficulty:** `Easy`

**Tags:** `Bit Manipulation`

### Solution One

```c++
class Solution {
public:
    int hammingWeight(uint32_t n) {
        int count = 0;
        while (n)
        {
            count++;
            n &= n - 1;
        }
        return count;
    }
};
```

### Solution Two

```c++
class Solution {
public:
    int hammingWeight(uint32_t n) {
        int count = 0;
        while (n)
        {
            if (n & 1)
            {
                count++;
            }
            n >>= 1;
        }
        return count;
    }
};
```
