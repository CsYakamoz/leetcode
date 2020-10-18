## [537. Complex Number Multiplication](https://leetcode.com/problems/complex-number-multiplication/description/)

### Description

Given two strings representing two [complex numbers](https://en.wikipedia.org/wiki/Complex_number).

You need to return a string representing their multiplication. Note i2 = -1 according to the definition.

**Example 1:**

```
Input: "1+1i", "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.

```

**Example 2:**

```
Input: "1+-1i", "1+-1i"
Output: "0+-2i"
Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
```

**Note:**

1. The input strings will not have extra blank.
2. The input strings will be given in the form of **a+bi**, where the integer **a** and **b** will both belong to the range of [-100, 100]. And **the output should be also in this form**.

**Difficulty:** `Medium`

**Tags:** `Math` `String`

### Solution One

```c++
class Solution {
public:
    string complexNumberMultiply(string a, string b) {
        complex c1 = constructComplex(a);
        complex c2 = constructComplex(b);
        int x = c1.a * c2.a - c1.b * c2.b;
        int y = c1.a * c2.b + c1.b * c2.a;
        return to_string(x) + '+' + to_string(y) + 'i';
    }

private:
    struct complex
    {
        int a;
        int b;
        complex(int i, int j) : a(i), b(j) {}
    };

    complex constructComplex(string s)
    {
        auto oper = s.find('+');
        int i = stoi(s.substr(0, oper));
        int j = stoi(s.substr(oper + 1, s.size() - oper - 1));
        return complex(i, j);
    }
};
```

### Solution Two - In Top Solutions

[c++ using stringstream](https://discuss.leetcode.com/topic/84382/c-using-stringstream)

> stringstream is very useful to extract num from a string

```c++
class Solution {
public:
    string complexNumberMultiply(string a, string b) {
        int ra, ia, rb, ib;
        char buff;
        stringstream aa(a), bb(b), ans;
        aa >> ra >> buff >> ia >> buff;
        bb >> rb >> buff >> ib >> buff;
        ans << ra*rb - ia*ib << "+" << ra*ib + rb*ia << "i";
        return ans.str();
    }
};
```
