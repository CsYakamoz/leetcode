## [448. Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/)

### Description

Given an array of integers where 1 ≤ a[i] ≤ _n_ (_n_ = size of array), some elements appear twice and others appear once.

Find all the elements of [1, *n*] inclusive that do not appear in this array.

Could you do it without extra space and in O(_n_) runtime? You may assume the returned list does not count as extra space.

**Example:**

```
Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
```

**Difficulty:** `Easy`

**Tags:** `Array`

### Solution One

```c++
class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        vector<int> res;
        for (size_t i = 0; i < nums.size(); i++)
        {
            int index = abs(nums[i]) - 1;
            if (nums[index] > 0)
            {
                nums[index] = -nums[index];
            }
        }
        for (size_t i = 0; i < nums.size(); i++)
        {
            if (nums[i] > 0)
            {
                res.push_back(i + 1);
            }
        }
        return res;
    }
};
```

### Solution Two - In Top Solutions

[5-line Java Easy-understanding](https://discuss.leetcode.com/topic/66063/5-line-java-easy-understanding)

```java
public List<Integer> findDisappearedNumbers(int[] nums) {
        List<Integer> res = new ArrayList<>();
        int n = nums.length;
        for (int i = 0; i < nums.length; i ++) nums[(nums[i]-1) % n] += n;
        for (int i = 0; i < nums.length; i ++) if (nums[i] <= n) res.add(i+1);
        return res;
    }
```
