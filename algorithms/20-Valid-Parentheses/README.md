## [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/#/description)

### Description

Given a string containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

The brackets must close in the correct order, `"()"` and `"()[]{}"` are all valid but `"(]"` and `"([)]"` are not.

**Difficult:** `Easy`

**Tags:** `Stack` `String`

### Solution One

`Stack`

遇到左括号，就将其添加到`v`中

遇到右括号，就判断`v`中是否有左括号且左括号与右括号匹配

如果判断结果为`true`，则继续，否则返回`false`

在循环结束后，还要判断`v`是否还有左括号，因为可能左括号的数量较多

```c++
class Solution {
public:
    bool isValid(string s) {
        set<char> map = { '(', '[', '{' };
        vector<char> v;
        char c;
        int index = -1;
        for (size_t i = 0; i < s.size(); i++)
        {
            if (map.find(s[i]) != map.end())
            {
                v.push_back(s[i]);
                ++index;
            }
            else
            {
                switch (s[i])
                {
                case ')':
                    c = '(';
                    break;
                case ']':
                    c = '[';
                    break;
                case '}':
                    c = '{';
                    break;
                }
                if (index >= 0 && c == v[index])
                {
                    v.pop_back();
                    --index;
                }
                else
                {
                    return false;
                }
            }
        }
        if (index >= 0)
        {
            return false;
        }
        else
        {
            return true;
        }
    }
};
```
