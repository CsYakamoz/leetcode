## [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/description/)

### Description

Given *n* pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given *n* = 3, a solution set is:

```
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```



**Difficult:** `Medium`

**Tags:** `String` `Backtracking`



### Solution One

```c++
class Solution {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> res;
        string str;
        dfs(n, n, str, res);
        return res;
    }

private:
    void dfs(int left, int right, string &str, vector<string> &res) {
        if (left == 0 && right == 0) {
            res.push_back(str);
            return;
        }

        if (left > 0) {
            str.push_back('(');
            dfs(left - 1, right, str, res);
            str.pop_back();
        }

        if (right > left) {
            str.push_back(')');
            dfs(left, right - 1, str, res);
            str.pop_back();
        }
    }
};
```



