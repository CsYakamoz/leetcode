## [451. Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/description/)

### Description

Given a string, sort it in decreasing order based on the frequency of characters.

**Example 1:**

```
Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.

```

**Example 2:**

```
Input:
"cccaaa"

Output:
"cccaaa"

Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.

```

**Example 3:**

```
Input:
"Aabb"

Output:
"bbAa"

Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
```

**Difficult:** `Medium`

**Tags:** `Hash Table` `Heap`

### Solution One

```c++
class Solution {
public:
    string frequencySort(string s) {
        map<char, int> m;
        for (auto i : s)
            m[i]++;
        vector<pair<char, int>> v(m.begin(), m.end());
        sort(v.rbegin(), v.rend(),
            [](pair<char, int> x, pair<char, int> y) {
            return x.second < y.second;
        });
        string res;
        for (auto i : v)
        {
            while (i.second != 0)
            {
                res.push_back(i.first);
                i.second--;
            }
        }
        return res;
    }
};
```

### Solution Two

```c++
class Solution {
public:
    string frequencySort(string s) {
        map<char, int> m;
        for (auto i : s)
            m[i]++;
        vector<pair<char, int>> v(m.begin(), m.end());
        make_heap(v.begin(), v.end(),
            [](pair<char, int> x, pair<char, int> y) {
            return x.second < y.second;
        });
        string res;
        while (!v.empty())
        {
            auto i = v.front();
            res += string(i.second, i.first);
            pop_heap(v.begin(), v.end(),
                [](pair<char, int> x, pair<char, int> y) {
                return x.second < y.second;
            });
            v.pop_back();
        }
        return res;
    }
};
```

### Solution Three - In Top Solutions

[C++ O(n) solution without sort()](https://discuss.leetcode.com/topic/66045/c-o-n-solution-without-sort)

```c++
class Solution {
public:
    string frequencySort(string s) {
        unordered_map<char,int> freq;
        vector<string> bucket(s.size()+1, "");
        string res;

        //count frequency of each character
        for(char c:s) freq[c]++;
        //put character into frequency bucket
        for(auto& it:freq) {
            int n = it.second;
            char c = it.first;
            bucket[n].append(n, c);
        }
        //form descending sorted string
        for(int i=s.size(); i>0; i--) {
            if(!bucket[i].empty())
                res.append(bucket[i]);
        }
        return res;
    }
};
```
