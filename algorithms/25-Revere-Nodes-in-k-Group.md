## [25. Revere Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/#/description)

### Description

Given a linked list, reverse the nodes of a linked list _k_ at a time and return its modified list.

_k_ is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of _k_ then left-out nodes in the end should remain as it is.

You may not alter the values in the nodes, only nodes itself may be changed.

Only constant memory is allowed.

For example,
Given this linked list: `1->2->3->4->5`

For _k_ = 2, you should return: `2->1->4->3->5`

For _k_ = 3, you should return: `3->2->1->4->5`

**Difficult:** `Hard`

**Tags:** `Linked List`

### Solution One

首先将 k 个结点添加到 vector 中

再将 vector 中结点对换

```c++
class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        // If the List is [] or the k is equal to 1, nothing to do
        if (head == nullptr || k == 1)
        {
            return head;
        }

        ListNode *newHead = new ListNode(0);
        newHead->next = head;
        ListNode *current = newHead;
        vector<ListNode *> p;

        while (current->next)
        {
            ListNode *tmpCurrent = current;
            int n = 0;
            while (current->next && n != k)
            {
                p.push_back(current->next);
                current = current->next;
                ++n;
            }
            current = tmpCurrent;
            if (n == k)
            {
                int i = 0, j = k - 1;
                while (i < j && i + 1 != j)
                {
                    ListNode *tmp = p[j]->next;
                    current->next = p[j];
                    p[j]->next = p[i + 1];
                    p[j - 1]->next = p[i];
                    p[i]->next = tmp;
                    current = current->next;
                    ++i;
                    --j;
                }
                if (i + 1 == j)
                {
                    ListNode *tmp = current->next;
                    current->next = p[j];
                    p[i]->next = p[j]->next;
                    p[j]->next = tmp;
                }
                current = p[0];
                p.clear();
            }
            else
            {
                break;
            }
        }
        return newHead->next;
    }
};
```
