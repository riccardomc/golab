package linkedlists

//Sum returns the sum of l1 and l2
func Sum(l1, l2 *SinglyLinkedList) *SinglyLinkedList {
	l := New()

	carry := 0
	c1, c2 := l1.Head, l2.Head
	for ; c1 != nil && c2 != nil; c1, c2 = c1.Next, c2.Next {
		v := c1.Value.(int) + c2.Value.(int) + carry
		appendValue(l, v%10)
		carry = v / 10
	}

	for ; c1 != nil; c1 = c1.Next {
		v := c1.Value.(int) + carry
		appendValue(l, v%10)
		carry = v / 10
	}

	for ; c2 != nil; c2 = c2.Next {
		v := c2.Value.(int) + carry
		appendValue(l, v%10)
		carry = v / 10
	}

	if carry > 0 {
		appendValue(l, carry)
	}

	return l
}

func appendValue(l *SinglyLinkedList, v interface{}) {
	var c *Node

	if l.Head == nil {
		l.PushValue(v)
		return
	}

	for c = l.Head; c.Next != nil; {
		c = c.Next
	}

	c.Next = &Node{
		Value: v,
		Next:  nil,
	}
}
