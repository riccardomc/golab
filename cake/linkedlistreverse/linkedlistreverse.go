package linkedlistreverse

type node struct {
	key  int
	next *node
}

func reverse(l *node) {
	if l.next == nil {
		return
	}
	prev := l.next
	curr := l.next.next

	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}

	l.next = prev
}
