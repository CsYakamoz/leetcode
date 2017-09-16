## [670. Maximum Swap](https://leetcode.com/problems/maximum-swap/description/)

### Description

Given a non-negative integer, you could swap two digits **at most** once to get the maximum valued number. Return the maximum valued number you could get.

**Example 1:**

```
Input: 2736
Output: 7236
Explanation: Swap the number 2 and the number 7.

```

**Example 2:**

```
Input: 9973
Output: 9973
Explanation: No swap.

```

**Note:**

1. The given number is in the range [0, 108]



**Difficult:** `Medium`

**Tags:** `Array` `Math`



### Solution One

```c++
class Solution {
public:
    int maximumSwap(int num) {
        if (num < 10)
            return num;

        vector<int> digits;
        int i = 0;
        do
        {
            int mod = num % 10;
            digits.push_back(mod);
            num /= 10;
        } while (num != 0);

        vector<int> tmp = digits;
        sort(tmp.begin(), tmp.end());

        for (int i = digits.size() - 1; i >= 0; i--)
        {
            if (tmp[i] != digits[i])
            {
                for (int j = 0; j < i; j++)
                {
                    if (digits[j] == tmp[i])
                    {
                        swap(digits[j], digits[i]);
                        break;
                    }
                }
                break;
            }
        }

        int res = 0;
        for (int i = digits.size() - 1; i >= 0; i--)
            res = res * 10 + digits[i];

        return res;
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int maximumSwap(int num) {
        string s = to_string(num);
        int len = s.size();
        for(int i = 0; i < len; ++i){
            vector<int> rec;
            for(int j = i+1; j < len; ++j){
                if( (rec.empty() || s[j] >= s[rec.back()]) && s[j] > s[i]) rec.push_back(j);
            }
            if(rec.size() > 0) {
                swap(s[i], s[rec.back()]);
                break;
            }
        }
        return stoi(s);
    }
};
```



