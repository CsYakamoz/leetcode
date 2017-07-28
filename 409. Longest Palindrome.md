## [409. Longest Palindrome](https://leetcode.com/problems/longest-palindrome/tabs/description)

### Description

Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example `"Aa"` is not considered a palindrome here.

**Note:**
Assume the length of given string will not exceed 1,010.

**Example:**

```
Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
```



**Difficult:** `Easy`

**Tags:** `Hash Table`



### Solution One

```c++
class Solution {
public:
    int longestPalindrome(string s) {
        vector<int> index(128, 0);
        int res = 0;
        for (auto i : s)
        {
            index[i]++;
            if (index[i] == 2)
            {
                res += 2;
                index[i] = 0;
            }
        }
        return res < s.size() ? res + 1 : res;
    }
};
```



### Solution Two

[What are the odds? (Python & C++)](https://discuss.leetcode.com/topic/61338/what-are-the-odds-python-c)

