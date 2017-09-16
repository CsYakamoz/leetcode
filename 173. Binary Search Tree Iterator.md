## [173. Binary Search Tree Iterator](https://leetcode.com/problems/binary-search-tree-iterator/description/)

### Description

Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling `next()` will return the next smallest number in the BST.

**Note: **`next()` and `hasNext()` should run in average O(1) time and uses O(*h*) memory, where *h* is the height of the tree.



**Difficult:** `Medium`

**Tags:** `Tree` `Stack` `Design`



### Solution One

思路来源于题目 [653. Two Sum IV - Input is a BST](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/) 中的 Top Solution - [[C++\] Clean Code - O(n) time O(lg n) space - BinaryTree Iterator](https://discuss.leetcode.com/topic/98391/c-clean-code-o-n-time-o-lg-n-space-binarytree-iterator)

```c++
class BSTIterator {
public:
    BSTIterator(TreeNode *root) {
        while (root)
        {
            inorder.push(root);
            root = root->left;
        }
    }

    /** @return whether we have a next smallest number */
    bool hasNext() {
        return !inorder.empty();
    }

    /** @return the next smallest number */
    int next() {
        TreeNode *node = inorder.top();
        inorder.pop();
        int val = node->val;
        node = node->right;
        while (node)
        {
            inorder.push(node);
            node = node->left;
        }
        return val;
    }

private:
    stack<TreeNode*> inorder;
};
```



