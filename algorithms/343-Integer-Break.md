## [343. Integer Break](https://leetcode.com/problems/integer-break/description/)

### Description

Given a positive integer _n_, break it into the sum of **at least** two positive integers and maximize the product of those integers. Return the maximum product you can get.

For example, given _n_ = 2, return 1 (2 = 1 + 1); given _n_ = 10, return 36 (10 = 3 + 3 + 4).

**Note**: You may assume that _n_ is not less than 2 and not larger than 58.

**Difficult:** `Medium`

**Tags:** `Dynamic Programming`

### Solution One

```c++
class Solution {
public:
    int integerBreak(int n) {
        vector<int> dp{ 0,1 };
        for (int i = 2; i <= n; i++)
        {
            int x = 1;
            int y = i - 1;
            int val = 0;
            while (x <= y)
            {
                val = max(val, max(x, dp[x]) * max(y, dp[y]));
                x++;
                y--;
            }
            dp.push_back(val);
        }
        return dp[n];
    }
};
```

### Solution Two - In Top Solutions

[Why factor 2 or 3? The math behind this problem.](https://discuss.leetcode.com/topic/43055/why-factor-2-or-3-the-math-behind-this-problem)

> I saw many solutions were referring to factors of 2 and 3. But why these two magic numbers? Why other factors do not work?
> Let's study the math behind it.
>
> For convenience, say **n** is sufficiently large and can be broken into any smaller real positive numbers. We now try to calculate which real number generates the largest product.
> Assume we break **n** into **(n / x)** **x**'s, then the product will be **xn/x**, and we want to maximize it.
>
> Taking its derivative gives us **n \* x^n/x-2^ \* (1 - ln(x))**.
> The derivative is positive when **0 < x < e**, and equal to **0** when **x = e**, then becomes negative when **x > e**,
> which indicates that the product increases as **x** increases, then reaches its maximum when **x = e**, then starts dropping.
>
> This reveals the fact that if **n** is sufficiently large and we are allowed to break **n** into real numbers,
> the best idea is to break it into nearly all **e**'s.
> On the other hand, if **n** is sufficiently large and we can only break **n** into integers, we should choose integers that are closer to **e**.
> The only potential candidates are **2** and **3** since **2 < e < 3**, but we will generally prefer **3** to **2**. Why?
>
> Of course, one can prove it based on the formula above, but there is a more natural way shown as follows.
>
> **6 = 2 + 2 + 2 = 3 + 3**. But **2 \* 2 _ 2 < 3 _ 3**.
> Therefore, if there are three **2**'s in the decomposition, we can replace them by two **3**'s to gain a larger product.
>
> All the analysis above assumes **n** is significantly large. When **n** is small (say **n <= 10**), it may contain flaws.
> For instance, when **n = 4**, we have **2 \* 2 > 3 \* 1**.
> To fix it, we keep breaking **n** into **3**'s until **n** gets smaller than **10**, then solve the problem by brute-force.

### Solution Three - In Top Solutions

[Easy to understand C++ with explanation](https://discuss.leetcode.com/topic/43042/easy-to-understand-c-with-explanation)

> For any integer `p` strictly greater than `4`, it has the property such that `3 * (p - 3) > p`, which means breaking it into two integers `3` and `p - 3` makes the product larger while keeping the sum unchanged. If `p - 3` is still greater than `4`, we should break it again into `3` and `p - 6`, giving `3 * 3 * (p - 6)`, and so on, until we cannot break it (less than or equal to 4) anymore.
>
> For integer `4`, breaking it into `2 * 2` or keeping it as `4` does not change its contribution to the product.
> We cannot have more than two `4`s, because `2 * 3 * 3 > 4 * 4`. We cannot have more than three `2`s because `3 * 3 > 2 * 2 * 2`.

```c++
class Solution {
public:
    long long integerBreak(long long n) {
        if(n == 2) return 1;
        if(n == 3) return 2;
        if(n == 4) return 4;
        if(n == 5) return 6;
        if(n == 6) return 9;
        return 3 * integerBreak(n - 3);
    }
};
```
