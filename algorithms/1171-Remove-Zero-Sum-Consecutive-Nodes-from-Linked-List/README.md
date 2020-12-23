## [1171. Remove Zero Sum Consecutive Nodes from Linked List](https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/)

### Description

<p>Given the <code>head</code> of a linked list, we repeatedly delete consecutive sequences of nodes that sum to <code>0</code> until there are no such sequences.</p>

<p>After doing so, return the head of the final linked list.&nbsp; You may return any such answer.</p>

<p>&nbsp;</p>
<p>(Note that in the examples below, all sequences are serializations of <code>ListNode</code> objects.)</p>

<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> head = [1,2,-3,3,1]
<strong>Output:</strong> [3,1]
<strong>Note:</strong> The answer [1,2,1] would also be accepted.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> head = [1,2,3,-3,4]
<strong>Output:</strong> [1,2,4]
</pre>

<p><strong>Example 3:</strong></p>

<pre>
<strong>Input:</strong> head = [1,2,3,-3,-2]
<strong>Output:</strong> [1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The given linked list will contain between <code>1</code> and <code>1000</code> nodes.</li>
	<li>Each node in the linked list has <code>-1000 &lt;= node.val &lt;= 1000</code>.</li>
</ul>

**Difficulty:** `Medium`

**Tags:** `Linked List`

### Solution One

```go
func removeZeroSumSublists(head *ListNode) *ListNode {
	dict := make(map[int]int)
	dict[0] = 0

	prefixSum := make([]*ListNode, 0)
	prefixSum = append(
		prefixSum,
		&ListNode{Val: 0, Next: head},
	)

	for i, sum := 1, 0; head != nil; {
		sum += head.Val

		idx, exist := dict[sum]
		if exist {
			prefixSum[idx].Next = head.Next

			for j, key := idx+1, sum; j < len(prefixSum) && j < i; j++ {
				key += prefixSum[j].Val
				delete(dict, key)
			}

			i = idx
			prefixSum = prefixSum[0 : idx+1]
		} else {
			dict[sum] = i
			prefixSum = append(
				prefixSum,
				head,
			)
		}

		head = head.Next
		i++
	}

	return prefixSum[0].Next
}
```

### Solution Two - In Top Solutions

reference: [[Java/C++/Python] Greedily Skip with HashMap](https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/discuss/366319/JavaC%2B%2BPython-Greedily-Skip-with-HashMap)

```java
    public ListNode removeZeroSumSublists(ListNode head) {
        int prefix = 0;
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        Map<Integer, ListNode> seen = new HashMap<>();
        seen.put(0, dummy);
        for (ListNode i = dummy; i != null; i = i.next) {
            prefix += i.val;
            seen.put(prefix, i);
        }
        prefix = 0;
        for (ListNode i = dummy; i != null; i = i.next) {
            prefix += i.val;
            i.next = seen.get(prefix).next;
        }
        return dummy.next;
    }
```
