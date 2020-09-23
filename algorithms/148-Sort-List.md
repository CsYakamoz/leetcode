## [148. Sort List](https://leetcode.com/problems/sort-list/description/)

### Description

Sort a linked list in _O_(_n_ log _n_) time using constant space complexity.

**Difficult:** `Medium`

**Tags:** `Linked List` `Sort`

### Solution One

```c++
class Solution {
public:
    ListNode* sortList(ListNode* head) {
        if (head == nullptr)
        {
            return head;
        }
        vector<ListNode*> lists;
        while (head)
        {
            lists.push_back(head);
            head = head->next;
        }
        sort(lists.begin(), lists.end(),
            [](ListNode *x, ListNode *y) {
            return  x->val < y->val;
        });
        for (size_t i = 1; i < lists.size(); i++)
        {
            lists[i - 1]->next = lists[i];
        }
        lists.back()->next = nullptr;
        return lists[0];
    }
};
```

### Solution Two - In Top Solutions

[Java merge sort solution](https://discuss.leetcode.com/topic/18100/java-merge-sort-solution)

```java
public class Solution {

  public ListNode sortList(ListNode head) {
    if (head == null || head.next == null)
      return head;

    // step 1. cut the list to two halves
    ListNode prev = null, slow = head, fast = head;

    while (fast != null && fast.next != null) {
      prev = slow;
      slow = slow.next;
      fast = fast.next.next;
    }

    prev.next = null;

    // step 2. sort each half
    ListNode l1 = sortList(head);
    ListNode l2 = sortList(slow);

    // step 3. merge l1 and l2
    return merge(l1, l2);
  }

  ListNode merge(ListNode l1, ListNode l2) {
    ListNode l = new ListNode(0), p = l;

    while (l1 != null && l2 != null) {
      if (l1.val < l2.val) {
        p.next = l1;
        l1 = l1.next;
      } else {
        p.next = l2;
        l2 = l2.next;
      }
      p = p.next;
    }

    if (l1 != null)
      p.next = l1;

    if (l2 != null)
      p.next = l2;

    return l.next;
  }

}
```
