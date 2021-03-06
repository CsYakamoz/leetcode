## [50. Pow(x, n)](https://leetcode.com/problems/powx-n/#/description)

### Description

Implement pow(x, n)

**Difficulty:** `Medium`

**Tags:** `Binary Search` `Math`

### Solution One - In Top Solutions

对于一个正整数，如 9，我们可以写成：1001，即$2^3+2^0$

那么将有 $x^9 = x^{2^3} + x^{2^0}$

那么对于任意 x，n

则有：

$$
x^n = x^{2^i} + x^{2^{i-1}} + \cdots + x^{2^0}
$$

为了防止 n 为 INT_MIN，故`absN`的类型为 long long，因为用 int 类型存储 -INT_MIN 会溢出

用此思路，时间复杂度为: $T(log_2 N)$

```c++
class Solution {
public:
    double myPow(double x, int n) {
        double result = 1;
        if (n < 0)
        {
            x = 1 / x;
        }
        long long absN = llabs((long long)n);
        while (absN > 0)
        {
            if (absN & 1)
            {
                result *= x;
            }
            absN >>= 1;
            x *= x;
        }
        return result;
    }
};
```

### Solution Two - In Top Solutions

$x^9 = x * (x^2)^4 = x * (x^4)^2 = x * (x^8)^1$

$x^8 = (x^2)^4 = (x^4)^2 = (x^8)^1$

此处使用(long long)n 的理由同上

```c++
class Solution {
public:
    double myPow(double x, int n) {
        if (n == 0)
        {
            return 1;
        }
        else if (n < 0)
        {
            return doPow(1 / x, -(long long)n);
        }
        else
        {
            return doPow(x, n);
        }
    }
private:
    double doPow(double x, long long n)
    {
        if (n == 1)
        {
            return x;
        }
        return n % 2 == 0 ? doPow(x*x, n / 2) : x*doPow(x*x, n / 2);
    }
};
```
