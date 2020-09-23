## [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/description/)

### Description

Given a collection of intervals, merge all overlapping intervals.

For example,
Given `[1,3],[2,6],[8,10],[15,18]`,
return `[1,6],[8,10],[15,18]`.



**Difficult:** `Medium`

**Tags:** `Array` `Sort`



### Solution One - In Top Solutions

[C++ 10 line solution. easing understanding](https://discuss.leetcode.com/topic/20263/c-10-line-solution-easing-understanding)

```c++
class Solution {
public:
    vector<Interval> merge(vector<Interval>& intervals) {
        if (intervals.size() <= 1)
            return intervals;
        vector<Interval> res;
        sort(intervals.begin(), intervals.end(), [](Interval &x, Interval &y) {
            return x.start < y.start;
        });
        res.push_back(intervals[0]);
        for (size_t i = 1; i < intervals.size(); i++)
        {
            if (intervals[i].start <= res.back().end)
                res.back().end = max(res.back().end, intervals[i].end);
            else
                res.push_back(intervals[i]);
        }
        return res;
    }
};
```



