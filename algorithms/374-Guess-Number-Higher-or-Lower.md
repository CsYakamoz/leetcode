## [374. Guess Number Higher or Lower](https://leetcode.com/problems/guess-number-higher-or-lower/description/)

### Description

We are playing the Guess Game. The game is as follows:

I pick a number from **1** to **n**. You have to guess which number I picked.

Every time you guess wrong, I'll tell you whether the number is higher or lower.

You call a pre-defined API `guess(int num)` which returns 3 possible results (`-1`, `1`, or `0`):

```
-1 : My number is lower
 1 : My number is higher
 0 : Congrats! You got it!

```

**Example:**

```
n = 10, I pick 6.

Return 6.
```

**Difficult:** `Easy`

**Tags:** `Binary Search`

### Solution One

```c++
// Forward declaration of guess API.
// @param num, your guess
// @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
int guess(int num);

class Solution {
public:
    int guessNumber(int n) {
        int left = 1;
        int right = n;
        int val;
        while (left <= right)
        {
            int mid = (right - left) / 2 + left;
            int res = guess(mid);
            if (res == 0)
            {
                val = mid;
                break;
            }
            else if (res == -1)
            {
                right = mid - 1;
            }
            else
            {
                left = mid + 1;
            }
        }
        return val;
    }
};
```

### Other Solutions

[374. Guess Number Higher or Lower - Solution](https://leetcode.com/problems/guess-number-higher-or-lower/solution/#approach-3-ternary-search-accepted)

这里介绍了 **Ternary Search**，同时也说明了原因——为什么 **Ternary Search** 不比 **Binary Serach** 更广泛的使用
