## [171. Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/#/description)

### Description

Related to question [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/)

Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

```
    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
```

**Difficulty:** `Easy`

**Tags:** `Math`

### Solution One

```c++
class Solution {
public:
    int titleToNumber(string s) {
        int res = 0;
        for (auto i : s)
        {
            res = res * 26 + i - 'A' + 1;
        }
        return res;
    }
};
```
