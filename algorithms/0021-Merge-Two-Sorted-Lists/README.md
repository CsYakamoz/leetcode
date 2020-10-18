## [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/#/description)

### Description

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

**Difficulty:** `Easy`

**Tags:** `Linked List`

### Solution One

当`p1`和`p2`都非空时，选择值较小的结点添加到`current`后面

当`p1`或`p2`为空时，将非空的结点添加到`current`后面

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
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode *head = new ListNode(0);
        ListNode *current = head;
        ListNode *p1 = l1;
        ListNode *p2 = l2;
        while (p1 || p2)
        {
            if (p1 == nullptr)
            {
                current->next = p2;
                current = current->next;
                p2 = p2->next;
            }
            else if (p2 == nullptr)
            {
                current->next = p1;
                current = current->next;
                p1 = p1->next;
            }
            else
            {
                if (p1->val < p2->val)
                {
                    current->next = p1;
                    current = current->next;
                    p1 = p1->next;
                }
                else
                {
                    current->next = p2;
                    current = current->next;
                    p2 = p2->next;
                }
            }
        }
        return head->next;
    }
};
```

### Solution Two - In Top Solutions

与**Solution One**类似，但这里通过判断是否同时为非空来判断是否结束循环

```c++
class Solution {
public:
    ListNode *mergeTwoLists(ListNode *l1, ListNode *l2) {
        ListNode dummy(INT_MIN);
        ListNode *tail = &dummy;

        while (l1 && l2) {
            if (l1->val < l2->val) {
                tail->next = l1;
                l1 = l1->next;
            } else {
                tail->next = l2;
                l2 = l2->next;
            }
            tail = tail->next;
        }

        tail->next = l1 ? l1 : l2;
        return dummy.next;
    }
};
```
