## [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/#/description)

### Description

Given a string **s**, find the longest palindromic substring in **s**. You may assume that the maximum length of **s** is 1000.

**Example**:

```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```

**Example**:

```
Input: "cbbd"
Output: "bb"
```

**Difficult:** `Medium`

**Tags:** `String`

### Solution One - In Top Solutions

对于字符串中的任意一个字符$C_i$，以其为中心，向两边扩展并检测是否相等，相等则继续扩展，否则表明以该字符$C_i$为中心的回文子字符串达到最长了

```c++
class Solution {
public:
    string longestPalindrome(string s) {
        string::size_type max_length = 0;
        string::size_type start = 0;
        for (string::size_type i = 0; i < s.size();)
        {
            if (i + max_length/2 >= s.size())
            {
                break;
            }
            string::size_type j = i, k = i;
            while (k < s.size()-1 && s[k] == s[k+1])
            {
                ++k;	// 跳过相同的字符
            }
            i = k + 1;
            while (j > 0 && k < s.size()-1 && s[j - 1] == s[k + 1])
            {
                --j;
                ++k;
            }
            if (max_length < k - j + 1)
            {
                start = j;
                max_length = k - j + 1;
            }
        }
        return s.substr(start, max_length);
    }
};
```
