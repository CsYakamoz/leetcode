## [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/description/)

### Description

Given an array of _n_ integers where _n_ > 1, `nums`, return an array `output` such that `output[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

Solve it **without division** and in O(_n_).

For example, given `[1,2,3,4]`, return `[24,12,8,6]`.

**Follow up:**
Could you solve it with constant space complexity? (Note: The output array **does not** count as extra space for the purpose of space complexity analysis.)

**Difficult:** `Medium`

**Tags:** `Array`

### Solution One - In Top Solutions

[My simple Java solution](https://discuss.leetcode.com/topic/19033/my-simple-java-solution)

> Use `tmp` to store temporary multiply result by two directions. Then fill it into `result`. Bingo!

```c++
class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        vector<int> result(nums.size(), 1);

        for (int i = 0, tmp = 1; i < nums.size(); i++)
        {
            result[i] = tmp;
            tmp *= nums[i];
        }

        for (int i = nums.size() - 1, tmp = 1; i >= 0; i--)
        {
            result[i] *= tmp;
            tmp *= nums[i];
        }

        return result;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        vector<int> result(nums.size());
        result[0] = 1;
        for(int i = 1;i<nums.size();i++){
          result[i] = result[i-1]*nums[i-1];
        }
        int right = 1;
        for(int j = nums.size()-1;j>=0;j--){
          result[j] *= right;
          right *= nums[j];
        }
        return result;
    }
};
```
