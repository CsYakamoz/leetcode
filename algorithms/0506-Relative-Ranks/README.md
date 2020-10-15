## [506. Relative Ranks](https://leetcode.com/problems/relative-ranks/description/)

### Description

Given scores of **N** athletes, find their relative ranks and the people with the top three highest scores, who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".

**Example 1:**

```
Input: [5, 4, 3, 2, 1]
Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal".
For the left two athletes, you just need to output their relative ranks according to their scores.
```

**Note:**

1. N is a positive integer and won't exceed 10,000.
2. All the scores of athletes are guaranteed to be unique.

**Difficult:** `Easy`

**Tags:**

### Solution One

```c++
class Solution {
public:
    vector<string> findRelativeRanks(vector<int>& nums) {
        vector<pair<int, int>> vec;
        vector<string> res(nums.size(), "");
        vector<string> Medal = { "Gold Medal", "Silver Medal", "Bronze Medal" };
        for (size_t i = 0; i < nums.size(); i++)
        {
            vec.push_back(pair<int, int>(nums[i], i));
        }
        sort(vec.begin(), vec.end(),
            [](pair<int, int> x, pair<int, int> y) { return x.first > y.first; });
        for (size_t i = 0; i < vec.size(); i++)
        {
            if (i < 3)
            {
                res[vec[i].second] = Medal[i];
            }
            else
            {
                res[vec[i].second] = to_string(i + 1);
            }
        }
        return res;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    vector<string> findRelativeRanks(vector<int>& nums) {
        vector<pair<int, int>> np;
        for (int i = 0; i < nums.size(); ++i)
            np.push_back({i, nums[i]});

        sort(np.begin(), np.end(), [](pair<int, int> &p1, pair<int, int> &p2) {
           return p1.second > p2.second;
        });

        vector<string> res(nums.size(), "");
        vector<string> tp{"Gold Medal", "Silver Medal", "Bronze Medal"};
        for (int i = 0; i < np.size(); ++i) {
            if (i < 3)
                res[np[i].first] = tp[i];
            else
                res[np[i].first] = to_string(i+1);
        }
        return res;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    vector<string> findRelativeRanks(vector<int>& nums) {
        vector<string> res(nums.size());
        priority_queue<pair<int,int>> temp;
        for(size_t i=0,len=nums.size();i<len;++i)
            temp.push({nums[i],i});
        int count=1;
        while(temp.size())
        {
            auto elem=temp.top();
            temp.pop();
            if(count==1)
                res[elem.second]="Gold Medal";
            else if(count==2)
                res[elem.second]="Silver Medal";
            else if(count==3)
                res[elem.second]="Bronze Medal";
            else
                res[elem.second]=to_string(count);
            ++count;
        }
        return res;
    }
};
```
