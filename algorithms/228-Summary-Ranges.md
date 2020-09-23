## [228. Summary Ranges](https://leetcode.com/problems/summary-ranges/description/)

### Description

Given a sorted integer array without duplicates, return the summary of its ranges.

For example, given `[0,1,2,4,5,7]`, return `["0->2","4->5","7"].`

**Difficult:** `Medium`

**Tags:** `Array`

### Solution One

```c++
class Solution {
public:
    vector<string> summaryRanges(vector<int>& nums) {
        vector<string> res;
        if (nums.empty()) return res;

        int start = nums[0];
        int end = nums[0];
        for (size_t i = 1; i < nums.size(); i++)
        {
            if (nums[i] - nums[i-1] != 1)
            {
                if (start == end)
                    res.push_back(to_string(start));
                else
                    res.push_back(to_string(start) + "->" + to_string(end));
                start = nums[i];
                end = nums[i];
            }
            else
            {
                end = nums[i];
            }
        }
        if (start == end)
            res.push_back(to_string(start));
        else
            res.push_back(to_string(start) + "->" + to_string(end));
        return res;
    }
};
```

### Solution Two - In Top Solutions

[10 line c++ easy understand](https://discuss.leetcode.com/topic/17117/10-line-c-easy-understand)

```c++
   vector<string> summaryRanges(vector<int>& nums) {
    const int size_n = nums.size();
    vector<string> res;
    if ( 0 == size_n) return res;
    for (int i = 0; i < size_n;) {
        int start = i, end = i;
        while (end + 1 < size_n && nums[end+1] == nums[end] + 1) end++;
        if (end > start) res.push_back(to_string(nums[start]) + "->" + to_string(nums[end]));
        else res.push_back(to_string(nums[start]));
        i = end+1;
    }
    return res;
}
```
