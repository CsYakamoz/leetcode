## [26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/#/description)

### Description

Given a sorted array, remove the duplicates in place such that each element appear only *once* and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array *nums* = `[1,1,2]`,

Your function should return length = `2`, with the first two elements of *nums* being `1` and `2` respectively. It doesn't matter what you leave beyond the new length.

**Difficulty:** `Easy`

**Tags:** `Array` `Two Pointers`

### Solution One

计数的同时，使用`erase`将重复的删除

```c++
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int length = 0;
        auto p = nums.begin();
        while (p != nums.end())
        {
            ++length;
            auto val = *p++;
            while (p != nums.end() && *p == val)
            {
                p = nums.erase(p);
            }
        }
        return length;
    }
};
```

### Solution Two

使用 C++ 标准库函数 [std::unique](http://www.cplusplus.com/reference/algorithm/unique/?kw=unique) ，该函数会将相邻的重复项“消除”，并返回一个指向不重复值范围末尾的迭代器

**update:** 第 5 行可以直接删除，并返回`f - nums.begin()`

```c++
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        auto f = unique(nums.begin(), nums.end());
        nums.erase(f, nums.end());
        return nums.size();
    }
};
```

### Solution Three

`Two Pointers`

`length`表示前指针，同时也表示长度，`i`为后指针

思路为：每当 i 指向一个与 i - 1 不同的值，就将 nums[i] 复制到 nums[length]

```c++
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int length = 0;
        int i = 0;
        while (i < nums.size())
        {
            nums[length] = nums[i];
            ++length;
            ++i;
            while (i < nums.size() && nums[i] == nums[i-1])
            {
                // skip duplicates number
                ++i;
            }
        }
        return length;
    }
};
```
