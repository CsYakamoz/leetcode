## [75. Sort Colors](https://leetcode.com/problems/sort-colors/description/)

### Description

Given an array with _n_ objects colored red, white or blue, sort them so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

**Note:**
You are not suppose to use the library's sort function for this problem.

**Follow up:**
A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.

Could you come up with an one-pass algorithm using only constant space?

**Difficulty:** `Medium`

**Tags:** `Array` `Two Pointers` `Sort`

### Solution One

```c++
class Solution
{
public:
    void sortColors(vector<int>& nums)
    {
        int left = -1;
        int mid = 0;
        int right = nums.size();

        while (mid < right)
        {
            switch (nums[mid])
            {
            case 0:
                swap(nums[++left], nums[mid++]);
                break;
            case 1:
                mid++;
                break;
            case 2:
                swap(nums[mid], nums[--right]);
                break;
            }
        }
    }
};
```
