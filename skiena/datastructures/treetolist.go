package datastructures

import "github.com/riccardomc/golab/skiena/datastructures/binarytree"

//TreeToList returns a list representation of the tree
func TreeToList(t *binarytree.BTNode) *Node {

	l := &Node{-1, nil}

	var toList func(*binarytree.BTNode)

	toList = func(t *binarytree.BTNode) {
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
