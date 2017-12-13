## [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/description/)

### Description

Your are given an array of positive integers `nums`.

Count and print the number of (contiguous) subarrays where the product of all the elements in the subarray is less than `k`.

**Example 1:**

```
Input: nums = [10, 5, 2, 6], k = 100
Output: 8
Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

```

**Note:**

* `0 < nums.length <= 50000`.
* `0 < nums[i] < 1000`.
* `0 <= k < 10^6`.



**Difficult:** `Medium`

**Tags:** `Array` `Two Pointers`



### Solution One

```c++
class Solution
{
public:
    int numSubarrayProductLessThanK(std::vector<int> &nums, int k)
    {
        if (k <= 1)
            return 0;

        long long product = 1;
        int count = 0;
        int begin = 0;
        int end = 0;

        while (end < nums.size())
        {
            product *= nums[end];

            if (product >= k)
            {
                count += sum(end - begin);

                while (begin <= end && product >= k)
                    product /= nums[begin++];

                // Remove Duplicate Subarray
                count -= sum(end - begin);
            }

            end++;
        }

        count += sum(end - begin);
        return count;
    }

private:
    int sum(int len)
    {
        if (len <= 0)
            return 0;

        int res = 1;
        if (len & 1)
        {
            res *= (len + 1) / 2;
            res *= len;
        }
        else
        {
            res *= len / 2;
            res *= len + 1;
        }
        return res;
    }
};
```



### Solution Two - In Top Solutions

[C++, concise solution, O(n)](https://discuss.leetcode.com/topic/107978/c-concise-solution-o-n)

```c++
class Solution {
public:
    int numSubarrayProductLessThanK(vector<int>& nums, int k) {
        if (k <= 1) return 0;
        int n = nums.size(), prod = 1, ans = 0, left = 0;
        for (int i = 0; i < n; i++) {
            prod *= nums[i];
            while (prod >= k) prod /= nums[left++];
            ans += i - left + 1;
        }
        return ans;
    }
};
```



