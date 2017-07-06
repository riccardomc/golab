package linkedlistcycle

import (
	"math"
)

type node struct {
	key  int
	next *node
}

func newLinkedList() *node {
	return &node{math.MinInt64, nil}
}

func (l *node) push(key int) {
	l.next = &node{key, l.next}
}

func cycle(l *node) bool {
	if l == nil {
		return false
	}

	nodeMap := make(map[*node]bool, 0)
	for c := l.next; c != nil; c = c.next {
		if _, ok := nodeMap[c]; ok {
			return true
		}
		nodeMap[c] = true
	}

	return false
}

func cycleLap(l *node) bool {

	fast := l.next
	slow := l.next

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			return true
		}
	}

	return false
}
