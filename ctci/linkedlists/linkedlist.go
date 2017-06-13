package linkedlists

// Node is the base element of my linked list
type Node struct {
	Value interface{}
	Next  *Node
}

// SinglyLinkedList is a simple implementation of a singly linked list
type SinglyLinkedList struct {
	Head *Node
	Size int
}

// New returns a new singly linked list
func New() *SinglyLinkedList {
	return new(SinglyLinkedList).Init()
}

// Init initializes a SinglyLinkedList
func (l *SinglyLinkedList) Init() *SinglyLinkedList {
	l.Head = nil
	l.Size = 0
	return l
}

// PushNode inserts the Node n at the beginning of the list
func (l *SinglyLinkedList) PushNode(n *Node) {
	n.Next = l.Head
	l.Head = n
	l.Size++
}

// PushValue creates a Node with Value v and inserts it at the beginning of the
// list
func (l *SinglyLinkedList) PushValue(v interface{}) {
	l.PushNode(&Node{Value: v, Next: nil})
}

// Pop returns the element at the beginning of the list
func (l *SinglyLinkedList) Pop() *Node {
	if l.Head == nil {
		return nil
	}

	node := l.Head
	l.Head = l.Head.Next
	node.Next = nil
	l.Size--
	return node
}

// Compare returns true if the list is the same as otherL
func (l *SinglyLinkedList) Compare(l1 *SinglyLinkedList) bool {

	if l1 == nil {
		return false
	}

	if l.Size != l1.Size {
		return false
	}

	for c, c1 := l.Head, l1.Head; c != nil; c, c1 = c.Next, c1.Next {
		if c.Value != c1.Value {
			return false
		}
	}

	return true
}

// ToString returns a string representation of the list
func (l *SinglyLinkedList) ToString(toString func(interface{}) string) string {
	s := "["
	for c := l.Head; c != nil; c = c.Next {
		s += toString(c.Value)
		s += " "
	}
	s += "]"
	return s
}
