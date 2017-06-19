package linkedlists

// Kth returns the kth to last Node in the SinglyLinkedList l
func Kth(l *SinglyLinkedList, k int) *Node {
	var kth *Node

	for counter, c := 0, l.Head; c != nil; c = c.Next {
		if counter > k {
			kth = kth.Next
			continue
		}

		if counter == k {
			kth = l.Head
		}

		counter++
	}

	return kth
}
