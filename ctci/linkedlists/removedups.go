package linkedlists

//RemoveDups remove duplicates from the SinglyLinkedList l
func RemoveDups(l *SinglyLinkedList) *SinglyLinkedList {
	if l.Head == nil {
		return l
	}

	found := make([]bool, l.Size)
	found[l.Head.Value.(int)] = true

	for prev, c := l.Head, l.Head.Next; c != nil; {
		if found[c.Value.(int)] {
			prev.Next = c.Next
			c.Next = nil
			c = prev.Next
		} else {
			found[c.Value.(int)] = true
			prev = c
			c = c.Next
		}
	}
	return l
}
