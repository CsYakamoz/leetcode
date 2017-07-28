## [168. Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/#/description)

### Description

Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

```
    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 
```



**Difficult:** `Easy`

**Tags:** `Math`



### Solution One

此处引用 Top Solutions 的解释

[Python solution with explanation](https://discuss.leetcode.com/topic/6245/python-solution-with-explanation)

```c++
class Solution {
public:
    string convertToTitle(int n) {
        string res;
        string dict = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        do
        {
            n--;
            res.push_back(dict[n % 26]);
            n /= 26;
        } while (n);
        reverse(res.begin(), res.end());
        return res;
    }
};
```



