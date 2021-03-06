## [9. Palindrome Number](https://leetcode.com/problems/palindrome-number/#/description)

### Description

Determine whether an integer is a palindrome. Do this without extra space.

**Difficulty:** `Easy`

**Tags:** `Math`

### Solution One

直接将数字反转，再判断是否相等

```c++
class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0)
        {
            return false;
        }
        long long result = 0;
        int temp = x;
        while (temp != 0)
        {
            result = result * 10 + temp % 10;
            temp /= 10;
        }
        return result == x;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    bool isPalindrome(int x) {
        if(x<0|| (x!=0 &&x%10==0)) return false;
          // 若最低位0，则反转后位数肯定不相同，所以直接返回false
        int sum=0;
        while(x>sum)
        {
            sum = sum*10+x%10;
            x = x/10;
        }
        return (x==sum)||(x==sum/10);
    }
};
```

### Solution Three

结合之前所做的题目*Reverse Integer*与思路二

```c++
class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0 || (x != 0 && x % 10 == 0) )
        {
            return false;
        }
        else
        {
            return x == reverse(x);
        }
    }
    int reverse(int x) {
        int ans = 0;
        while (x) {
            int temp = ans * 10 + x % 10;
            if (temp / 10 != ans)
                return 0;
            ans = temp;
            x /= 10;
        }
        return ans;
    }
};
```
