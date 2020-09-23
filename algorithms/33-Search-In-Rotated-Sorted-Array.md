## [33. Search In Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/description/)

### Description

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., `0 1 2 4 5 6 7` might become `4 5 6 7 0 1 2`).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.



**Difficult:** `Medium`

**Tags:** `Array`



### Solution One

第一层 while 循环用来确定 `nums` 向右旋转了多少次，k 的范围为 [1, length]

第二层就是普通的二分查找，既然 `nums` 向右旋转了 k 次，那么对应的 mid 也要旋转 k 次

时间复杂度：O(N) 

```c++
class Solution {
public:
    int search(vector<int>& nums, int target) {
        int length = nums.size();

        int k = 1;
        while (k < nums.size() && nums[k] > nums[k - 1])
            k++;

        int left = 0;
        int right = length - 1;
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            int m = (mid + k) % length;
            if (nums[m] == target)
            {
                return m;
            }
            else if (nums[m] < target)
            {
                left = mid + 1;
            }
            else
            {
                right = mid - 1;
            }
        }
        return -1;
    }
};
```



### Solution Two - In Top Solutions

[Concise O(log N) Binary search solution](https://discuss.leetcode.com/topic/3538/concise-o-log-n-binary-search-solution)

这里第一层循环是用二分查找法来寻找序列中最小值的索引

也可以理解为确定 `nums` 向右循环了多少次，好过 **One** 中的线性查找法，k 的范围为 [0, length -1]

时间复杂度：O(log N)

**Update:** 第一层循环目的与题目 [153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/) 一样

```java
class Solution {
public:
    int search(int A[], int n, int target) {
        int lo=0,hi=n-1;
        // find the index of the smallest value using binary search.
        // Loop will terminate since mid < hi, and lo or hi will shrink by at least 1.
        // Proof by contradiction that mid < hi: if mid==hi, then lo==hi and loop would have been terminated.
        while(lo<hi){
            int mid=(lo+hi)/2;
            if(A[mid]>A[hi]) lo=mid+1;
            else hi=mid;
        }
        // lo==hi is the index of the smallest value and also the number of places rotated.
        int rot=lo;
        lo=0;hi=n-1;
        // The usual binary search and accounting for rotation.
        while(lo<=hi){
            int mid=(lo+hi)/2;
            int realmid=(mid+rot)%n;
            if(A[realmid]==target)return realmid;
            if(A[realmid]<target)lo=mid+1;
            else hi=mid-1;
        }
        return -1;
    }
};
```

