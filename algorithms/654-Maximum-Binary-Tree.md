## [654. Maximum Binary Tree](https://leetcode.com/problems/maximum-binary-tree/description/)

### Description

Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

1. The root is the maximum number in the array.
2. The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
3. The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.

Construct the maximum tree by the given array and output the root node of this tree.

**Example 1:**

```
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    / 
     2  0   
       \
        1

```

**Note:**

1. The size of the given array will be in the range [1,1000].



**Difficult:** `Medium`

**Tags:** `Tree`



### Solution One

```c++
class Solution {
public:
    TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
        // The size of the given array will be in the range [1,1000]
        // We don't need to check whether nums is empty
        return helper(nums, 0, nums.size() - 1);
    }

    TreeNode* helper(vector<int> &nums, int left, int right)
    {
        if (left > right)
        {
            return nullptr;
        }
        int index = findMaxIndex(nums, left, right);
        TreeNode *root = new TreeNode(nums[index]);
        root->left = helper(nums, left, index - 1);
        root->right = helper(nums, index + 1, right);
        return root;
    }

    int findMaxIndex(vector<int> &nums, int left, int right)
    {
        int maxValue = INT_MIN;
        int index = -1;
        while (left <= right)
        {
            if (maxValue < nums[left])
            {
                maxValue = nums[left];
                index = left;
            }
            left++;
        }
        return index;
    }
};
```



### Solution Two - In Top Solutions

这个思路贼强！！！

```c++
class Solution {
public:
    TreeNode* insert(TreeNode* root,int x)
    {
        if(root == NULL)
            root=new TreeNode(x);
        else
        {
            if(x < root->val)
                root->right = insert(root->right,x);
            else 
            {
                TreeNode* newRoot = new TreeNode(x);
                newRoot->left=root;
                root=newRoot;
            }
        }
        return root;
    }
    TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
        TreeNode* ans=NULL;
        int len = nums.size();
        for(int i=0;i<len;i++)
            ans = insert(ans,nums[i]);
        return ans;
    }
};
```



### Solution Three - In Top Solutions

这个思路也贼强！！！

[C++ O(N) solution](https://discuss.leetcode.com/topic/98509/c-o-n-solution)

> This solution is inspired by [@votrubac](https://discuss.leetcode.com/uid/76475)
> Here is his brilliant solution
> <https://discuss.leetcode.com/topic/98454/c-9-lines-o-n-log-n-map>
>
> The key idea is:
>
> 1. We scan numbers from left to right, build the tree one node by one step;
> 2. We use a stack to keep some (not all) tree nodes and ensure a decreasing order;
> 3. For each number, we keep pop the stack until empty or a bigger number; The bigger number (if exist, it will be still in stack) is current number's root, and the last popped number (if exist) is current number's right child (temporarily, this relationship may change in the future); Then we push current number into the stack.

```c++
class Solution {
public:
    TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
        vector<TreeNode*> stk;
        for (int i = 0; i < nums.size(); ++i)
        {
            TreeNode* cur = new TreeNode(nums[i]);
            while (!stk.empty() && stk.back()->val < nums[i])
            {
                cur->left = stk.back();
                stk.pop_back();
            }
            if (!stk.empty())
                stk.back()->right = cur;
            stk.push_back(cur);
        }
        return stk.front();
    }
};
```

> This solution will be slightly faster than normal recursive solution.
>
> Again, great thanks to [@votrubac](https://discuss.leetcode.com/uid/76475) !!!