## [415. Add Strings](https://leetcode.com/problems/add-strings/#/description)

### Description

Given two non-negative integers `num1` and `num2` represented as string, return the sum of `num1` and `num2`.

**Note:**

1. The length of both `num1` and `num2` is < 5100.
2. Both `num1` and `num2` contains only digits `0-9`.
3. Both `num1` and `num2` does not contain any leading zero.
4. You **must not user any built-in BigInteger library** or **convert the inputs to integer** directly.

**Difficulty:** `Easy`

**Tags:** `Math`

### Solution One

```c++
class Solution {
public:
    string addStrings(string num1, string num2) {
        if (num1.empty() || num2.empty())
        {
            return num1 + num2;
        }
        string result;
        int sum = 0;
        int i = num1.size() - 1, j = num2.size() - 1;
        while (i >= 0 || j >= 0 || sum)
        {
            int oper1 = 0, oper2 = 0;
            if (i >= 0)
            {
                oper1 = num1[i--] - 48;
            }
            if (j >= 0)
            {
                oper2 = num2[j--] - 48;
            }
            sum += oper1 + oper2;
            result = (char)(sum % 10 + 48) + result;
            sum /= 10;
        }
        return result;
    }
};
```

### Solution Two - In Top Solutions

[3ms 5 lines Concise C++ Solution without extra space. The loop should stop early!](https://discuss.leetcode.com/topic/62281/3ms-5-lines-concise-c-solution-without-extra-space-the-loop-should-stop-early)

```c++
class Solution {
public:
    string addStrings(string num1, string num2) {
        if (num1.empty() || num2.empty())
        {
            return num1 + num2;
        }
        if (num2.size() > num1.size())
        {
            return addStrings(num2, num1);
        }
        int i = num1.size() - 1;
        int j = num2.size() - 1;
        int sum = 0;
        while (i >= 0 && (j >= 0 || sum))
        {
            sum += j >= 0 ? num2[j] - '0' : 0;
            sum += num1[i] - '0';
            num1[i] = sum % 10 + '0';
            sum /= 10;
            j--;
            i--;
        }
        return (sum ? "1" : "") + num1;
    }
};
```

### Solution Three

```c++
class Solution {
public:
    string addStrings(string num1, string num2) {
        if (num1.empty() || num2.empty())
        {
            return num1.empty() ? num2 : num1;
        }
        string res;
        int i = num1.size() - 1;
        int j = num2.size() - 1;
        int sum = 0;
        while (i >= 0 || j >= 0 || sum )
        {
            if (i >= 0)
            {
                sum += num1[i] - '0';
            }
            if (j >= 0)
            {
                sum += num2[j] - '0';
            }
            res += to_string(sum % 10);
            sum /= 10;
            i--;
            j--;
        }
        reverse(res.begin(), res.end());
        return res;
    }
};
```
