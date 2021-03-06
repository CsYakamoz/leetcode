## [6. ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/#/description)

### Description

The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

```c++
string convert(string text, int nRows);
```

`convert("PAYPALISHIRING", 3)` should return `"PAHNAPLSIIGYIR"`.

**Difficulty:** `Medium`

**Tags:** `String`

### Solution One

将字符串 s 的字符按顺序($C_0,C_1,\cdots,C_n$)添加到相应行中

使用变量`dir`判断往上走还是往下走

```c++
class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows == 1)
        {
            return s;
        }
        vector<vector<char>> Row(numRows);
        string::size_type i = 0;
        string::size_type line = 0;
        bool dir = true;
        while (i != s.size())
        {
            if (dir)
            {
                Row[line].push_back(s[i++]);
                if (line == numRows - 1)
                {
                    line = numRows - 2;
                    dir = false;
                    continue;
                }
                ++line;
            }
            else
            {
                Row[line].push_back(s[i++]);
                if (line == 0)
                {
                    line = 1;
                    dir = true;
                    continue;
                }
                --line;
            }
        }
        string result;
        for (auto i : Row)
            for (auto j : i)
                result += j;
        return result;
    }
};
```

### Solution Two

首先计算出每一行中两个字符下标的差的规律`left = (numRows - line - 1) * 2`和`right = line * 2`

然后根据`left`和`right`的值，判断进行哪个循环（`left`和`right`相同或者`left`和`right`其中一个为 0，则下标差固定为`max(left, right)`，否则需要使下标差为`left`或者`right`）

两种循环都将该行应有的字符添加到`result`

```c++
class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows == 1)
        {
            return s;
        }
        string result;
        string::size_type line = 0;
        string::size_type left, right;
        string::size_type i;
        while (line != numRows)
        {
            i = line;
            left = (numRows - line - 1) * 2;
            right = line * 2;
            if (left == right || left == 0 || right == 0)
            {
                left = left > right ? left : right;
                while (i < s.size())
                {
                    result += s[i];
                    i += left;
                }
            }
            else
            {
                string::size_type temp[2] = { left, right };
                size_t j = 0;
                while (i < s.size())
                {
                    result += s[i];
                    i += temp[j];
                    j = (j + 1) % 2;
                }
            }
            ++line;
        }
        return result;
    }
};
```
