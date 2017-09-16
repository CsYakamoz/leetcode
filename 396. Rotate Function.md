## [396. Rotate Function](https://leetcode.com/problems/rotate-function/description/)

### Description

Given an array of integers `A` and let *n* to be its length.

Assume `Bk` to be an array obtained by rotating the array `A` *k* positions clock-wise, we define a "rotation function" `F` on `A` as follow:

`F(k) = 0 * Bk[0] + 1 * Bk[1] + ... + (n-1) * Bk[n-1]`.

Calculate the maximum value of `F(0), F(1), ..., F(n-1)`.

**Note:**
*n* is guaranteed to be less than 10^5.

**Example:**

```
A = [4, 3, 2, 6]

F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26

So the maximum value of F(0), F(1), F(2), F(3) is F(3) = 26.
```



**Difficult:** `Medium`

**Tags:** `Math`



### Solution One

> **15 / 17** test cases passed.
>
> Status: **Time Limit Exceeded**

```c++
class Solution {
public:
    int maxRotateFunction(vector<int>& A) {
        if (A.empty())
            return 0;

        int result = INT_MIN;
        for (size_t i = 0; i < A.size(); i++)
        {
            int start = (A.size() - i) % A.size();
            int count = 0;
            int val = 0;
            int j = start;
            do
            {
                val += count++ * A[j];
                j = (j + 1) % A.size();
            } while (j != start);
            result = max(result, val);
        }
        return result;
    }
};
```



### Solution Two

```c++
class Solution {
public:
    int maxRotateFunction(vector<int>& A) {
        if (A.size() < 2)
            return 0;

        int length = A.size();
        int F = 0;
        int sum = 0;

        for (int i = 0; i < length; i++)
        {
            F += i * A[i];  // F(0)
            sum += A[i];
        }

        int res = F;
        for (int i = 1; i < length; i++)
        {   // F(i)
            F += length * A[i - 1] - sum;
            res = max(res, F);
        }
        return res;
    }
};
```



