## [436. Find Right Interval](https://leetcode.com/problems/find-right-interval/description/)

### Description

Given a set of intervals, for each of the interval i, check if there exists an interval j whose start point is bigger than or equal to the end point of the interval i, which can be called that j is on the "right" of i.

For any interval i, you need to store the minimum interval j's index, which means that the interval j has the minimum start point to build the"right"relationship for interval i. If the interval j doesn't exist, store -1 for the interval i. Finally, you need output the stored value of each interval as an array.

**Note:**

1. You may assume the interval's end point is always bigger than its start point.
2. You may assume none of these intervals have the same start point.

**Example 1:**

```
Input: [ [1,2] ]

Output: [-1]

Explanation: There is only one interval in the collection, so it outputs -1.

```

**Example 2:**

```
Input: [ [3,4], [2,3], [1,2] ]

Output: [-1, 0, 1]

Explanation: There is no satisfied "right" interval for [3,4].
For [2,3], the interval [3,4] has minimum-"right" start point;
For [1,2], the interval [2,3] has minimum-"right" start point.

```

**Example 3:**

```
Input: [ [1,4], [2,3], [3,4] ]

Output: [-1, 2, -1]

Explanation: There is no satisfied "right" interval for [1,4] and [3,4].
For [2,3], the interval [3,4] has minimum-"right" start point.
```

**Difficult:** `Medium`

**Tags:** `Binary Search`

### Solution One

```c++
class Solution {
public:
    vector<int> findRightInterval(vector<Interval>& intervals) {
        int length = intervals.size();

        vector<pair<int,int>> start;
        for (int i = 0; i < length; i++)
            start.push_back({ intervals[i].start, i });

        sort(start.begin(), start.end(), [](pair<int, int> &x, pair<int,int> &y) {
            return x.first < y.first;
        });

        vector<int> result;
        for (int i = 0; i < length; i++)
        {
            int end = intervals[i].end;
            result.push_back(binarySearch(start, end));
        }
        return result;
    }

private:
    int binarySearch(const vector<pair<int, int>> &nums, const int val)
    {
        int left = 0;
        int right = nums.size();
        while (left < right)
        {
            int mid = left + (right - left) / 2;
            int start = nums[mid].first;
            if (start >= val)
            {
                right = mid;
            }
            else
            {
                left = mid + 1;
            }
        }
        // right points to the first elements
        // that is bigger than or equal to the val
        return right == nums.size() ? -1 : nums[right].second;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
    struct cmp {
        bool operator() (pair<int,int>& a, pair<int,int>& b) {
            return a.first < b.first;
        }
        bool operator() (pair<int,int>& a, int b) {
            return a.first < b;
        }
        bool operator() (int a, pair<int,int>& b) {
            return a < b.first;
        }
    };
public:
    vector<int> findRightInterval(vector<Interval>& intervals) {
        int n = intervals.size();
        vector<pair<int, int>> starts(n);
        for (int i = 0; i < n; ++i) {
            starts[i] = {intervals[i].start, i};
        }
        sort(starts.begin(), starts.end());
        vector<int> res(n);
        for (int i = 0; i < n; ++i) {
            // auto it = lower_bound(starts.begin(), starts.end(), make_pair(intervals[i].end, -1));
            auto it = lower_bound(starts.begin(), starts.end(), intervals[i].end, cmp());
            res[i] = it == starts.end()? -1 : it->second;
        }
        return res;
    }
};
```
