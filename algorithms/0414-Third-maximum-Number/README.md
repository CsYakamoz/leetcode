## [414. Third maximum Number](https://leetcode.com/problems/third-maximum-number/description/)

### Description

Given a **non-empty** array of integers, return the **third** maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).

**Example 1:**

```
Input: [3, 2, 1]

Output: 1

Explanation: The third maximum is 1.

```

**Example 2:**

```
Input: [1, 2]

Output: 2

Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

```

**Example 3:**

```
Input: [2, 2, 3, 1]

Output: 1

Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.
```

**Difficulty:** `Easy`

**Tags:** `Array`

### Solution One

`Hash Table`

第 4 行相当于`for (auto i: nums) s.insert(i);`

```c++
class Solution {
public:
    int thirdMax(vector<int>& nums) {
        set<int> s(nums.begin(), nums.end());
        auto iter = s.rbegin();
        if (s.size() > 2)
        {
            iter++;
            iter++;
            return *iter;
        }
        else
        {
            return *iter;
        }
    }
};
```

### Solution Two - In Top Solutions

[Short Clear C++ solution, no set or pq.](https://discuss.leetcode.com/topic/67186/short-clear-c-solution-no-set-or-pq)

```c++
int thirdMax(vector<int>& nums) {
    long long a, b, c;
    a = b = c = LLONG_MIN;
    for (auto num : nums) {
        if (num <= c || num == b || num == a) continue;
        c = num;
        if (c > b) swap(b, c);
        if (b > a) swap(a, b);
    }
    return c == LLONG_MIN ? a : c;
}
```
