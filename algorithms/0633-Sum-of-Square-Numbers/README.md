## [633. Sum of Square numbers](https://leetcode.com/problems/sum-of-square-numbers/#/description)

### Description

Given a non-negative integer `c`, your task is to decide whether there're two integers `a` and `b` such that a2 + b2 = c.

**Example 1:**

```
Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5

```

**Example 2:**

```
Input: 3
Output: False
```

**Difficulty:** `Easy`

**Tags:** `Math`

### Solution One

`Two Pointers`

```c++
class Solution {
public:
    bool judgeSquareSum(int c) {
        int a = 0;
        int b = sqrt(c);
        while (a <= b)
        {
            int sum = a * a + b * b;
            if (sum == c)
            {
                return true;
            }
            else if (sum < c)
            {
                a++;
            }
            else
            {
                b--;
            }
        }
        return false;
    }
};
```

### Other Solutions

[633. Sum of Square Numbers - S](https://leetcode.com/problems/sum-of-square-numbers/#/solution)
