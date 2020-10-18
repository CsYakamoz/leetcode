## [14. Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/#/description)

### Description

Write a function to find the longest common prefix string amongst an array of strings.

**Difficulty:** `Easy`

**Tags:** `String`

### Solution One

令`result = strs[0]`，以`result`为基准

接着，获取每个`strs[i]`(0 < i < strs.size()) 和`result`共同前缀，并将共同前缀赋值给`result`

这样当循环结束时，就得到最长的公共前缀

**Note：**第 10 行的循环条件可以改为 `while (i < strs.size() && !result.empty())`

```c++
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.size() == 0)
        {
            return string("");
        }
        string result = strs[0];
        size_t i = 1;
        while (i < strs.size())
        {
            size_t j = 0;
            size_t length = max(result.size(), strs[i].size());
            while (j < length && result[j] == strs[i][j])
            {
                ++j;
            }
            result = result.substr(0, j);
            ++i;
        }
        return result;
    }
};
```

### Solution Two - In Top Solutions

`Sort`

[Sorted the array, Java solution, 2 ms](https://discuss.leetcode.com/topic/27913/sorted-the-array-java-solution-2-ms)

Sort the array first, and then you can simply compare the first and last elements in the sorted array.

```c++
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.size() == 0)
        {
            return string("");
        }
        sort(strs.begin(), strs.end());
        string a = strs[0];
        string b = strs[strs.size() - 1];
        string result;
        size_t length = max(a.size(), b.size());
        for (size_t i = 0; i < length; i++)
        {
            if (a[i] == b[i])
            {
                result += a[i];
            }
            else
            {
                return result;
            }
        }
        return result;
    }
};
```

### Solution Three - In Top Solutions

判断所有`strs[i]`的第`k`个字符是否相同，若相同，则添加到`result`

否则，直接返回`result`

```c++
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.size() == 0)
        {
            return string("");
        }
        string result("");
        for (size_t k = 0; k < strs[0].size(); k++)
        {
            size_t i = 1;
            for (; i < strs.size() && strs[i].size() > k; i++)
            {
                if (strs[0][k] != strs[i][k])
                {
                    return result;
                }
            }
            if (i == strs.size())
            {
                result += strs[0][k];
            }
        }
        return result;
    }
};
```
