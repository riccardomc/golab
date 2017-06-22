package datastructures

//TreeToList returns a list representation of the tree
func TreeToList(t *BTNode) *Node {

	l := &Node{-1, nil}

	var toList func(*BTNode)

	toList = func(t *BTNode) {
		if t == nil {
			return
		}

		toList(t.Left)
		l.In(t.Key)
		toList(t.Right)
	}

	toList(t)

	return l
}
