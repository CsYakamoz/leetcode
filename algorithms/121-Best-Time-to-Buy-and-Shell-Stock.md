## [121. Best Time to Buy and Shell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/)

### Description

Say you have an array for which the *i*th element is the price of a given stock on day *i*.

If you were only permitted to complete at most one transaction (ie, buy one and sell one share of the stock), design an algorithm to find the maximum profit.

**Example 1:**

```
Input: [7, 1, 5, 3, 6, 4]
Output: 5

max. difference = 6-1 = 5 (not 7-1 = 6, as selling price needs to be larger than buying price)

```

**Example 2:**

```
Input: [7, 6, 4, 3, 1]
Output: 0

In this case, no transaction is done, i.e. max profit = 0.
```



**Difficult:** `Easy`

**Tags:** `Array` `Dynamic Programming`



### Solution One

```c++
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if (prices.empty())
            return 0;

        int minPrice = prices[0];
        int maxProfit = 0;
        for (size_t i = 1; i < prices.size(); i++)
        {
            if (prices[i] > minPrice)
                maxProfit = max(maxProfit, prices[i] - minPrice);
            else if (prices[i] < minPrice)
                minPrice = prices[i];
        }
        return maxProfit;
    }
};
```



