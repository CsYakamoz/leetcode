## [135. Candy](https://leetcode.com/problems/candy/description/)

### Description

There are _N_ children standing in a line. Each child is assigned a rating value.

You are giving candies to these children subjected to the following requirements:

- Each child must have at least one candy.
- Children with a higher rating get more candies than their neighbors.

What is the minimum candies you must give?

**Difficulty:** `Hard`

**Tags:** `Greedy`

### Solution One

```c++
class Solution {
public:
    int candy(vector<int>& ratings) {
        int res = 1;
        int prev = 1;
        int oldScore = 0;
        bool isSorted = true;
        for (size_t i = 1; i < ratings.size(); i++)
        {
            if (ratings[i] > ratings[i - 1])
            {
                if (!isSorted)
                {
                    isSorted = true;
                    prev = 1;
                }
                prev++;
            }
            else if (ratings[i] < ratings[i - 1])
            {
                if (isSorted)
                {
                    isSorted = false;
                    oldScore = prev;
                    prev = 0;
                }
                prev++;
                if (oldScore == prev)
                {
                    oldScore++;
                    res++;
                }
            }
            else
            {
                isSorted = true;
                prev = 1;
            }
            res += prev;
        }
        return res;
    }
};
```

### Other Solutions

[135. Candy - Solution](https://leetcode.com/problems/candy/solution/)
