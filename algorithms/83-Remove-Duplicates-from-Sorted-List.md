## [83. Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/)

### Description

Given a sorted linked list, delete all duplicates such that each element appear only *once*.

For example,
Given `1->1->2`, return `1->2`.
Given `1->1->2->3->3`, return `1->2->3`.



**Difficult:** `Easy`

**Tags:** `Linked List`



### Solution One

```c++
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        ListNode *res = head;
        while (head)
        {
            while (head->next && (head->val == head->next->val))
            {
                head->next = head->next->next;
            }
            head = head->next;
        }
        return res;
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if (head == NULL) {
            return NULL;
        }
        // 
        // dummy--tail -- NULL
        //        head -- xx  xx  xx
        ListNode *dummy = new ListNode(0);
        ListNode *tail = head;
        dummy->next = tail;
        
        ListNode *cur = head->next;
        tail->next = NULL;
        
        while (cur != NULL) {
            ListNode *cur_next = cur->next;
            if (cur->val != tail->val) {
                tail->next = cur;
                tail = tail->next;
                tail->next = NULL;
            }
            cur =cur_next;
        }
        return dummy->next;
    }
};
```



