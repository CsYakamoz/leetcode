## [273. Integer to English Words](https://leetcode.com/problems/integer-to-english-words/description/)

### Description

Convert a non-negative integer to its english words representation. Given input is guaranteed to be less than 2^31^ - 1.

For example,

```
123 -> "One Hundred Twenty Three"
12345 -> "Twelve Thousand Three Hundred Forty Five"
1234567 -> "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
```



**Difficult:**  `Hard`

**Tags:**  `Math`  `String`



### Solution One

```c++
class Solution {
public:
    string numberToWords(int num)
    {
        if (num == 0) return "Zero";

        vector<string> record;
        for (size_t i = 0; num != 0 && i < 4; i++)
        {
            int n = get_3_digits(num);
            if (n != 0)
                record.push_back(handle_3_digits(n) + a[i]);
        }
        string result;
        for (int i = record.size() - 1; i >= 0; i--)
        {
            result += record[i];
            if (i != 0) result.push_back(' ');
        }
        return result;
    }

private:
    string handle_3_digits(int num)
    {
        string res;
        if (num > 99) {
            res += b[num / 100] + " Hundred";
            num = num % 100;
            if (num != 0) res.push_back(' ');
        }
        if (num <= 19) {
            res += c[num];
        } else {
            res += d[num / 10];
            num = num % 10;
            if (num != 0) {
                res.push_back(' ');
                res += b[num];
            }
        }
        return res;
    }

    int get_3_digits(int &num)
    {
        string str;
        for (size_t i = 0; num != 0 && i < 3; i++) {
            str.push_back(num % 10 + '0');
            num /= 10;
        }
        reverse(str.begin(), str.end()); 
        return stoi(str);
    }

    vector<string> a{ "", " Thousand", " Million", " Billion" };
    vector<string> b{ "", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine" };
    vector<string> c{ "", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
            "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen" };
    vector<string> d{ "", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety" };
};
```



### Solution Two - In Top Solutions

[My clean Java solution, very easy to understand](https://discuss.leetcode.com/topic/23054/my-clean-java-solution-very-easy-to-understand)

```python
class Solution(object):

    less_than_20 = ["", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
            "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"]
    tens = ["", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"]
    thousands = ["", "Thousand", "Million", "Billion"]

    def numberToWords(self, num):
        """
        :type num: int
        :rtype: str
        """
        if num == 0:
            return "Zero"
        words = ""
        i = 0
        while num != 0:
            if num % 1000 != 0:
                words = self.helper(num % 1000) + Solution.thousands[i] + ' ' + words
            num //= 1000
            i = i + 1

        return words.rstrip()


    def helper(self, num):
        if num == 0:
            return ''
        elif num < 20:
            return Solution.less_than_20[num] + ' '
        elif num < 100:
            return Solution.tens[num // 10] + ' ' + self.helper(num % 10)
        else:
            return Solution.less_than_20[num // 100] + " Hundred " + self.helper(num % 100)
```



