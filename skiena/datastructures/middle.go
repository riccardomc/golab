package datastructures

//Middle returns the middle node of a list
func Middle(l *Node) *Node {
	if l == nil {
		return nil
	}
	counter := 0
	middle := l.Next

	for c := l.Next; c != nil; c = c.Next {
		counter++
		if counter%2 == 0 {
			middle = middle.Next
		}
	}
	return middle
}
