## [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/description/)

### Description

You are climbing a stair case. It takes _n_ steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

**Note:** Given _n_ will be a positive integer.

**Difficult:** `Easy`

**Tags:** `Dynamic Programming`

### Solution One

```c++
class Solution {
public:
    int climbStairs(int n) {
        if (n < 3)
        {
            return n;
        }
        int pre2way = 1;
        int pre1way = 2;
        int i = 3;
        while (i <= n)
        {
            int way = pre1way + pre2way;
            pre2way = pre1way;
            pre1way = way;
            i++;
        }
        return pre1way;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
	int climbStairs(int n) {
		vector<int> viWays(n, 1);
		if (n == 0)
			return 0;
		if (n == 1)
			return 1;
		if (n == 2)
			return 2;
		viWays[n - 2] = 2;
		for (int i = n - 3; i >= 0; i--)
		{
			viWays[i] = viWays[i + 1] + viWays[i + 2];
		}
		return viWays[0];
	}
};
```

### Others Solutions

[70. Climbing Stairs - Solutions](https://leetcode.com/problems/climbing-stairs/solution/)
