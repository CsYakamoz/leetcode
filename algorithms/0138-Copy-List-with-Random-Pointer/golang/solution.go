package golang

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// reference Hint 3
	p := head
	for p != nil {
		cp := &Node{Val: p.Val, Next: p.Next, Random: nil}
		p.Next = cp
		p = cp.Next
	}

	result := head.Next

	// reference Hint 4
	p = head
	for p != nil {
		cp := p.Next
		if p.Random != nil {
			cp.Random = p.Random.Next
		}

		p = cp.Next
	}

	// restore the original linked list
	p = head
	for p != nil {
		cp := p.Next
		p.Next = cp.Next
		p = p.Next
		if p != nil {
			cp.Next = p.Next
		}
	}

	return result
}
