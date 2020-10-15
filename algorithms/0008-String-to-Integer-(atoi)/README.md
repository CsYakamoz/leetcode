## [8. String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/)

### Description

Implement atoi to convert a string to an integer.

**Hint:** Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

**Notes:** It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

**Difficult:** `Medium`

**Tags:** `Math` `String`

### Solution One

第一个循环首先将查看数字（包括符号）前面是否有非数字（不包括空格)，有则直接返回 0

第二个循环跳过空格

第三个循环判断有多少个符号，大于一个则直接返回 0

第三个循环将数字添加到`result`，并且判断是否大于`INT_MAX`或者小于`INT_MIN`

```c++
class Solution {
public:
    int myAtoi(string str) {
        if (str.size() == 0)
        {
            return 0;
        }
        int signNum = 0;
        int sign = 1;
        string::size_type i = 0;
        long long result = 0;
        set<char> dict = { '-', '+', '0', '1', '2', '3', '4', '5','6', '7','8','9', ' ' };
        while (i < str.size())
        {
            if (dict.find(str[i]) == dict.end())
            {
                return 0;
            }
            else if (dict.find(str[i]) != dict.end())
            {
                break;
            }
            ++i;
        }
        while (i < str.size() && str[i] == ' ')
        {
            i++;
        }
        while (i < str.size() && (str[i] == '-' || str[i] == '+'))
        {
            if (str[i] == '-')
            {
                sign = -1;
                ++signNum;
            }
            else
            {
                ++signNum;
            }
            ++i;
        }
        if (signNum > 1)
        {
            return 0;
        }
        while (i < str.size())
        {
            if (str[i] < '0' || str[i] > '9')
            {
                break;
            }
            result = result * 10 + (str[i] - 48) * sign;
            if (result > INT_MAX)
            {
                return INT_MAX;
            }
            else if (result < INT_MIN)
            {
                return INT_MIN;
            }
            ++i;
        }
        return result;
    }
};
```

### Solution Two - In Top Solutions

与**One**类似，只是少了第一个循环和第三个循环中符号数是否大于 1

这里跳过空格后，直接查看是否为符号，接着直接数字添加到`result`

```c++
class Solution {
public:
    int myAtoi(string str) {
        if (str.size() == 0)
        {
            return 0;
        }
        int sign = 1;
        string::size_type i = 0;
        long long result = 0;
        while (i < str.size() && str[i] == ' ')
        {
            i++;
        }
        if (i < str.size() && (str[i] == '-' || str[i] == '+'))
        {
            if (str[i] == '-')
            {
                sign = -1;
            }
            ++i;
        }
        while (i < str.size())
        {
            if (str[i] < '0' || str[i] > '9')
            {
                break;
            }
            result = result * 10 + (str[i] - 48) * sign;
            if (result > INT_MAX)
            {
                return INT_MAX;
            }
            else if (result < INT_MIN)
            {
                return INT_MIN;
            }
            ++i;
        }
        return result;
    }
};
```
