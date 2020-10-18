## [412. Fizz Buzz](https://leetcode.com/problems/fizz-buzz/#/description)

### Description

Write a program that outputs the string representation of numbers from 1 to _n_.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

**Example:**

```
n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
```

**Difficulty:** `Easy`

**Tags:**

### Solution One

```c++
class Solution {
public:
    vector<string> fizzBuzz(int n) {
        vector<string> result;
        for (int i = 1; i <= n; i++)
        {
            string s;
            if (i % 15 == 0)
            {
                s = "FizzBuzz";
            }
            else if (i % 5 == 0)
            {
                s = "Buzz";
            }
            else if (i % 3 == 0)
            {
                s = "Fizz";
            }
            else
            {
                s = to_string(i);
            }
            result.push_back(s);
        }
        return result;
    }
};
```

### Solution Two

```c++
class Solution {
public:
    vector<string> fizzBuzz(int n) {
        vector<string> result(n);
        vector<bool> isVisted(n);
        for (size_t i = 15; i <= n; i += 15) {
            if (!isVisted[i - 1]) {
                result[i - 1] = "FizzBuzz";
                isVisted[i - 1] = true;
            }
        }
        for (size_t i = 3; i <= n; i += 3) {
            if (!isVisted[i - 1]) {
                result[i - 1] = "Fizz";
                isVisted[i - 1] = true;
            }
        }
        for (size_t i = 5; i <= n; i += 5) {
            if (!isVisted[i - 1]) {
                result[i - 1] = "Buzz";
                isVisted[i - 1] = true;
            }
        }
        for (size_t i = 1; i <= n; i++)
        {
            if (!isVisted[i-1])
            {
                result[i - 1] = to_string(i);
            }
        }
        return result;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    vector<string> fizzBuzz(int n) {
        vector<string> result(n);
        for (int i = 1; i <= n; i++)
        {
            if (i % 3 == 0)
            {
                result[i - 1] += "Fizz";
            }
            if (i % 5 == 0)
            {
                result[i - 1] += "Buzz";
            }
            if (result[i-1] == "")
            {
                result[i - 1] += to_string(i);
            }
        }
        return result;
    }
};
```
