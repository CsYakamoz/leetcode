## [117. Populating Next Right Pointers in Each Node II](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/description/)

### Description

Follow up for problem "_Populating Next Right Pointers in Each Node_".

What if the given tree could be any binary tree? Would your previous solution still work?

**Note:**

- You may only use constant extra space.

For example,
Given the following binary tree,

```
         1
       /  \
      2    3
     / \    \
    4   5    7

```

After calling your function, the tree should look like:

```
         1 -> NULL
       /  \
      2 -> 3 -> NULL
     / \    \
    4-> 5 -> 7 -> NULL
```

**Difficult:** `Medium`

**Tags:** `Tree` `Depth-first Search`

### Solution One

```c++
class Solution {
public:
    void connect(TreeLinkNode *root) {
        if (root == nullptr)
            return;

        if (root->right)
        {
            connect(root->right, root->next);
            connect(root->right);
        }
        if (root->left)
        {
            if (root->right)
                root->left->next = root->right;
            else
                connect(root->left, root->next);

            connect(root->left);
        }
    }

private:
    void connect(TreeLinkNode *&root, TreeLinkNode *node)
    {
        while (node)
        {
            if (node->left || node->right)
            {
                root->next = node->left ? node->left : node->right;
                break;
            }
            node = node->next;
        }
    }
};
```

### Solution Two - In Top Solutions

```c++
class Solution {
public:
    void connect(TreeLinkNode *root) {

        if(!root) return;

        TreeLinkNode* p = root->next;

        while(p){
            if(p->left){
                p=p->left;
                break;
            }

            if(p->right){
                p=p->right;
                break;
            }
            p=p->next;
        }

        if(root->left) root->left->next = root->right?root->right:p;
        if(root->right) root->right->next = p;

        connect(root->right);
        connect(root->left);


    }
};
```

### Solution Three

个人最初想到的解法，但不符合要求：**You may only use constant extra space.**

这里依旧贴出来

```c++
class Solution {
public:
    void connect(TreeLinkNode *root) {
        vector<TreeLinkNode*> node;
        connect(root, 0, node);
    }

private:
    void connect(TreeLinkNode *root, int level, vector<TreeLinkNode*> &node)
    {
        if (root == nullptr)
            return;

        if (level == node.size())
            node.push_back(root);
        else
        {
            node[level]->next = root;
            node[level] = root;
        }

        connect(root->left, level + 1, node);
        connect(root->right, level + 1, node);
    }
};
```

### Solution Four - In Top Solutions

[Simple solution using constant space](https://discuss.leetcode.com/topic/8447/simple-solution-using-constant-space)

> The idea is simple: level-order traversal.
> You can see the following code:

```java
public class Solution {
    public void connect(TreeLinkNode root) {

        while(root != null){
            TreeLinkNode tempChild = new TreeLinkNode(0);
            TreeLinkNode currentChild = tempChild;
            while(root!=null){
                if(root.left != null) { currentChild.next = root.left; currentChild = currentChild.next;}
                if(root.right != null) { currentChild.next = root.right; currentChild = currentChild.next;}
                root = root.next;
            }
            root = tempChild.next;
        }
    }
}
```
