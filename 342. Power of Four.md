## [342. Power of Four](https://leetcode.com/problems/power-of-four/description/)

### Description

Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

**Example:**
Given num = 16, return true. Given num = 5, return false.

**Follow up**: Could you solve it without loops/recursion?



**Difficult:** `Easy`

**Tags:** `Bit Manipulation`



### Solution One

`num & num - 1 == 0` 确保 32 位中，只有 1 位是 1，其余为 0

```c++
class Solution {
public:
    bool isPowerOfFour(int num) {
        if (num <= 0)
        {
            return false;
        }
        int a = num & 0x55555555;
        // num & num - 1 ensures there is just 1 bit
        // a ensure the 1 bit is in power of 4
        if ((num & num - 1) == 0 && a )
        {
            return true;
        }
        return false;
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    bool isPowerOfFour(int num) {
        if (num <= 0) return false;
        bitset<32> bits(num);
        if (bits.count() != 1) return false;
        for (int i=0; i<32; i+=2)
            if (bits[i]) return true;
        return false;
    }
};
```



### Solution Three - In Top Solutions

[1 line C++ solution without confusing bit manipulations](https://discuss.leetcode.com/topic/42914/1-line-c-solution-without-confusing-bit-manipulations)

> (4^n - 1) % 3 == 0
> another proof:
> (1) 4^n - 1 = (2^n + 1) * (2^n - 1)
> (2) among any 3 consecutive numbers, there must be one that is a multiple of 3
> among (2^n-1), (2^n), (2^n+1), one of them must be a multiple of 3, and (2^n) cannot be the one, therefore either (2^n-1) or (2^n+1) must be a multiple of 3, and 4^n-1 must be a multiple of 3 as well.
>
> @aesi123 from LeetCode

```c++
bool isPowerOfFour(int num) {
    return num > 0 && (num & (num - 1)) == 0 && (num - 1) % 3 == 0;
}
```



[Simple C++ O(1) solution without 0x55555555](https://discuss.leetcode.com/topic/44430/simple-c-o-1-solution-without-0x55555555)

```c++
class Solution {
public:
    bool isPowerOfFour(int num) {
        return ((num-1)&num)==0 && (num-1)%3==0;
    }
};
```



