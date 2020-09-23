## [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/tabs/description)

### Description

Given two strings *s* and *t*, write a function to determine if *t* is an anagram of *s*.

For example,
*s* = "anagram", *t* = "nagaram", return true.
*s* = "rat", *t* = "car", return false.

**Note:**
You may assume the string contains only lowercase alphabets.

**Follow up:**
What if the inputs contain unicode characters? How would you adapt your solution to such case?



**Difficult:** `Easy`

**Tags:** `Hash Table` `Sort`



### Solution One

`Hash Table`

```
class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.size() != t.size())
        {
            return false;
        }
        vector<int> alphabets(26, 0);
        for (char i : s)
            alphabets[i - 'a']++;
        for (char i : t)
        {
            alphabets[i - 'a']--;
            if (alphabets[i - 'a'] < 0)
            {
                return false;
            }
        }
        return true;
    }
};
```



### Solution Two

`Sort`

```
class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.size() != t.size())
        {
            return false;
        }
        sort(s.begin(), s.end());
        sort(t.begin(), t.end());
        for (size_t i = 0; i < s.size(); i++)
        {
            if (s[i] != t[i])
            {
                return false;
            }
        }
        return true;
    }
};
```



### Solution Three - In Top Solutions

```c++
class Solution {
public:
    bool isAnagram(string s, string t) {
        int m[26][2]={0};
        for(auto c:s)m[c-'a'][0]++;
        for(auto c:t)m[c-'a'][1]++;
        for(int i =0;i<25;i++)
        {
            if(m[i][0]!=m[i][1])return false;
        }
        return true;
    }
};
```



