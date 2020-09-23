## [100. Same Tree](https://leetcode.com/problems/same-tree/#/description)

### Description

Given two binary trees, write a function to check if they are equal or not.

Two binary trees are considered equal if they are structurally identical and the nodes have the same value.



**Difficult:** `Easy`

**Tags:** `Tree` `Depth-first Search`



### Solution One

```c++
class Solution {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        if (p != nullptr && q != nullptr)
        {	// Both of them are not null
            return p->val == q->val && 
                isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
        }
        else if (p == nullptr && q == nullptr)
        {	// Both of them are null
            return true;
        }
        else
        {	// One of them is null, the other is not null
            return false;
        }
    }
};
```



### Solution Two - In Top Solutions

[My non-recursive method](https://discuss.leetcode.com/topic/7513/my-non-recursive-method)

```java
public boolean isSameTree(TreeNode p, TreeNode q) {
         Stack<TreeNode> stack_p = new Stack <> ();       
         Stack<TreeNode> stack_q = new Stack <> ();
         if (p != null) stack_p.push( p ) ;
         if (q != null) stack_q.push( q ) ;
         while (!stack_p.isEmpty() && !stack_q.isEmpty()) {
             TreeNode pn = stack_p.pop() ;
             TreeNode qn = stack_q.pop() ;	    	
             if (pn.val != qn.val) return false ;
             if (pn.right != null) stack_p.push(pn.right) ;
             if (qn.right != null) stack_q.push(qn.right) ;
             if (stack_p.size() != stack_q.size()) return false ;
             if (pn.left != null) stack_p.push(pn.left) ;	    	 	    	 
             if (qn.left != null) stack_q.push(qn.left) ;
             if (stack_p.size() != stack_q.size()) return false ;
         }		     
         return stack_p.size() == stack_q.size() ;	 
     }
```



