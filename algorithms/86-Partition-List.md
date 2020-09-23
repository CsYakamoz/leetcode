## [86. Partition List](https://leetcode.com/problems/partition-list/description/)

### Description

Given a linked list and a value *x*, partition it such that all nodes less than *x* come before nodes greater than or equal to *x*.

You should preserve the original relative order of the nodes in each of the two partitions.

For example,
Given `1->4->3->2->5->2` and *x* = 3,
return `1->2->2->4->3->5`.



**Difficult:** `Medium`

**Tags:** `Linked List` `Two Pointers`



### Solution One

使用二级指针

```c++
class Solution {
public:
    ListNode* partition(ListNode* head, int x) {
        ListNode **i = &head;
        while (*i && (*i)->val < x)
        {
            i = &(*i)->next;
        }
        if (*i == nullptr)
        {
            return head;
        }
        ListNode **j = &(*i)->next;
        ListNode *secondPartitionHead = *i;

        while (*j)
        {
            if ((*j)->val >= x)
            {
                j = &(*j)->next;
            }
            else
            {
                auto node = *j;
                *j = node->next;
                *i = node;
                i = &node->next;
            }
        }
        *i = secondPartitionHead;
        return head;
    }
};
```



### Solution Two - In Top Solutions

[Very concise one pass solution](https://discuss.leetcode.com/topic/7005/very-concise-one-pass-solution)

```c++
ListNode *partition(ListNode *head, int x) {
    ListNode node1(0), node2(0);
    ListNode *p1 = &node1, *p2 = &node2;
    while (head) {
        if (head->val < x)
            p1 = p1->next = head;
        else
            p2 = p2->next = head;
        head = head->next;
    }
    p2->next = NULL;
    p1->next = node2.next;
    return node1.next;
}
```



