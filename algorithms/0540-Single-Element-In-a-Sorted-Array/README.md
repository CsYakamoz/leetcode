## [540. Single Element In a Sorted Array](https://leetcode.com/problems/single-element-in-a-sorted-array/description/)

### Description

Given a sorted array consisting of only integers where every element appears twice except for one element which appears once. Find this single element that appears only once.

**Example 1:**

```
Input: [1,1,2,3,3,4,4,8,8]
Output: 2

```

**Example 2:**

```
Input: [3,3,7,7,10,11,11]
Output: 10

```

**Note:** Your solution should run in O(log n) time and O(1) space.

**Difficult:** `Medium`

**Tags:**

### Solution One

```c++
class Solution {
public:
    int singleNonDuplicate(vector<int>& nums) {
        int left = 0;
        int right = nums.size() - 1;
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            if (mid == 0 || mid == nums.size() - 1)
                return nums[mid];
            else if (nums[mid] != nums[mid - 1] && nums[mid] != nums[mid + 1])
                return nums[mid];
            else if (((right - left) / 2) % 2 == 1)
            {
                if (nums[mid] == nums[mid - 1]) left = mid + 1;
                else right = mid - 1;
            }
            else
            {
                if (nums[mid] == nums[mid - 1]) right = mid;
                else left = mid;
            }
        }
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    int singleNonDuplicate(vector<int>& nums) {
        int first = 0, last = nums.size() - 1;
        int middle;
        while(first < last) {
            middle = first + (last - first) / 2;
            if((middle - first) % 2) {
                if(nums[middle] != nums[middle-1]) last = middle - 1;
                else first = middle + 1;
            }
            else {
                if(nums[middle] == nums[middle-1]) last = middle - 2;
                else if(nums[middle] == nums[middle+1]) first = middle + 2;
                else return nums[middle];
            }
        }
        return nums[first];
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    int singleNonDuplicate(vector<int>& nums) {
        int i = 0, j = nums.size()-1;
        while(i < j){
            int mid = i + (j-i)/2;
            if(mid%2 ==0){//mid is even
                if(nums[mid]==nums[mid+1]) //eg., 1 1 2 2 3
                    i = mid+2;
                else //eg., 1 2 2 3 3
                    j = mid-2;
            }
            else{//mid is odd
                if(nums[mid] == nums[mid-1])
                    i = mid+1; //eg.,1 1 2
                else
                    j = mid-1;//eg., 1 2 2
            }
        }
        return nums[i];
    }
};
```
