## [461. Hamming Distance](https://leetcode.com/problems/hamming-distance/#/description)

### Description

The [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance) between two integers is the number of positions at which the corresponding bits are different.

Given two integers `x` and `y`, calculate the Hamming distance.

**Note:**
0 ≤ `x`, `y` < 231.

**Example:**

```
Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.
```

**Difficulty:** `Easy`

**Tags:** `Bit Manipulation`

### Solution One - In Top Solutions

```c++
class Solution {
public:
    int hammingDistance(int x, int y) {
        unsigned val = x ^ y;
        int num = 0;
        while (val != 0)
        {
            ++num;
            val &= val - 1;
        }
        return num;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int hammingDistance(int x, int y) {
        int a1,a2,b1,b2;
        int c = 0;

        if( x > y)
        {
            a1=x;
            b1=y;
        }
        else
        {
            a1=y;
            b1=x;
        }
        while(a1)
        {
            a2=a1%2;
            b2=b1%2;
            if(a2!=b2)
            {
                c++;
            }
            a1/=2;
            b1/=2;
        }
        if(a1!=b1)
        {
            c++;
        }

        return c;

    }
};
```
