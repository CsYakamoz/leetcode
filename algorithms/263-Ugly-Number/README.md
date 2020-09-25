## [263. Ugly Number](https://leetcode.com/problems/ugly-number/description/)

### Description

Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include `2, 3, 5`. For example, `6, 8` are ugly while `14` is not ugly since it includes another prime factor `7`.

Note that `1` is typically treated as an ugly number.

**Difficult:** `Easy`

**Tags:** `Math`

### Solution One

如果一个数是`Ugly Number`，则它可以表示成`num = 2x * 3y * 5z`

第一个循环使其变成`num / 2x = 3y * 5z`

第二个循环使其变成`num / 2x / 3y = 5z`

第三个循环使其变成`num / 2x / 3y / 5z = 1`

故最后只要判断其是否为 1，即可

```c++
class Solution {
public:
    bool isUgly(int num) {
        if (num < 1)
        {
            return false;
        }
        while (num % 2 == 0)
        {
            num /= 2;
        }
        while (num % 3 == 0)
        {
            num /= 3;
        }
        while (num % 5 == 0)
        {
            num /= 5;
        }
        return num == 1;
    }
};
```
