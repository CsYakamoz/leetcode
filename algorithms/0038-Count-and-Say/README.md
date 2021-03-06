## [38. Count and Say](https://leetcode.com/problems/count-and-say/#/description)

### Description

The count-and-say sequence is the sequence of integers beginning as follows:
`1, 11, 21, 1211, 111221, ...`

`1` is read off as `"one 1"` or `11`.
`11` is read off as `"two 1s"` or `21`.
`21` is read off as `"one 2`, then `one 1"` or `1211`.

Given an integer _n_, generate the $n^{th}$ sequence.

Note: The sequence of integers will be represented as a string.

**Difficulty:** `Easy`

**Tags:** `String`

### Solution One

计算有多少个`1`或`2`，接着添加到`tmp`

```c++
class Solution {
public:
    string countAndSay(int n) {
        string s = "1";
        if (n < 2)
        {
            return s;
        }
        for (int i = 1; i < n; i++)
        {
            string tmp;
            size_t j = 1;
            auto c = s[0];
            int count = 1;
            while (j < s.size())
            {
                if (c == s[j])
                {
                    count++;
                }
                else
                {
                    tmp += to_string(count) + c;
                    c = s[j];
                    count = 1;
                }
                j++;
            }
            tmp += to_string(count) + c;
            s = tmp;
        }
        return s;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    string countAndSay(int n) {
        if (n <= 0) {
            return string();
        }
        string output = string("1");
        while ((n-1) > 0) {
            output = nextSequence(output);
            n--;
        }
        return output;
    }

    string nextSequence(string& input) {
        int count;
        char currentChar = 0;
        string output;
        for (string::iterator it = input.begin(); it != input.end(); it++) {
            if (currentChar == 0) {
                currentChar = *it;
                count = 1;
            } else {
                if (*it == currentChar) {
                    count++;
                } else {
                    output += ('0' + count);
                    output += (currentChar);
                    currentChar = *it;
                    count = 1;
                }
            }
        }
        // cout << "count is " << count << "current char is " << (currentChar - '0');
        output += ('0' + count);
        output += (currentChar);
        return output;
    }
};
```
