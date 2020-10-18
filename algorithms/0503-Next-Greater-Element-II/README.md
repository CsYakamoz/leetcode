## [503. Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/#/description)

### Description

Given a circular array (the next element of the last element is the first element of the array), print the Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.

**Example 1:**

```
Input: [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number;
The second 1's next greater number needs to search circularly, which is also 2.

```

**Note:** The length of given array won't exceed 10000.

**Difficulty:** `Medium`

**Tags:** `Stack`

### Solution One

原始办法，对于第`i`个元素，从`i+1`位置开始寻找第一个比`nums[i]`大的元素

```c++
class Solution {
public:
    vector<int> nextGreaterElements(vector<int>& nums) {
        vector<int> result(nums.size());
        for (size_t i = 0; i < nums.size(); i++)
        {
            auto j = (i + 1) % nums.size();
            int val = -1;
            while (j != i)
            {
                if (nums[j] > nums[i])
                {
                    val = nums[j];
                    break;
                }
                j = (j + 1) % nums.size();
            }
            result[i] = val;
        }
        return result;
    }
};
```

### Solution Two - In Top Solutions

思路与 _Next Greater Element I_ 类似

```c++
class Solution {
public:
    vector<int> nextGreaterElements(vector<int>& nums) {
        vector<int> result(nums.size(), -1);
        stack<int> s;
        for (size_t i = 0; i < 2 * nums.size(); i++)
        {
            int val = nums[i % nums.size()];
            while (!s.empty() && nums[s.top()] < val)
            {
                result[s.top()] = val;
                s.pop();
            }
            if (i < nums.size())
            {
                s.push(i);
            }
        }
        return result;
    }
};
```
