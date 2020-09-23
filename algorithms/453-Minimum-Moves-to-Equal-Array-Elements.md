## [453. Minimum Moves to Equal Array Elements](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/description/)

### Description

Given a **non-empty** integer array of size *n*, find the minimum number of moves required to make all array elements equal, where a move is incrementing *n* - 1 elements by 1.

**Example:**

```
Input:
[1,2,3]

Output:
3

Explanation:
Only three moves are needed (remember each move increments two elements):

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
```



**Difficult:** `Easy`

**Tags:** `Math`



### Solution One - In Top Solutions

[Java O(n) solution. Short.](https://discuss.leetcode.com/topic/66557/java-o-n-solution-short)

> Adding `1` to `n - 1` elements is the same as subtracting `1` from one element, w.r.t goal of making the elements in the array equal.
> So, best way to do this is make all the elements in the array equal to the `min` element.
> `sum(array) - n * minimum`

```c++
class Solution {
public:
    int minMoves(vector<int>& nums) {
        int res = 0;
        int minVal = INT_MAX;
        for (auto i : nums)
            minVal = min(minVal, i);
        for (auto i : nums)
            res += i - minVal;
        return res;
    }
};
```



### Solution Two - In Top Solutions

[It is a math question](https://discuss.leetcode.com/topic/66737/it-is-a-math-question)

> let's define sum as the sum of all the numbers, before any moves; minNum as the min number int the list; n is the length of the list;
>
> After, say m moves, we get all the numbers as x , and we will get the following equation
>
> ```
>  sum + m * (n - 1) = x * n
>
> ```
>
> and actually,
>
> ```
>   x = minNum + m
>
> ```
>
> and finally, we will get
>
> ```
>   sum - minNum * n = m
>
> ```
>
> So, it is clear and easy now.

