## [508. Most Frequent Subtree Sum](https://leetcode.com/problems/most-frequent-subtree-sum/description/)

### Description

Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.

**Examples 1**
Input:

```
  5
 /  \
2   -3

```

return [2, -3, 4], since all the values happen only once, return all of them in any order.

**Examples 2**
Input:

```
  5
 /  \
2   -5

```

return [2], since 2 happens twice, however -5 only occur once.

**Note:** You may assume the sum of values in any subtree is in the range of 32-bit signed integer.



**Difficult:** `Medium`

**Tags:** `Hash Table` `Tree`



### Solution One

```c++
class Solution {
public:
    vector<int> findFrequentTreeSum(TreeNode* root) {
        getSum(root);

        int maxOccur = 0;
        vector<int> res;

        for (auto i : hash)
        {
            if (maxOccur < i.second)
            {
                maxOccur = i.second;
                res.clear();
                res.push_back(i.first);
            }
            else if (maxOccur == i.second)
            {
                res.push_back(i.first);
            }
        }

        return res;
    }


private:
    unordered_map<int, int> hash;

    int getSum(TreeNode *root)
    {
        if (root == nullptr)
            return 0;

        int lVal = getSum(root->left);
        int rVal = getSum(root->right);
        int sum = lVal + rVal + root->val;
        hash[sum]++;
        return sum;
    }
};
```



### Solution Two - In Top Solutions

[Short Clean C++ O(n) Solution](https://discuss.leetcode.com/topic/77763/short-clean-c-o-n-solution)

```c++
class Solution {
public:
    vector<int> findFrequentTreeSum(TreeNode* root) {
        unordered_map<int,int> counts;
        int maxCount = 0;
        countSubtreeSums(root, counts, maxCount);
        
        
        vector<int> maxSums;
        for(const auto& x :  counts){
            if(x.second == maxCount) maxSums.push_back(x.first);
        }
        return maxSums;
    }
    
    int countSubtreeSums(TreeNode *r, unordered_map<int,int> &counts, int& maxCount){
        if(r == nullptr) return 0;
        
        int sum = r->val;
        sum += countSubtreeSums(r->left, counts, maxCount);
        sum += countSubtreeSums(r->right, counts, maxCount);
        ++counts[sum];
        maxCount = max(maxCount, counts[sum]);
        return sum;
    }
};
```



