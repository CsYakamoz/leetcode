## [202. Happy Number](https://leetcode.com/problems/happy-number/#/description)

### Description

Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

**Example: **19 is a happy number

- 1^2^ + 9^2^ = 82
- 8^2^ + 2^2^ = 68
- 6^2^ + 8^2^ = 100
- 1^2^ + 0^2^ + 0^2^ = 1

**Difficulty:** `Easy`

**Tags:** `Hash Table` `Math`

### Solution One

`Hash Table`

```c++
class Solution {
public:
    bool isHappy(int n) {
        set<int> s;
        while (n > 3)
        {	// 3 ^ 2 < 10
            if (s.find(n) == s.end())
            {
                s.insert(n);
                int sum = 0;
                while (n)
                {
                    int mod = n % 10;
                    sum += mod * mod;
                    n /= 10;
                }
                n = sum;
            }
            else
            {
                return false;
            }
        }
        return n == 1;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    /*
    Whether the number is or not, its square sum sequence is always a list with cycle
    类似于检测单向链表是否环路，使用双指针法（保证循环必然退出）
    */
    bool isHappy(int n) {
        int fast = n, slow = n;
        do{
            slow = getDigitSum(slow);
            fast = getDigitSum(fast);
            fast = getDigitSum(fast);
        } while (slow != fast);

        if (slow == 1)
            return true;
        return false;
    }

    inline int getDigitSum(int n)
    {
        int sum = 0;
        while (n)
        {
            sum += (n % 10) * (n % 10);
            n /= 10;
        }
        return sum;
    }
};
```
