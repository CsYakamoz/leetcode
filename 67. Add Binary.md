## [67. Add Binary](https://leetcode.com/problems/add-binary/description/)

### Description

Given two binary strings, return their sum (also a binary string).

For example,
a = `"11"`
b = `"1"`
Return `"100"`.



**Difficult:** `Easy`

**Tags:** `Math` `String`



### Solution One

```c++
class Solution {
public:
    string addBinary(string a, string b) {
        if (a.size() < b.size())
        {   // Need to ensure that the length of a is greater than or equal to b
            return addBinary(b, a);
        }
        int sum = 0;
        int i = a.size() - 1;
        int j = b.size() - 1;
        while (i >= 0 && (j >= 0 || sum))
        {
            sum += a[i] - '0';
            if (j >= 0) sum += b[j--] - '0';
            a[i] = sum % 2 + '0';
            sum /= 2;
            i--;
        }
        if (sum)
        {
            a.insert(a.begin(), '1');
        }
        return a;
    }
};
```



