## [692. Top K Frequent Words](https://leetcode.com/problems/top-k-frequent-words/description/)

### Description

Given a non-empty list of words, return the _k_ most frequent elements.

Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

**Example 1:**

```
Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
Output: ["i", "love"]
Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.

```

**Example 2:**

```
Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
Output: ["the", "is", "sunny", "day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
    with the number of occurrence being 4, 3, 2 and 1 respectively.

```

**Note:**

1. You may assume _k_ is always valid, 1 ≤ _k_ ≤ number of unique elements.
2. Input words contain only lowercase letters.

**Follow up:**

1. Try to solve it in _O_(_n_ log _k_) time and _O_(_n_) extra space.
2. Can you solve it in _O_(_n_) time with only _O_(_k_) extra space?

**Difficult:** `Medium`

**Tags:** `Hash Table` `Heap` `Trie`

### Solution One

```c++
class Solution {
public:
    vector<string> topKFrequent(vector<string>& words, int k)
    {
        vector<string> res;
        unordered_map<string, int> hash;
        int maxFrequent = INT_MIN;

        for (auto &str : words) {
            hash[str]++;
            maxFrequent = max(hash[str], maxFrequent);
        }

        vector<vector<string>> vec(maxFrequent + 1);
        for (auto &iter : hash)
            vec[iter.second].push_back(iter.first);

        for (auto iter = vec.rbegin(); iter != vec.rend() && k != 0; iter++)
            if (!iter->empty()) {
                sort(iter->begin(), iter->end());

                for (auto ptr = iter->begin(); ptr != iter->end() && k != 0; ptr++, k--)
                    res.push_back(*ptr);
            }

        return res;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Compare{
public:
    bool operator()(pair<int, string> a, pair<int, string> b){
        if(a.first==b.first)
            return a.second>b.second;
        return a.first<b.first;
    }
};

class Solution {
public:
    // O(N logK)
    vector<string> topKFrequent1(vector<string>& words, int k) {
        unordered_map<string, int> m;
        for(auto w:words) m[w]++;
        priority_queue<pair<int, string>, vector<pair<int, string> >, Compare> pQ;
        for(auto it:m) pQ.push({it.second, it.first});
        vector<string> res;
        while(k){
            res.push_back(pQ.top().second);
            pQ.pop();
            k--;
        }
        return res;
    }
    // O(N) bucket solution
    vector<string> topKFrequent(vector<string>& words, int k) {
        unordered_map<string, int> m;
        for(auto w:words) m[w]++;
        int mx = -1;
        for(auto it:m) mx = max(mx, it.second);
        vector<set<string>> bucket(mx+1, set<string>());
        for(auto it:m) bucket[it.second].insert(it.first);
        vector<string> res;
        int end = bucket.size()-1;
        while(k){
            for(auto s:bucket[end]){
                res.push_back(s);
                bucket[end].erase(s);
                k--;
                if(k==0) break;
            }
            while(bucket[end].empty()) end--;
        }
        return res;
    }
};
```
