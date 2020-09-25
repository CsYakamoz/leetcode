## [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/#/description)

Merge _k_ sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

**Difficult:** `Hard`

**Tags:** `Divide and Conquer` `Linked List` `Heap`

### Solution One

`Sort`

首先将所有结点添加到 vector 中， 然后排序（关键字： val）

然后将前后链接起来

```c++
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        vector<ListNode*> v;
        for (auto i : lists)
        {
            while (i)
            {
                v.push_back(i);
                i = i->next;
            }
        }
        if (v.empty())
        {
            return nullptr;
        }
        sort(v.begin(), v.end(), [](ListNode *x, ListNode *y) { return x->val < y->val; });
        for (size_t i = 1; i < v.size(); i++)
        {
            v[i - 1]->next = v[i];
        }
        v.back()->next = nullptr;
        return v[0];
    }
};
```

### Solution Two

使用 _Merge Two Sorted Lists_ 中的思路

将第一个与第二个合并，得到新的一个

将新的一个与第三个合并，得到新的一个

...

最后合并为一个 List

```c++
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode *result = lists.size() == 0 ? nullptr : lists[0];
        for (size_t i = 1; i < lists.size(); i++)
        {
            result = merge(result, lists[i]);
        }
        return result;
    }

private:
    ListNode* merge(ListNode *l1, ListNode *l2)
    {
        ListNode *head = new ListNode(0);
        ListNode *current = head;
        while (l1 || l2)
        {
            if (l1 == nullptr)
            {
                current->next = l2;
                current = current->next;
                l2 = l2->next;
            }
            else if (l2 == nullptr)
            {
                current->next = l1;
                current = current->next;
                l1 = l1->next;
            }
            else
            {
                if (l1->val < l2->val)
                {
                    current->next = l1;
                    current = current->next;
                    l1 = l1->next;
                }
                else
                {
                    current->next = l2;
                    current = current->next;
                    l2 = l2->next;
                }
            }
        }
        return head->next;
    }
};
```

### Solution Three - In Top Solutions

```c++
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode *result = lists.size() == 0 ? nullptr : lists[0];
        for (size_t i = 1; i < lists.size(); i++)
        {
            result = merge(result, lists[i]);
        }
        return result;
    }

private:
    ListNode* merge(ListNode *l1, ListNode *l2)
    {
        ListNode *head = new ListNode(0);
        ListNode *current = head;
        while (l1 && l2)
        {
            if (l1->val < l2->val)
            {
                current->next = l1;
                l1 = l1->next;
            }
            else
            {
                current->next = l2;
                l2 = l2->next;
            }
            current = current->next;
        }
        current->next = l1 ? l1 : l2;
        return head->next;
    }
};
```
