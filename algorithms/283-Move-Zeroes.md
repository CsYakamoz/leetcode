## [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/#/description)

### Description

Given an array `nums`, write a function to move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.

For example, given `nums = [0, 1, 0, 3, 12]`, after calling your function, `nums` should be `[1, 3, 12, 0, 0]`.

**Note**:

1. You must do this **in-place** without making a copy of the array.
2. Minimize the total number of operations.

**Difficult:** `Easy`

**Tags:** `Array` `Two Pointers`

### Solution One

`Tow Pointers`

`front`指向第一个零，接着`back`从`front + 1`开始寻找非零数

每找到一个，就将非零移到`front`指向的位置

上面循环结束后，就将数组后面的位置赋值为`0`

```c++
class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int front = 0;
        while (front < nums.size() && nums[front] != 0)
        {	// Point to the first zero
            front++;
        }
        int back = front + 1;
        while (back < nums.size())
        {
            while (back < nums.size() && nums[back] == 0)
            {	// Point to the next non-zero
                back++;
            }
            if (back < nums.size())
            {	// Moves it
                nums[front++] = nums[back++];
            }
        }
        while (front < nums.size())
        {
            nums[front++] = 0;
        }
    }
};
```

### Solution Two - In Top Solutions

总体思路与**One**相似，但这里使用`swap`函数来交换值，少了**One**中“将数组后面的位置赋值为`0`”

```c++
class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int index = 0;
        for(int i=0; i<nums.size(); i++)
            if(nums[i]!=0)
                swap(nums[index++], nums[i]);
    }
};
```
