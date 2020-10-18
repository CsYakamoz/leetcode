## [754. Reach a Nnumber](https://leetcode.com/problems/reach-a-number/description/)

### Description

You are standing at position `0` on an infinite number line. There is a goal at position `target`.

On each move, you can either go left or right. During the _n_-th move (starting from 1), you take _n_ steps.

Return the minimum number of steps required to reach the destination.

**Example 1:**

```
Input: target = 3
Output: 2
Explanation:
On the first move we step from 0 to 1.
On the second step we step from 1 to 3.
```

**Example 2:**

```
Input: target = 2
Output: 3
Explanation:
On the first move we step from 0 to 1.
On the second move we step  from 1 to -1.
On the third move we step from -1 to 2.
```

**Note:**

`target` will be a non-zero integer in the range `[-10^9, 10^9]`.

**Difficulty:** `Medium`

**Tags:** `Math`

### Solution One

```c++
class Solution {
public:
    int reachNumber(int target) {
        if (target == 0) {
            return 0;
        }

        target = abs(target);
        int left = 1;
        int right = target;
        long long sum = 1;
        while (left < right) {
            int mid = (right - left) / 2 + left;
            sum = Solution::sum(mid);
            if (sum >= target) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        sum = Solution::sum(right);
        while ((sum - target) & 1) {
            sum += ++right;
        }

        return right;
    }

private:
    long long sum(int n) {
        long long sum = 1;
        if (n & 1) {
            sum *= (1 + n) / 2;
            sum *= n;
        } else {
            sum *= n / 2;
            sum *= (1 + n);
        }
        return sum;
    }
};
```

### Solution

[754. Reach a Number](https://leetcode.com/problems/reach-a-number/solution/)
