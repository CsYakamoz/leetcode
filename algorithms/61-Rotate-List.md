## [61. Rotate List](https://leetcode.com/problems/rotate-list/description/)

### Description

Given a list, rotate the list to the right by *k* places, where *k* is non-negative.

For example:
Given `1->2->3->4->5->NULL` and *k* = `2`,
return `4->5->1->2->3->NULL`.



**Difficult:** `Medium`

**Tags:** `Linked List` `Two Pointers`



### Solution One

思路来源于题目 [189. Rotate Array](https://leetcode.com/problems/rotate-array/solution/#approach-4-using-reverse-accepted) 

```c++
class Solution {
public:
    ListNode* rotateRight(ListNode* head, int k) {
        if (head == nullptr) return nullptr;

        // Reverse all nodes
        head = reverseList(head);
        k = k % length;

        if (k == 0) return reverseList(head);

        // Reverse first k nodes
        ListNode *h1 = reverseList(head, k);
        // Reverse remaining nodes
        h2 = reverseList(h2);
        // Connect
        tail->next = h2;

        return h1;
    }

private:
    int length;
    ListNode *h2;
    ListNode *tail;

    ListNode* reverseList(ListNode *root)
    {
        ListNode *prev = nullptr;
        while (root)
        {
            length++;
            ListNode *node = root->next;
            root->next = prev;
            prev = root;
            root = node;
        }
        return prev;
    }

    ListNode* reverseList(ListNode *root, int k)
    {
        ListNode *prev = nullptr;
        tail = root;
        while (k--)
        {
            ListNode *node = root->next;
            root->next = prev;
            prev = root;
            root = node;
        }
        h2 = root;
        return prev;
    }
};
```



### Solution Two - In Top Solutions

[My clean C++ code, quite standard (find tail and reconnect the list)](https://discuss.leetcode.com/topic/14470/my-clean-c-code-quite-standard-find-tail-and-reconnect-the-list)

> There is no trick for this problem. Some people used slow/fast pointers to find the tail node but I don't see the benefit (in the sense that it doesn't reduce the pointer move op) to do so. So I just used one loop to find the length first.

```c++
class Solution {
public:
    ListNode* rotateRight(ListNode* head, int k) {
        if(!head) return head;
        
        int len=1; // number of nodes
        ListNode *newH, *tail;
        newH=tail=head;
        
        while(tail->next)  // get the number of nodes in the list
        {
            tail = tail->next;
            len++;
        }
        tail->next = head; // circle the link

        if(k %= len) 
        {
            for(auto i=0; i<len-k; i++) tail = tail->next; // the tail node is the (len-k)-th node (1st node is head)
        }
        newH = tail->next; 
        tail->next = NULL;
        return newH;
    }
};
```

