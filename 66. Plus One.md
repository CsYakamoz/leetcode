## [66. Plus One](https://leetcode.com/problems/plus-one/tabs/description)

### Description

Given a non-negative integer represented as a **non-empty** array of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.



**Difficult:** `Easy`

**Tags:** `Array` `Math`



### Solution One

```c++
class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        bool carry = true;
        for (int i = digits.size() - 1; carry && i >= 0; i--)
        {
            digits[i] += 1;
            carry = digits[i] / 10;
            digits[i] %= 10;
        }
        if (carry)
        {
            digits.insert(digits.begin(), 1);
        }
        return digits;
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        int lastIndex = digits.size()-1;
        while(lastIndex >= 0){
            if(digits[lastIndex] + 1 >= 10){   
                digits[lastIndex] = 0;
                lastIndex--;
            }
            else{
                digits[lastIndex] += 1;
                break;
            }
        }
        if(lastIndex == -1) digits.insert(digits.begin(), 1);
        return digits;
            
    }
};
```



