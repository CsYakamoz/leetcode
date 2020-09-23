## [53. Maximuim Subarray](https://leetcode.com/problems/maximum-subarray/description/)

### Description

Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array `[-2,1,-3,4,-1,2,1,-5,4]`,
the contiguous subarray `[4,-1,2,1]` has the largest sum = `6`.

**More practice:**

If you have figured out the O(*n*) solution, try coding another solution using the divide and conquer approach, which is more subtle.



**Difficult:** `Easy`

**Tags:** `Array` `Divide and Conquer` `Dynamic Programming`



### Solution One

`DP`

```c++
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int prev;
        int res = INT_MIN;
        for (size_t i = 0; i < nums.size(); i++)
        {
            if (i == 0)
            {
                res = prev = nums[i];
            }
            else
            {
                prev = max(prev + nums[i], nums[i]);
                res = max(res, prev);
            }
        }
        return res;
    }
};
```



### Solution Two - In Top Solutions

`DP`

```c++
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int max_sum = std::numeric_limits<int>::min();
        int sum = 0;
        
        for (auto n : nums) {
            sum = (sum + n) > n ? sum + n : n;
            max_sum = max(max_sum, sum);
        }
        return max_sum;
    }
};
```



### Solution Three - In Top Solutions

`DC`

[How to solve "Maximum Subarray" by using the divide and conquer approach ?](https://discuss.leetcode.com/topic/426/how-to-solve-maximum-subarray-by-using-the-divide-and-conquer-approach)

