## [747. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/description/)

### Description

On a staircase, the `i`-th step has some non-negative cost `cost[i]` assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

**Example 1:**

```
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
```

**Example 2:**

```
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
```

**Note:**

1. `cost` will have a length in the range `[2, 1000]`.
2. Every `cost[i]` will be an integer in the range `[0, 999]`.

**Difficult:** `Easy`

**Tags:** `Array` `Dynamic Programming`

### Solution One

```c++
class Solution
{
  public:
    // every cost[i] will be an integer in the range[0, 999]
    int minCostClimbingStairs(vector<int> &cost)
    {
        // cost will have a length in the range [2, 1000]
        if (cost.size() == 2)
            return min(cost[0], cost[1]);

        vector<int> res;
        res.push_back(cost[0]);
        res.push_back(cost[1]);

        for (size_t i = 2; i < cost.size(); i++)
        {
            res.push_back(cost[i] + min(res[i - 1], res[i - 2]));
        }

        return min(res[res.size() - 1], res[res.size() - 2]);
    }
};
```
