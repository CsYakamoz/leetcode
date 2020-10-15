## [507. Perfect Number](https://leetcode.com/problems/perfect-number/description/)

### Description

We define the Perfect Number is a **positive** integer that is equal to the sum of all its **positive** divisors except itself.

Now, given an

integer

n, write a function that returns true when it is a perfect number and false when it is not.

**Example:**

```
Input: 28
Output: True
Explanation: 28 = 1 + 2 + 4 + 7 + 14

```

**Note:** The input number **n** will not exceed 100,000,000. (1e8)

**Difficult:** `Easy`

**Tags:** `Math`

### Solution One

根据题目要求且 1 不是完全数，所以一开始判断`num`是否小于等于 1

由于 1 肯定是因子，所以`sum`被初始化为 1

一开始 `for`循环的条件是写成`i * i <= num`，但不知道为什么不行（根据**Solution**，是可以的，估计自己哪里出错了）

所以改成使用`ending`，意思是刚刚得到的较大因子，1 和 28，28 赋值给`ending`， 2 和 14， 14 赋值给`ending`，4 和 7， 7 赋值给`ending`

同时`ending`是为了避免添加同样的因子，例如刚才的 14 和 7

最后一个`if`是也是为了避免添加同样的因子

```c++
class Solution {
public:
    bool checkPerfectNumber(int num) {
        if (num <= 1)
        {
            return false;
        }
        int sum = 1;    // 1 is positive divisor
        int ending = num;
        // If we initialize i as 1, the num will be added to sum
        for (int i = 2; i < ending; i++)
        {
            if (num % i == 0)
            {
                sum += i + num / i;
                ending = num / i;
                // Avoid adding (num / i) repeatedly, but it is unuseful if i * i == num
            }
        }
        int sqrt_x = sqrt(num);
        if (sqrt_x * sqrt_x == num)
        {
            sum -= sqrt_x;
        }
        return sum == num;
    }
};
```

### Solution Two - In Top Solutions

[Simple Java Solution](https://discuss.leetcode.com/topic/84259/simple-java-solution)

这里不考虑 `i == num / i`的原因，见上面链接

```c++
public class Solution {
    public boolean checkPerfectNumber(int num) {
        if (num == 1) return false;

        int sum = 0;
        for (int i = 2; i <= Math.sqrt(num); i++) {
            if (num % i == 0) {
                sum += i + num / i;
            }
        }
        sum++;

        return sum == num;
    }
}
```

### Solution Three - In Top Solutions

根据维基百科，在`int`所能表示的范围中，只有下列 5 个是完全数

```c++
class Solution {
public:
    bool checkPerfectNumber(int num) {
        static unordered_set<int> n = {6, 28, 496, 8128, 33550336};
        return n.count(num);
    }
};
```

### Other Solutions

[507. Perfect Number - Solution](https://leetcode.com/problems/perfect-number/solution/)
