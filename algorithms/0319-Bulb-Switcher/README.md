## [319. Bulb Switcher](https://leetcode.com/problems/bulb-switcher/description/)

### Description

There are _n_ bulbs that are initially off. You first turn on all the bulbs. Then, you turn off every second bulb. On the third round, you toggle every third bulb (turning on if it's off or turning off if it's on). For the *i*th round, you toggle every _i_ bulb. For the *n*th round, you only toggle the last bulb. Find how many bulbs are on after _n_ rounds.

**Example:**

```
Given n = 3.

At first, the three bulbs are [off, off, off].
After first round, the three bulbs are [on, on, on].
After second round, the three bulbs are [on, off, on].
After third round, the three bulbs are [on, off, off].

So you should return 1, because there is only one bulb is on.
```

**Difficulty:** `Medium`

**Tags:** `Math` `Brainteaser`

### Solution One

找规律

```c++
class Solution {
public:
    int bulbSwitch(int n) {
        if (n == 0) return 0;

        int i = 3;
        int length = 3;
        while (n > i)
        {
            i += length + 2;
            length += 2;
        }

        return length / 2;
    }
};
```

### Solution Two - In Top Solutions

[Share my o(1) solution with explanation](https://discuss.leetcode.com/topic/39558/share-my-o-1-solution-with-explanation)

第 i 个灯泡，只会在第 k 个回合时改变状态（k 为 i 的因子）

例如：第 6 个灯泡，那么其状态只会在第 1，2，3，6 个回合被改变

如果因子数为奇数，则该灯泡最后为 ON，否则为 OFF

故问题转化为：**1 ~ N 中，有多少个数的因子数为奇数？**

那么在什么情况下，一个数的因子数为奇数？

这里分素数和非素数两种情况讨论：

素数情况下：因子只有两个：1 和 本身，因子数为奇数个

非素数情况下：1 和 本身，还有成对的 i / p 和 p (p >= 2)

若不存在 i / p == p，则因子数为偶数个

若存在，则因子数为奇数个，此时 i 为平方数

故问题又转化为：**1 ~ N 中有多少个平方数？**

```c++
class Solution {
public:
    int bulbSwitch(int n) {
        return sqrt(n);
    }
};
```

> As we know that there are n bulbs, let's name them as 1, 2, 3, 4, ..., n.
>
> 1. You first turn on all the bulbs.
> 2. Then, you turn off every second bulb.(2, 4, 6, ...)
> 3. On the third round, you toggle every third bulb.(3, 6, 9, ...)
> 4. For the ith round, you toggle every i bulb.(i, 2i, 3i, ...)
> 5. For the nth round, you only toggle the last bulb.(n)
>
> ---
>
> If n > 6, you can find that bulb 6 is toggled in round 2 and 3.
>
> Later, it will also be toggled in round 6, and round 6 will be the last round when bulb 6 is toggled.
>
> Here, **2,3 and 6 are all factors of 6 (except 1).**
>
> ---
>
> ## **Prove:**
>
> We can come to the conclusion that **the bulb i is toggled k times.**
>
> Here, **k** is **the number of i's factors (except 1)**.
>
> **k + 1** will be **the total number of i's factors**
>
> ---
>
> For example:
>
> - **Factors of 6: 1, 2, 3, 6 (3 factors except 1, so it will be toggled 3 times)**
> - **Factors of 7: 1, 7 (1 factors except 1, so it will be toggled once)**
>   ....
>
> Now, the key problem here is to judge **whether k is even or odd.**
>
> ---
>
> Since **all bulbs are on at the beginning**, we can get:
>
> - **If k is odd, the bulb will be off in the end.(after odd times of toggling).**
> - **If k is even, the bulb i will be on in the end (after even times of toggling).**
>
> As we all know, **a natural number can divided by 1 and itself**, and **all factors appear in pairs**.
>
> **When we know that p is i's factor, we are sure q = i/p is also i's factor.**
>
> **If i has no factor p that makes p = i/p, k+ 1 is even.**
>
> **If i has a factor p that makes p = i/p (i = p^2, i is a perfect square of p), k+ 1 is odd.**
>
> ---
>
> So we get that **in the end**:
>
> - If **i** is a **perfect square** , _k_+ 1 is odd, **k is even (bulb i is on)**.
> - If **i** is **NOT** a **perfect square** , _k_+ 1 is even, **k is odd (bulb i is off)**.
>
> ---
>
> We want to find **how many bulbs are on** after _n_ rounds (**In the end**).
>
> That means we need to find out **how many perfect square numbers are NO MORE than n**.
>
> The **number of perfect square numbers which are no more than n**, is the **square root of the maximum perfect square number which is NO MORE than n**
>
> ---
>
> ## **Result:**
>
> The **square root of the maximum perfect square number which is NO MORE than n** is the
> **integer part of sqrt(n).**
>
> (**If i = 1, it will NEVER be toggled, k is 0 (even) here which meets the requirement.**)
