## [268. Missing Number](https://leetcode.com/problems/missing-number/#/description)

### Description

Given an array containing _n_ distinct numbers taken from `0, 1, 2, ..., n`, find the one that is missing from the array.

For example,
Given _nums_ = `[0, 1, 3]` return `2`.

**Note**:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

**Difficult:** `Easy`

**Tags:** `Array` `Math` `Bit Manipulation`

### Solution One

`Math`

通过等差数列求和公式，获得 `0 + 1 + ... + n`的值

再减去`nums`中的值，便得到所要的结果

```c++
class Solution {
public:
    int missingNumber(vector<int>& nums) {
        // formula = (a1 + an) * n / 2;
        // res = 0 + 1 + ... + n
        // In this case, a1 is 0, an is n, the length is n + 1
        int res = nums.size() * (nums.size() + 1) / 2;
        for (const auto &i : nums)
            res -= i;
        return res;
    }
};
```

### Solution Two

|  n  | 0 ^ ,,,,,, ^ n  | result |
| :-: | :-------------: | :----: |
|  0  | 0 ^ ....... ^ 0 |   0    |
|  1  | 0 ^ ...... ^ 1  |   1    |
|  2  | 0 ^ ...... ^ 2  |   3    |
|  3  | 0 ^ ...... ^ 3  |   0    |
|  4  | 0 ^ ...... ^ 4  |   4    |
|  5  | 0 ^ ...... ^ 5  |   1    |
|  6  | 0 ^ ...... ^ 6  |   7    |
|  7  | 0 ^ ...... ^ 7  |   0    |

这里只列出 0 ~ 7 ，若想看得更多，可以写个循环看

通过上面可以发现规律，四个四个为一组，`4k + j( 0 <= j <= 3)`

- j = 0，则 result = n
- j = 1，则 result = 1
- j = 2，则 result = n + 1
- j = 3，则 result = 0

```c++
class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int res;
        switch (nums.size() % 4)
        {
        case 0:
            res = nums.size();
            break;
        case 1:
            res = 1;
            break;
        case 2:
            res = nums.size() + 1;
            break;
        case 3:
            res = 0;
            break;
        }
        for (const auto i : nums)
            res ^= i;
        return res;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    int missingNumber(vector<int>& nums) {
      int result = nums.size();
      for (int i=0;i<nums.size();i++){
          result ^= i;
          result ^= nums[i];
      }
        return result;
    }
};
```
