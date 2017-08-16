## [57. Insert Interval](https://leetcode.com/problems/insert-interval/description/)

### Description

Given a set of *non-overlapping* intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

**Example 1:**
Given intervals `[1,3],[6,9]`, insert and merge `[2,5]` in as `[1,5],[6,9]`.

**Example 2:**
Given `[1,2],[3,5],[6,7],[8,10],[12,16]`, insert and merge `[4,9]` in as `[1,2],[3,10],[12,16]`.

This is because the new interval `[4,9]` overlaps with `[3,5],[6,7],[8,10]`.



**Difficult:** `Medium`

**Tags:** `Hard`



### Solution One

来源于 [56. Merget Interval](https://leetcode.com/problems/merge-intervals/description/) 中的 Top Solutions [C++ 10 line solution. easing understanding](https://discuss.leetcode.com/topic/20263/c-10-line-solution-easing-understanding)

```c++
class Solution {
public:
    vector<Interval> insert(vector<Interval>& intervals, Interval newInterval) {
        intervals.push_back(newInterval);
        if (intervals.size() == 1) {
            return intervals;
        }
        vector<Interval> res;
        sort(intervals.begin(), intervals.end(), [](Interval &x, Interval &y) {
            return x.start < y.start;
        });
        res.push_back(intervals[0]);
        for (size_t i = 0; i < intervals.size(); i++)
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



