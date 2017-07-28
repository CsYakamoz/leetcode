## [160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/#/description)

### Description

Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:

```
A:          a1 → a2
                   ↘
                     c1 → c2 → c3
                   ↗            
B:     b1 → b2 → b3

```

begin to intersect at node c1.

**Notes:**

- If the two linked lists have no intersection at all, return `null`.
- The linked lists must retain their original structure after the function returns.
- You may assume there are no cycles anywhere in the entire linked structure.
- Your code should preferably run in O(n) time and use only O(1) memory.



**Difficult:** `Easy`

**Tags:** `Linked List`



### Solution One

`Hash Table`

```c++
class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if (headA == nullptr || headB == nullptr)
        {
            return nullptr;
        }
        set<ListNode*> s;
        while (headA)
        {
            s.insert(headA);
            headA = headA->next;
        }
        while (headB)
        {
            if (s.find(headB) != s.end())
            {
                break;
            }
            headB = headB->next;
        }
        return headB;
    }
};
```



### Solution Two - In Top Solutions

[Concise JAVA solution, O(1) memory O(n) time](https://discuss.leetcode.com/topic/5492/concise-java-solution-o-1-memory-o-n-time)

[Java solution without knowing the difference in len!](https://discuss.leetcode.com/topic/28067/java-solution-without-knowing-the-difference-in-len)

