## [650. 2 Keys Keyboard](https://leetcode.com/problems/2-keys-keyboard/description/)

### Description

Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:

1. `Copy All`: You can copy all the characters present on the notepad (partial copy is not allowed).
2. `Paste`: You can paste the characters which are copied **last time**.

Given a number `n`. You have to get **exactly** `n` 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get `n` 'A'.

**Example 1:**

```
Input: 3
Output: 3
Explanation:
Intitally, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.

```

**Note:**

1. The `n` will be in the range [1, 1000].

**Difficult:** `Medium`

**Tags:** `Dynamic Programming`

### Solution One

分三种情况讨论：

1. n 为偶数：

   只要在有 n / 2 个 `A` 时，执行操作 `Copy All`、`Paste`，即可得到 n 个`A`

2. n 为奇数且不是素数

   因为不是素数，所以可以找出其除了自身之外的最大因子（假设为 factor），在有 factor 个`A` 时，首先执行 `Copy All` 操作，然后执行 `n / factor - 1` 次 `Paste` 操作，即可得到 n 个`A`

3. n 为奇数且是素数

   因为是素数，所以只能一个一个来复制，只要执行 `Copy All`，和 `n - 1` 次 `Paste` 操作，即可得到 n 个 `A`

**Update:** 因为 `n` 的范围为 [1, 1000]， 所以可以直接弄个素数数组 `prime`（长度为 168，而 log 168 约等于 7），当 n 为奇数时，在 `prime` 中进行二分查找，即可判断是否为素数

```c++
class Solution {
public:
    int minSteps(int n) {
        if (n == 1) return 0;
        if (n % 2 == 0)
        {   // n is even
            return minSteps(n / 2) + 2;
        }
        else
        {   // n is odd
            int factor;
            if (isPrime(n, factor))
                // n is prime
                return n;
            else
                // n is nor prime
                return n / factor + minSteps(factor);
        }
    }

private:
    bool isPrime(int n, int &factor)
    {
        for (int i = 2; i * i <= n; i++)
            if (n % i == 0)
            {
                factor = n / i;
                return false;
            }
        return true;
    }
};
```

### Solution Two - In Top Solutions

这里直接将上面三种情况整合了在一起

**One** 中的三个 return 整合起来，就是下面的

`return i + minSteps(n / i)`

```c++
class Solution {
public:
    int minSteps(int n) {
        if (n==1) return 0;

        int half = sqrt(n)+1;
        for (int i=2; i < half; ++i) {
            if ((n%i) == 0) {
                return (i+minSteps(n/i));
            }
        }

        return n;
    }
};
```
