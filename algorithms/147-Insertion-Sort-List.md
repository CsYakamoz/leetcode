## [147. Insertion Sort List](https://leetcode.com/problems/insertion-sort-list/description/)

### Description

Sort a linked list using insertion sort.



**Difficult:** `Medium`

**Tags:** `Linked List` `Sort`



### Solution One

Using double pointer.

Example:

```
4 -> 3 -> 2 -> 1 -> nullptr

nullptr <- 4,     3 -> 2 -> 1 ->nullptr
    nullptr <- 3 <- 4,     2 -> 1 -> nullptr
        nullptr <- 2 <- 3 <- 4,     1 -> nullptr
            nullptr <- 1 <- 2 <- 3 <- 4

Reverse List:
    1 -> 2 -> 3 -> 4 -> nullptr
```

```c++
class Solution {
public:
    ListNode* insertionSortList(ListNode* head) {
        if (head == nullptr || head->next == nullptr)
            return head;

        ListNode *prev = head;
        head = head->next;
        prev->next = nullptr;

        while (head)
        {
            ListNode *node = head->next;
            head->next = prev;
            prev = head;
            head = node;
            prev = insert(prev);
        }

        head = prev;
        prev = nullptr;
        while (head)
        {
            ListNode *node = head->next;
            head->next = prev;
            prev = head;
            head = node;
        }
        return prev;
    }

private:
    ListNode* insert(ListNode *head)
    {
        ListNode **curr = &head;
        while (*curr)
        {
            ListNode *entry = *curr;
            if (entry->next && entry->val < entry->next->val)
            {
                *curr = entry->next;
                curr = &(*curr)->next;
                entry->next = *curr;
                *curr = entry;
            }
            else
            {
                break;
            }
        }
        return head;
    }
};
```



### Solution Two - In Top Solutions

```c++
class Solution {
public:
    ListNode* insertionSortList(ListNode* head) {
        ListNode *dummy = new ListNode(0), *p=dummy;
        
        vector<int> vals;
        while(head){
            vals.push_back(head->val);
            head = head->next;
        }
        sort(vals.begin(), vals.end());
        for(auto& val : vals){
            p->next = new ListNode(val);
            p = p->next;
        }
        
        
        return dummy->next;
    }
};
```



### Solution Three - In Top Solutions

创建一个新的链表，逐个将结点添加到新的链表。

```c++
class Solution {
public:
	ListNode* insertionSortList(ListNode* head) {
		if (head == nullptr)
			return head;
		ListNode * cur = nullptr;
		ListNode *newHead = new ListNode(0);
		newHead->next = head;
		ListNode *tail = head;
		head = head->next;
		while (head)
		{
			if (head->val < tail->val)
			{
				ListNode* prev = newHead;
				ListNode* cur = prev->next;
				//search for first that's greater than head
				while (cur->val <= head->val)
				{
					prev = cur;
					cur=cur->next;
				}
				tail->next = head->next;
				prev->next = head;
				head->next = cur;
				head = tail->next;
			}
			else
			{
				tail = head;
				head = head->next;
			}
		}
		return newHead->next;
	}
};
```



