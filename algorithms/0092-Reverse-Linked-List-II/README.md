## [92. Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/description/)

### Description

Reverse a linked list from position _m_ to _n_. Do it in-place and in one-pass.

For example:
Given `1->2->3->4->5->NULL`, _m_ = 2 and _n_ = 4,

return `1->4->3->2->5->NULL`.

**Note:**
Given _m_, _n_ satisfy the following condition:
1 ? _m_ ? _n_ ? length of list.

**Difficulty:** `Medium`

**Tags:** `Linked List`

### Solution One

```c++
class Solution {
public:
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        if (m == n) return head;

        int i = 1;
        ListNode **curr = &head;
        while (i++ < m)
            curr = &(*curr)->next;

        *curr = reverseList(*curr, n - m + 1);

        return head;
    }

    ListNode* reverseList(ListNode *begin, int length)
    {
        ListNode *pre = nullptr;
        ListNode *tail = begin;
        for (int i = 0; i < length; i++)
        {
            ListNode *node = begin->next;
            begin->next = pre;
            pre = begin;
            begin = node;
        }
        tail->next = begin;
        return pre;
    }
};
```

### Solution Two - In Top Solutions

[Simple Java solution with clear explanation](https://discuss.leetcode.com/topic/8976/simple-java-solution-with-clear-explanation)

> Simply just reverse the list along the way using 4 pointers: dummy, pre, start, then

```java
public ListNode reverseBetween(ListNode head, int m, int n) {
    if(head == null) return null;
    ListNode dummy = new ListNode(0); // create a dummy node to mark the head of this list
    dummy.next = head;
    ListNode pre = dummy; // make a pointer pre as a marker for the node before reversing
    for(int i = 0; i<m-1; i++) pre = pre.next;

    ListNode start = pre.next; // a pointer to the beginning of a sub-list that will be reversed
    ListNode then = start.next; // a pointer to a node that will be reversed

    // 1 - 2 -3 - 4 - 5 ; m=2; n =4 ---> pre = 1, start = 2, then = 3
    // dummy-> 1 -> 2 -> 3 -> 4 -> 5

    for(int i=0; i<n-m; i++)
    {
        start.next = then.next;
        then.next = pre.next;
        pre.next = then;
        then = start.next;
    }

    // first reversing : dummy->1 - 3 - 2 - 4 - 5; pre = 1, start = 2, then = 4
    // second reversing: dummy->1 - 4 - 3 - 2 - 5; pre = 1, start = 2, then = 5 (finish)

    return dummy.next;

}
```
