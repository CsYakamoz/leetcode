## [80. Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/)

### Description

Follow up for "Remove Duplicates":
What if duplicates are allowed at most *twice*?

For example,
Given sorted array *nums* = `[1,1,1,2,2,3]`,

Your function should return length = `5`, with the first five elements of *nums* being `1`, `1`, `2`, `2`and `3`. It doesn't matter what you leave beyond the new length.



**Difficult:** `Medium`

**Tags:** `Array` `Two Pointers`



### Solution One

```c++
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int length = 0;
        size_t i = 0;
        while (i < nums.size())
        {
            nums[length++] = nums[i++];
            int count = 1;
            while (i < nums.size() && nums[i] == nums[i - 1])
            {
                i++;
                count++;
            }

            if (count >= 2)
                nums[length++] = nums[i - 1];
        }
        return length;
    }
};
```



