## [413. Arithmetic Slices](https://leetcode.com/problems/arithmetic-slices/description/)

### Description

A sequence of number is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.

For example, these are arithmetic sequence:

```
1, 3, 5, 7, 9
7, 7, 7, 7
3, -1, -5, -9
```

The following sequence is not arithmetic.

```
1, 1, 2, 5, 7
```

A zero-indexed array A consisting of N numbers is given. A slice of that array is any pair of integers (P, Q) such that 0 <= P < Q < N.

A slice (P, Q) of array A is called arithmetic if the sequence:
A[P], A[p + 1], ..., A[Q - 1], A[Q] is arithmetic. In particular, this means that P + 1 < Q.

The function should return the number of arithmetic slices in the array A.

**Example:**

```
A = [1, 2, 3, 4]

return: 3, for 3 arithmetic slices in A: [1, 2, 3], [2, 3, 4] and [1, 2, 3, 4] itself.
```

**Difficult:** `Medium`

**Tags:** `Dynamic Programming`

### Solution One

ArithmeticSlices （以下简称 AS）

`findArithmeticSlices` 辅助函数用来寻找一个长度为 3 的 AS

假设 `A[i] ... A[j]` 是 AS，且差为 `diff`，同时能组成 `prev` 个 AS

若 `A[j+1] - A[j] == diff`，则 `A[i] ... A[j+1]` 是 AS，并且能组成 `prev + 1` 个 AS

原因：对于 `A[i] ... A[j]` ，我们从右边取 3 个元素能组成 AS，取 4、5、6 个也能组成 AS，最多能组成 `j - i - 1` 个 AS

那么对于 `A[i] ... A[j], A[j+1]`，我们在 `A[i] ... A[j]` 右边取 2 个元素加上 `A[j+1]` 能组成 AS，取 3、4、5 个加上 `A[j+1]` 也能组成 AS， 最多能组成 `j - i` 个 AS

```c++
class Solution {
public:
    int numberOfArithmeticSlices(vector<int>& A) {
        if (A.size() < 3) return 0;

        int i = findArithmeticSlices(A, 0);
        if (i == A.size()) return 0;

        int res = 1;
        int diff = A[i + 1] - A[i];
        i += 3;
        int prev = 1;

        while (i < A.size())
        {
            if (A[i] - A[i-1] == diff)
            {
                res += prev + 1;
                prev = prev + 1;
                i++;
            }
            else
            {
                i = findArithmeticSlices(A, i - 1);
                if (i == A.size()) break;

                res++;
                diff = A[i + 1] - A[i];
                i += 3;
                prev = 1;
            }
        }
        return res;
    }

private:
    int findArithmeticSlices(vector<int> &A, int start)
    {
        for (int i = start; i < A.size() - 2; i++)
        {
            if (A[i + 1] - A[i] == A[i + 2] - A[i + 1])
            {
                return i;
            }
        }
        return A.size();
    }
};
```

### Other Solutions

[413. Arithmetic Slices - Solutions](https://leetcode.com/problems/arithmetic-slices/solution/)
