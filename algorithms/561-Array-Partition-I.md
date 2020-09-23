## [561. Array Partition I](https://leetcode.com/problems/array-partition-i/#/description)

### Description

Given an array of **2n** integers, your task is to group these integers into **n** pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

**Example 1:**

```
Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
```

**Note:**

1. **n** is a positive integer, which is in the range of [1, 10000].
2. All the integers in the array will be in the range of [-10000, 10000].



**Difficult:** `Easy`

**Tags:** `Array`



### Solution One

首先排序，接着，对于`nums[0]`而言，不管它跟谁成 pairs，`min(pairs)`的值都为`num[0]`，故我们需要牺牲第二小的数，即`nums[1]`，所以`nums[0]. nums[1]`成 pairs

同理，对于`nums[2]`而言，由于0，1已经使用了，故`nums[2]`是最小值，故跟谁成pairs，`min(pairs)`的值都为`nums[2]`，所以我们让`nums[2]`，`nums[3]`成pairs。

综上所述，排序后，我只需要去偶数位置的值



也可以反过来思考，假设我们逆序排序：

`nums[0]`为最大值，不管它和谁（假设为`m`）成pairs，`min(pairs)`的值都是`m`，故我们要让`m`最大，`nums[0]`已经用了，故`nums[1]`是最大，故让`nums[0], nums[1]`成pairs。

其它同理。

```c++
class Solution {
public:
    int arrayPairSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int result = 0;
        for (size_t i = 0; i < nums.size(); i += 2)
        {
            result += nums[i];
        }
        return result;
    }
};
```



### Solution Two

思路同思路一，但使用了Hash Table

```c++
class Solution {
public:
    int arrayPairSum(vector<int>& nums) {
        vector<int> hashtable(20001,0);
        for(size_t i=0;i<nums.size();i++)
        {
            hashtable[nums[i]+10000]++;
        }
        int ret=0;
        int flag=0;
        for(size_t i=0;i<20001;){
            if((hashtable[i]>0)&&(flag==0)){
                ret=ret+i-10000;
                flag=1;
                hashtable[i]--;
            }else if((hashtable[i]>0)&&(flag==1)){
                hashtable[i]--;
                flag=0;
            }else i++;
        }
        return ret;
    }
};
```



