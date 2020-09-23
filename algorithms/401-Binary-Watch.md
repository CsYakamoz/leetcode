## [401. Binary Watch](https://leetcode.com/problems/binary-watch/description/)

### Description

A binary watch has 4 LEDs on the top which represent the **hours** (**0-11**), and the 6 LEDs on the bottom represent the **minutes** (**0-59**).

Each LED represents a zero or one, with the least significant bit on the right.

![](https://upload.wikimedia.org/wikipedia/commons/8/8b/Binary_clock_samui_moon.jpg)

For example, the above binary watch reads "3:25".

Given a non-negative integer *n* which represents the number of LEDs that are currently on, return all possible times the watch could represent.

**Example:**

```
Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
```

**Note:**

- The order of output does not matter.
- The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
- The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".



**Difficult:** `Easy`

**Tags:** `Backtracking` `Bit Manipulation`



### Solution One

```c++
class Solution
{
public:
    vector<string> readBinaryWatch(int num)
    {
        if (num == 0)
            return { {"0:00"} };

        vector<string> res;
        pair<int, int> times;


        helper(0, num, times, res);
        return res;
    }

private:
    const vector<int> LED{ 1,2,4,8,1,2,4,8,16,32 };

    void helper(int i, int num, pair<int, int> &times, vector<string> &res)
    {
        if (num == 0)
        {
            res.push_back(to_string(times.first) + ':'
                + (times.second > 9 ? to_string(times.second) : '0' + to_string(times.second)));
            return;
        }
        for (; i < LED.size(); i++)
        {
            if (i < 4)
            {
                if (times.first + LED[i] > 11)
                    continue;
                times.first += LED[i];
                helper(i + 1, num - 1, times, res);
                times.first -= LED[i];
            }
            else
            {
                if (times.second + LED[i] > 59)
                    continue;
                times.second += LED[i];
                helper(i + 1, num - 1, times, res);
                times.second -= LED[i];
            }
        }
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
int countBit(int num)
{
    int count = 0;
    for (int i=0; i<7; i++)
        if (num&(1<<i)) count++;
    
    return count;
}
    
public:
    vector<string> readBinaryWatch(int num) {
        vector<string> rs;
        for (int h = 0; h < 12; h++)
        {
            int hbit = countBit(h);
            for (int m = 0; m < 60; m++)
            {
                if (hbit+countBit(m) == num)
                {
                    ostringstream ss;
                    ss << h << ":" << (m<10?"0":"") << m;                     
                    rs.push_back(ss.str());
                }
            }
        }
        
        return rs;
    }
};
```



