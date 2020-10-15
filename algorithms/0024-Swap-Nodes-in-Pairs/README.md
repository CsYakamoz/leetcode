## [24. Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/#/description)

### Description

Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given `1->2->3->4`, you should return the list as `2->1->4->3`.

Your algorithm should use only constant space. You may **not** modify the values in the list, only nodes itself can be changed.

**Difficult:** `Medium`

**Tags:** `Linked List`

### Solution One

定义三个指针`firstP`，`secondP`，`thirdP`

循环每次将调换后面两个指针对应结点的顺序，然后三个指针都向前移两次

```c++
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        ListNode *first = new ListNode(0);
        ListNode *firstP = first;
        ListNode *secondP = firstP->next = head;
        if (!secondP)
        {
            return secondP;
        }
        ListNode *thirdP = secondP->next;
        while (thirdP)
        {
            ListNode *tmp = firstP->next;
            firstP->next = secondP->next;
            secondP->next = thirdP->next;
            thirdP->next = tmp;
            secondP = firstP->next;
            thirdP = secondP->next;
            int n = 0;
            while (thirdP && n != 2)
            {
                thirdP = thirdP->next;
                secondP = secondP->next;
                firstP = firstP->next;
                ++n;
            }
        }
        return first->next;
    }
};
```

### Solution Two

目前还看不懂 Top Solution
