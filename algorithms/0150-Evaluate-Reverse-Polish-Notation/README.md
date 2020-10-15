## [150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/description/)

### Description

Evaluate the value of an arithmetic expression in [Reverse Polish Notation](http://en.wikipedia.org/wiki/Reverse_Polish_notation).

Valid operators are `+`, `-`, `*`, `/`. Each operand may be an integer or another expression.

Some examples:

```
  ["2", "1", "+", "3", "*"] -> ((2 + 1) * 3) -> 9
  ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6
```

**Difficult:** `Medium`

**Tags:** `Stack`

### Solution One

```c++
class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        string oper = "+-*/";
        stack<double> operand;
        for (string &s : tokens)
        {
            if (oper.find(s) != string::npos)
            {
                int val1 = operand.top();
                operand.pop();
                int val2 = operand.top();
                operand.pop();
                switch (s[0])
                {
                case '+':
                    operand.push(val2 + val1);
                    break;
                case '-':
                    operand.push(val2 - val1);
                    break;
                case '*':
                    operand.push(val2 * val1);
                    break;
                case '/':
                    operand.push(val2 / val1);
                    break;
                }
            }
            else
            {
                operand.push(stod(s));
            }
        }
        return operand.top();
    }
};
```
