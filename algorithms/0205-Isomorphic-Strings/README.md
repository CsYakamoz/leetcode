## [205. Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/#/description)

### Description

Given two string **s** and **t**, determine if they are isomorphic.

Two strings are isomorphic if the characters in **s** can be replaced to get **t**.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

For example,
Given `"egg"`, `"add"`, return true.

Given `"foo"`, `"bar"`, return false.

Given `"paper"`, `"title"`, return true.

**Note:**
You may assume both **s** and **t** have the same length.

**Difficulty:** `Easy`

**Tags:** `Hash Table`

### Solution One

```c++
class Solution {
public:
    bool isIsomorphic(string s, string t) {
        vector<int> indexOfs(128, -1);
        vector<int> indexOft(128, -1);
        size_t length = s.size();
        for(size_t i = 0; i < length; i++)
        {
            if (indexOfs[s[i]] != indexOft[t[i]])
            {
                return false;
            }
            indexOfs[s[i]] = i;
            indexOft[t[i]] = i;
        }
        return true;
    }
};
```
