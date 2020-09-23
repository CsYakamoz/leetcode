## [27. Remove Element](https://leetcode.com/problems/remove-element/#/description)

### Description

Given an array and a value, remove all instances of that value in place and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

**Example:**
Given input array *nums* = `[3,2,2,3]`, *val* = `3`

Your function should return length = 2, with the first two elements of *nums* being 2.



**Difficult:** `Easy`

**Tags:** `Array` `Two Pointers`



### Solution One

首先排序，接着在 vector 中查找第一个值为`val`的索引，然后计算值为`val`的数量，再从 vector 中删除这些值为`val`的元素

```c++
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        sort(nums.begin(), nums.end());
        auto start = find(nums.cbegin(), nums.cend(), val);
        auto end = start;
        while (end != nums.cend() && *end == val)
        {
            ++end;
        }
        nums.erase(start, end);
        return nums.size();
    }
};
```



### Solution Two

`Two Pointers`

首先排序，接着将使 `front` 指向第一个值为 `val` 的元素，`back`指向最后一个元素

每当找到一个值为`val`的元素，就交换`nums[front]`和`nums[back]`的值

19 行的 if 如果`back`已经指向值为`val`的元素，则不需要交换

```c++
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        sort(nums.begin(), nums.end());
        int front = 0;
        int back = nums.size() - 1;
        // The number of val
        int numsOfVal = 0;
        // Find the first val
        while (front < nums.size() && nums[front] != val)
        {
            ++front;
        }
        while (front <= back)
        {
            if (nums[front] == val)
            {
                ++numsOfVal;
                if (nums[back] != val)
                {
                    int tmp = nums[back];
                    nums[back] = val;
                    nums[front] = tmp;
                    --back;
                }
                  ++front;
            }
            else
            {
                break;
            }
        }
        return nums.size() - numsOfVal;
    }
};
```



### Solution Three - In Top Solutions

`Two Pointer`

每当找到一个值不为`val`的元素，就将它赋值到`begin`所指向的元素

```c++
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
         int begin=0;
         for(int i=0;i<nums.size();i++) if(nums[i]!=val) nums[begin++]=nums[i];
         return begin;
    }
};
```



