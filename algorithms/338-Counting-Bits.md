## [338. Counting Bits](https://leetcode.com/problems/counting-bits/#/description)

### Description

Given a non negative integer number **num**. For every numbers **i** in the range **0 ≤ i ≤ num** calculate the number of 1's in their binary representation and return them as an array.

**Example:**
For `num = 5` you should return `[0,1,1,2,1,2]`.

**Follow up:**

- It is very easy to come up with a solution with run time **O(n\*sizeof(integer))**. But can you do it in linear time **O(n)** /possibly in a single pass?
- Space complexity should be **O(n)**.
- Can you do it like a boss? Do it without using any builtin function like **__builtin_popcount** in c++ or in any other language.

**Hint:**

1. You should make use of what you have produced already.
2. Divide the numbers in ranges like [2-3], [4-7], [8-15] and so on. And try to generate new range from previous.
3. Or does the odd/even status of the number help you in calculating the number of 1s?



**Difficult:** `Medium`

**Tags:** `Dynamic Programming` `Bit Manipulation`



### Solution One

对`[0, num]`区间的每个数都计算其有多少个1

```c++
class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> result(num + 1);
        for (int i = 0; i < num+1; i++)
        {
            int tmp = i;
            int count = 0;
            while (tmp)
            {
                if (tmp & 1)
                {
                    ++count;
                }
                tmp >>= 1;
            }
            result[i] = count;
        }
        return result;
    }
};
```



### Solution Two

```c++
class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> result(num + 1);
        int i = 1;
        int sub = 2;
        if (i < num + 1)
        {
            result[i] = i;
            i++;
        }
        while (i < num + 1)
        {
            int count = sub;
            while (i < num + 1 && count)
            {
                result[i] = result[i - sub] + 1;
                --count;
                ++i;
            }
            sub *= 2;
        }
        return result;
    }
};
```



### Solution Three

```c++
class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> result(num + 1);
        int i = 0;
        while (i < num + 1)
        {
            if (i % 2)
            {
                result[i] = result[i / 2] + 1;
            }
            else
            {
                result[i] = result[i / 2];
            }
            ++i;
        }
        return result;
    }
};
```



### Solution Four - In Top Solutions

```c++
class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> result(num + 1);
        int i = 0;
        while (i < num + 1)
        {
            result[i] = result[i / 2] + i % 2;
            ++i;
        }
        return result;
    }
};
```


