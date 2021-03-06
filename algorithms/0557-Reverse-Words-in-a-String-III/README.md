## [557. Reverse Words in a String III](https://leetcode.com/problems/reverse-words-in-a-string-iii/#/description)

### Description

Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

**Example 1:**

```
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
```

**Note:** In the string, each word is separated by single space and there will not be any extra space in the string.

**Difficulty:** `Easy`

**Tags:** `String`

### Solution One

使用 istringstream 获得每个单词

调用 STL 中的 reverse

添加到 result

```c++
class Solution {
public:
    string reverseWords(string s) {
        if (s.empty())
        {
            return s;
        }
        string result;
        string word;
        istringstream iss(s);
        while (iss >> word)
        {
            reverse(word.begin(), word.end());
            result += word + " ";
        }
        result.pop_back();
        return result;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    string reverseWords(string s) {
        size_t first = 0, last = first;
        while (last != s.size())
        {
            if (s[last] == ' ')
            {
                reverse(&s[first], &s[last]);
                first = last + 1;
            }
            last++;
        }
        reverse(&s[first], &s[last]);
        return s;
    }
};
```
