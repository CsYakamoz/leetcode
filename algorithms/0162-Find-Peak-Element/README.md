## [162. Find Peak Element](https://leetcode.com/problems/find-peak-element/description/)

### Description

A peak element is an element that is greater than its neighbors.

Given an input array where `num[i] ≠ num[i+1]`, find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that `num[-1] = num[n] = -∞`.

For example, in array `[1, 2, 3, 1]`, 3 is a peak element and your function should return the index number 2.

**Note:**

Your solution should be in logarithmic complexity.

**Difficulty:** `Medium`

**Tags:** `Binary Search` `Array`

### Solution One

```c++
class Solution {
public:
    int findPeakElement(vector<int>& nums) {
        int length = nums.size();

        if (length == 1 || nums[0] > nums[1])
            return 0;
        else if (nums[length -1] > nums[length - 2])
            return length - 1;

        int index;
        int left = 0;
        int right = length - 1;
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            if (nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1])
            {
                index = mid;
                break;
            }
            else if (nums[mid] < nums[mid + 1])
            {
                left = mid + 1;
            }
            else
            {
                right = mid - 1;
            }
        }
        return index;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int findPeakElement(vector<int>& nums) {
        int left = 0;
        int right = nums.size() - 1;
        while (left + 1 < right){
            int mid = left + (right - left) / 2;
            if (nums.at(mid) > nums.at(mid - 1) && nums.at(mid) > nums.at(mid + 1)){
                return mid;
            }
            if (nums.at(mid -1) > nums.at(mid)){
                right = mid;
            }
            else{
                left = mid;
            }
        }
        if (nums.at(left) > nums.at(right)){
            return left;
        }
        else{
            return right;
        }
    }
};
```

### Other Solutions

[162. Find Peak Element](https://leetcode.com/problems/find-peak-element/solution/)

```java
public class Solution {
    public int findPeakElement(int[] nums) {
        int l = 0, r = nums.length - 1;
        while (l < r) {
            int mid = (l + r) / 2;
            if (nums[mid] > nums[mid + 1])
                r = mid;
            else
                l = mid + 1;
        }
        return l;
    }
}
```
