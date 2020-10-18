## [7. Reverse Integer](https://leetcode.com/problems/reverse-integer/#/description)

### Description

Reverse digits of an integer.

**Example1:** x = 123, return 321
**Example2:** x = -123, return -321

**Difficulty:** `Easy`

**Tags:** `Math`

### Solution One

```c++
class Solution {
public:
    int reverse(int x) {
        long long result = 0;		// long long有足够的位数存储溢出的int
        int sign = x > 0 ? 1 : -1;	// 存储符号位
        if (sign == -1)
        {
            x *= sign;				// 若x为负，先使其为正
        }
        while (x != 0)
        {
            result = result * 10 + x % 10;
            if (result > INT_MAX || result < INT_MIN)
            {
                return 0;		// 如果小于int的最小值或者大于int的最大值，则表明已溢出
                                  // 更改，这里由于result恒为正值，故不可能小于 INT_MIN
                                  // 个人认为通过的原因是因为，如果大于INT_MAX
                                  // 则恢复符号位后也会小于 INT_MIN
            }
            x /= 10;
        }
        result *= sign;			// 恢复符号位
        return result;
    }
};
```

### Solution Two - In Top Solutions

在自己的思路中，不知道负数除以 10 后所得余数也为负数，所以会考虑符号问题

这里则忽略了符号和在循环过程中是否溢出，只在返回时判断是否是否溢出

```c++
class Solution {
public:
    int reverse(int x) {
        long long res = 0;
        while(x) {
            res = res*10 + x%10;
            x /= 10;
        }
        return (res<INT_MIN || res>INT_MAX) ? 0 : res;
    }
};
```

### Solution Three - In Top Solutions

temp 溢出后，再除以 10 不可能会和之前未溢出一样，故这里可以用来是否相等来判断是否溢出

```c++
class Solution {
public:
    int reverse(int x) {
        int ans = 0;
        while (x) {
            int temp = ans * 10 + x % 10;
            if (temp / 10 != ans)
                return 0;
            ans = temp;
            x /= 10;
        }
        return ans;
    }
};
```
