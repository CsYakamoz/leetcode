## [119. Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/description/)

### Description

Given an index _k_, return the *k*th row of the Pascal's triangle.

For example, given _k_ = 3,
Return `[1,3,3,1]`.

**Note:**
Could you optimize your algorithm to use only _O_(_k_) extra space?

**Difficult:** `Easy`

**Tags:** `Array`

### Solution One - In Top Solutions

一开始想用组合公式直接算出第 k 行，但由于组合公式是阶乘公式，对于 Int 会溢出，对于 double 精度会不足。

所以改用下面的方法。

```c++
class Solution {
public:
    vector<int> getRow(int rowIndex) {
        vector<int> row(rowIndex + 1, 1);

       for (int i = 0; i <= rowIndex; i++)
       {
           for (int j = i - 1; j > 0; j--)
           {
               row[j] = row[j] + row[j - 1];
           }
       }
        return row;
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    vector<int> getRow(int rowIndex) {

        int reflect = (rowIndex + 1) - (rowIndex + 1) / 2;

        int index = (rowIndex + 1) / 2 - 1;

        long temp = 1;

        vector<int> output;

        for(int i = 0; i <= rowIndex; i++)
        {
            if(i < reflect)
            {
                if(i != 0)
                {
                    temp = temp * (rowIndex - i + 1) / i;
                }
                output.push_back(temp);
            }else{
                output.push_back(output[index]);
                index--;
            }
        }

        return output;

    }
};
```
