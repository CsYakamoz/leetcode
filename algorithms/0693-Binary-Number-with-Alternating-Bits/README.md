## [693. Binary Number with Alternating Bits](https://leetcode.com/problems/binary-number-with-alternating-bits/description/)

### Description

Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

**Example 1:**

```
Input: 5
Output: True
Explanation:
The binary representation of 5 is: 101

```

**Example 2:**

```
Input: 7
Output: False
Explanation:
The binary representation of 7 is: 111.

```

**Example 3:**

```
Input: 11
Output: False
Explanation:
The binary representation of 11 is: 1011.

```

**Example 4:**

```
Input: 10
Output: True
Explanation:
The binary representation of 10 is: 1010.
```

**Difficulty:** `Easy`

**Tags:** `Bit Manipulation`

### Solution One

```c++
class Solution {
public:
    bool hasAlternatingBits(int n)
    {
        string str;
        do {
            str.push_back(n % 2 + '0');
            n /= 2;
        } while (n != 0);

        for (size_t i = 1; i < str.size(); i++)
            if (str[i - 1] == str[i]) return false;

        return true;
    }
};
```

### Solution Two - In Top Solutions

[Binary Number with Alternating Bits - Discuss](https://discuss.leetcode.com/category/1550/binary-number-with-alternating-bits)
