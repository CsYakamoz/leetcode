## [405. Convert a Number to Hexadecimal](https://leetcode.com/problems/convert-a-number-to-hexadecimal/tabs/description)

### Description

Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, [two’s complement](https://en.wikipedia.org/wiki/Two%27s_complement) method is used.

**Note:**

1. All letters in hexadecimal (`a-f`) must be in lowercase.
2. The hexadecimal string must not contain extra leading `0`s. If the number is zero, it is represented by a single zero character `'0'`; otherwise, the first character in the hexadecimal string will not be the zero character.
3. The given number is guaranteed to fit within the range of a 32-bit signed integer.
4. You **must not use any method provided by the library** which converts/formats the number to hex directly.

**Example 1:**

```
Input:
26

Output:
"1a"

```

**Example 2:**

```
Input:
-1

Output:
"ffffffff"
```

**Difficulty:** `Easy`

**Tags:** `Bit Manipulation`

### Solution One

一开始，想用正常的方法——除以 16，得余数；但发现如果是负数，则会出现问题

后来看了题目类型，发现是位操作，于是想到每 4 位二进制等于 1 位十六进制，便取**每 4 位二进制来当作余数**

然后发现在 VS 2017 下（LeetCode 编译器未知），对负数进行右移，最高位保留符号位，此时只好使用`unsigned`类型

**注意：**此处不能使用`size_t`类型，`size_t`在 LeetCode 编译器中为 8 个字节，而不是 4 个字节

```c++
class Solution {
public:
    string toHex(int num) {
        unsigned x = num;
        string res;
        do
        {
            int mod = 0;
            int base = 1;
            for (size_t i = 0; x && i < 4; i++)
            {
                if (x & 1)
                {
                    mod += base;
                }
                x >>= 1;
                base *= 2;
            }
            if (mod > 9)
            {
                res.push_back(mod - 10 + 'a');
            }
            else
            {
                res.push_back(mod + '0');
            }
        } while (x);
        reverse(res.begin(), res.end());
        return res;
    }
};
```

### Solution Two - In Top Solutions

```c++
const string HEX = "0123456789abcdef";
class Solution {
public:
    string toHex(int num) {
        if (num == 0) return "0";
        string result;
        int count = 0;
        while (num && count++ < 8) {
            result = HEX[(num & 0xf)] + result;
            num >>= 4;
        }
        return result;
    }
};
```

### Solution Three - In Top Solutions

13 行的原因，目前还不懂

```c++
class Solution {
public:
    string toHex(int num) {
        if (num == 0)
        {
            return "0";
        }
        vector<char> unit = { '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f' };
        string res;
        long long number = num;
        if (number < 0)
        {
            number = number + 1 + 4294967295;
        }
        while (number != 0)
        {
            res.push_back(unit[number % 16]);
            number = number / 16;
        }
        int n = res.length();
        for (int i = 0; i < n / 2; i++)
        {
            char tmp;
            tmp = res[i];
            res[i] = res[n - i - 1];
            res[n - i - 1] = tmp;
        }
        return res;
    }
};
```
