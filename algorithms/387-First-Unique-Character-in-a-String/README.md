## [387. First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/#/description)

### Description

Given a string, find the first non-repeating character in it and return it's index. if it doesn't exist, return -1.

**Examples:**

```
s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
```

**Note:** You may assume the string contain only lowercase letters.

**Difficult:** `Easy`

**Tags:**

### Solution One

原始办法，二层循环

对于第 i 个元素，在 `[i + 1, s.size())`区间查看是否有相同元素

若有，则不是要找的元素

若无，则直接返回

```c++
class Solution {
public:
    int firstUniqChar(string s) {
        vector<char> index(128, -1);
        for (size_t i = 0; i < s.size(); i++)
        {
            if (index[s[i]] != 1)
            {
                bool isRepeated = false;
                for (size_t j = i + 1; j < s.size(); j++)
                {
                    if (s[i] == s[j])
                    {
                        isRepeated = true;
                    }
                }
                if (isRepeated == false)
                {
                    return i;
                }
                index[s[i]] = 1;
            }
        }
        return -1;
    }
};
```

### Solution Two

isExist 用来判断 s[i] 是否重复出现过， 若 isExist[s[i]] != -1，代表出现过（即重复），否则代表第一次出现

若第一次出现，则令 isExist[s[i]] = i，并且将 s[i] 放到 tmp 后面

若不是第一次出现，在 tmp 中寻找 s[i]，将其删除

```c++
class Solution {
public:
    int firstUniqChar(string s) {
        vector<int> isExist(128, -1);
        vector<char> tmp;
        for (size_t i = 0; i < s.size(); i++)
        {
            if (isExist[s[i]] != -1)
            {	// Not first find
                auto p = find(tmp.begin(), tmp.end(), s[i]);
                if (p != tmp.end())
                {
                    tmp.erase(p);
                }
            }
            else
            {	// first find
                isExist[s[i]] = i;
                tmp.push_back(s[i]);
            }
        }
        if (tmp.empty())
        {
            return -1;
        }
        else
        {
            return isExist[tmp[0]];
        }
    }
};
```

### Solution Three - In Top Solution

如果字符`c`是第一次出现，则`index[c] = i`，否则`index[c] = s.size()`

```c++
class Solution {
public:
    int firstUniqChar(string s) {
        vector<int> index(26, -1);
        for (size_t i = 0; i < s.size(); i++)
        {
            auto c = s[i] - 'a';
            if (index[c] == -1)
            {	// s[i] first appears
                index[c] = i;
            }
            else
            {	// s[i] Not first appear
                index[c] = s.size();
            }
        }
        int firstCharIndex = s.size();
        for (const auto i : index)
        {
            if (i != - 1)
            {
                firstCharIndex = min(firstCharIndex, i);
            }
        }
        return firstCharIndex == s.size() ? -1 : firstCharIndex;
    }
};
```
