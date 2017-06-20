package datastructures

import "strconv"

//Node is a classic linked list node thingie
type Node struct {
	Key  int
	Next *Node
}

// In brutally inserts at the beginning of the list
func (l *Node) In(v int) {
	l.Next = &Node{v, l.Next}
}

// ToString representation of the list
func (l *Node) ToString() string {
	var toString func(*Node, string) string

	toString = func(n *Node, s string) string {
		if n == nil {
			return s + " ]"
		}
		return toString(n.Next, s+" "+strconv.Itoa(n.Key))
	}

	return toString(l.Next, "[")
}

// Reverse the directions of the pointers of a list
func (l *Node) Reverse() *Node {
	var reverse func(*Node, *Node) *Node

	reverse = func(curr *Node, prev *Node) *Node {

		if curr == nil {
			return prev
		}

		next := curr.Next
		curr.Next = prev
		return reverse(next, curr)
	}

	l.Next = reverse(l.Next, nil)
	return l
}
