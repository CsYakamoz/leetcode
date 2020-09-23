## [347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/description/)

### Description

Given a non-empty array of integers, return the **k** most frequent elements.

For example,
Given `[1,1,1,2,2,3]` and k = 2, return `[1,2]`.

**Note: **

- You may assume *k* is always valid, 1 ? *k* ? number of unique elements.
- Your algorithm's time complexity **must be** better than O(*n* log *n*), where *n* is the array's size.





**Difficult:** `Medium`

**Tags:** `Hash Table` `Heap`



### Solution One

来源于题目 [451. Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/discuss/) 的 Top Solutions: [C++ O(n) solution without sort()](https://discuss.leetcode.com/topic/66045/c-o-n-solution-without-sort)

```c++
class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        vector<int> res;
        unordered_map<int, int> hash;
        int maxFrequent = INT_MIN;

        for (auto &i : nums) {
            hash[i]++;
            maxFrequent = max(maxFrequent, hash[i]);
        }

        vector<vector<int>> vec(maxFrequent + 1);
        for (auto &iter : hash)
            vec[iter.second].push_back(iter.first);

        for (auto iter = vec.crbegin(); iter != vec.crend() && k != 0; iter++) {
            if (!iter->empty()) {
                for (auto ptr = iter->begin(); ptr != iter->end() && k != 0; ptr++, k--)
                    res.push_back(*ptr);
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
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int,int> hash; //[num, freq]
        
        for(int n:nums)
            hash[n]++;
        
        priority_queue<pair<int,int>> heap;
        for(auto item:hash)
        {
            heap.push(make_pair(item.second, item.first));
        }
        
        vector<int> res;
        while(k--)
        {
            res.push_back(heap.top().second);
            heap.pop();
        }
        
        return res;
    }
};
```



