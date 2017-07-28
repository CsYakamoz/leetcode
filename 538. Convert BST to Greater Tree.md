## [538. Convert BST to Greater Tree](https://leetcode.com/problems/convert-bst-to-greater-tree/#/description)

###  Description

Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

**Example:**

```
Input: The root of a Binary Search Tree like this:
              5
            /   \
           2     13

Output: The root of a Greater Tree like this:
             18
            /   \
          20     13
```



**Difficult:** `Easy`

**Tags:** `Tree`



### Solution One

`DFS`  `Recursion`

此思路太奇怪了，我也解释不了，也莫名奇妙就通过了 

`rVal`意思为：以`root->right`为根结点的子树的所有结点的和，若`root`没有右结点，则值为空。

`lVal`意思同上。

`res`意思为：以`root`为根节点的树的所有结点的和。

`sum`意思为：我也解释不了

Top Solutions 上有个跟这思路差不多：[Java 6 lines](https://discuss.leetcode.com/topic/83477/java-6-lines)

```c++
class Solution {
public:
    TreeNode* convertBST(TreeNode* root) {
        helper(root, 0);
        return root;
    }

private:
    int helper(TreeNode *root, int sum) {
        if (root) {
            int rVal = helper(root->right, sum);
            int lVal = helper(root->left, root->val + rVal + sum);
            int res = root->val + rVal + lVal;
            root->val += rVal + sum;
            return res;
        } else {
            return 0;
        }
    }
};
```



### Solution Two - In Top Solutions

[c++ solution beats 100%](https://discuss.leetcode.com/topic/83534/c-solution-beats-100)

```c++
class Solution {
private:
    int cur_sum = 0;
public:
    void travel(TreeNode* root){
        if (!root) return;
        if (root->right) travel(root->right);
        
        root->val = (cur_sum += root->val);
        if (root->left) travel(root->left);
    }
    TreeNode* convertBST(TreeNode* root) {
        travel(root);
        return root;
    }
};
```



