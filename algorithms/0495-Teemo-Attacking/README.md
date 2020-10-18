## [495. Teemo Attacking](https://leetcode.com/problems/teemo-attacking/#/description)

### Description

In LLP world, there is a hero called Teemo and his attacking can make his enemy Ashe be in poisoned condition. Now, given the Teemo's attacking **ascending** time series towards Ashe and the poisoning time duration per Teemo's attacking, you need to output the total time that Ashe is in poisoned condition.

You may assume that Teemo attacks at the very beginning of a specific time point, and makes Ashe be in poisoned condition immediately.

**Example 1:**

```
Input: [1,4], 2
Output: 4
Explanation: At time point 1, Teemo starts attacking Ashe and makes Ashe be poisoned immediately.
This poisoned status will last 2 seconds until the end of time point 2.
And at time point 4, Teemo attacks Ashe again, and causes Ashe to be in poisoned status for another 2 seconds.
So you finally need to output 4.

```

**Example 2:**

```
Input: [1,2], 2
Output: 3
Explanation: At time point 1, Teemo starts attacking Ashe and makes Ashe be poisoned.
This poisoned status will last 2 seconds until the end of time point 2.
However, at the beginning of time point 2, Teemo attacks Ashe again who is already in poisoned status.
Since the poisoned status won't add up together, though the second poisoning attack will still work at time point 2, it will stop at the end of time point 3.
So you finally need to output 3.

```

**Note:**

1. You may assume the length of given time series array won't exceed 10000.
2. You may assume the numbers in the Teemo's attacking time series and his poisoning time duration per attacking are non-negative integers, which won't exceed 10,000,000.

**Difficulty:** `Medium`

**Tags:** `Array`

### Solution One

若`timeSeries[i] - timeSeries[i-1] < duration`成立

则两次攻击间隔时间小于`duration`，故`time += timeSeries[i] - timeSeries[i-1]`

否则，`time += duration`

对于最后一次攻击，直接加上`duration`即可

```c++
class Solution {
public:
    int findPoisonedDuration(vector<int>& timeSeries, int duration) {
        if (timeSeries.empty())
        {
            return 0;
        }
        int time = 0;
        size_t i = 1;
        while (i < timeSeries.size())
        {
            if (timeSeries[i] - timeSeries[i-1] < duration)
            {
                time += timeSeries[i] - timeSeries[i - 1];
            }
            else
            {
                time += duration;
            }
            ++i;
        }
        time += duration;
        return time;
    }
};
```

### Solution Two - In Top Solutions

[Python Solution for Teemo](https://discuss.leetcode.com/topic/77360/python-solution-for-teemo)

`timeSeries[i] - timeSeries[i-1]`表示两次攻击间隔时间

`duration - (timeSeries[i] - timeSeries[i-1])`表示持续时间减去两次攻击间隔时间

`max(0,duration - (timeSeries[i] - timeSeries[i-1]))`，因为`ans = duration * len(timeSeries)`，其中可能加多了时间，所以这里要减去

```python
class Solution(object):
    def findPoisonedDuration(self, timeSeries, duration):
        ans = duration * len(timeSeries)
        for i in range(1,len(timeSeries)):
            ans -= max(0, duration - (timeSeries[i] - timeSeries[i-1]))
        return ans
```
