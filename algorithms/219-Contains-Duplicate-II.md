## [219. Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/description/)

### Description

Given an array of integers and an integer *k*, find out whether there are two distinct indices *i* and *j* in the array such that **nums[i] = nums[j]** and the **absolute** difference between *i* and *j* is at most *k*.



**Difficult:** `Easy`

**Tags:** `Array` `Hash Table`



### Solution One

```c++
class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        unordered_map<int, int> m;
        for (size_t i = 0; i < nums.size(); i++)
        {
            if (m.find(nums[i]) != m.end())
            {
                if (i - m[nums[i]] <= k)
                {
                    return true;
                }
            }
            m[nums[i]] = i;
        }
        return false;
    }
};
```



### Solution Two - In Top Solutions

[C++ solution with unordered_set](https://discuss.leetcode.com/topic/15045/c-solution-with-unordered_set)

```c++
class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k)
    {
       unordered_set<int> s;
       
       if (k <= 0) return false;
       if (k >= nums.size()) k = nums.size() - 1;
       
       for (int i = 0; i < nums.size(); i++)
       {
           if (i > k) s.erase(nums[i - k - 1]);
           if (s.find(nums[i]) != s.end()) return true;
           s.insert(nums[i]);
       }
       
       return false;
    }
};
```



