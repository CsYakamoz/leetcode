## [344. Reverse String](https://leetcode.com/problems/reverse-string/#/description)

### Description

Write a function that takes a string as input and returns the string reversed.

**Example:**
Given s = "hello", return "olleh".

**Difficult:** `Easy`

**Tags:** `Two Pointer` `String`

### Solution One

1. 令 i = s.length() - 1，使其指向字符串 s 的最后一个字符
2. 若 i >= 0，则将 s[i]添加到字符串 result，否则返回字符串 result
3. i--

```c++
class Solution {
public:
    string reverseString(string s) {
        string result = "";
        for (int i = s.length() - 1; i >= 0; --i)
        {
            result += s[i];
        }
        return result;
    }
};
```

### Solution Two - In Top Solutions

将字符串两边的字符对换，判断条件用`left < right`而不用`left != right`，是因为有可能字符串的长度为偶数

```c++
class Solution {
public:
    string reverseString(string s) {
        int right = s.size() - 1;
        int left = 0;
        while (left < right)
            swap(s[left++], s[right--]);
        return s;
    }
};
```
