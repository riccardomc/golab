package secondlargest

import "github.com/riccardomc/golab/cake/validbst"

func secondLargest(tree *validbst.BST) *validbst.BST {
	var parent *validbst.BST

	c := tree

	for c.Right != nil {
		parent = c
		c = c.Right
	}

	if c.Left == nil {
		return parent
	}

	for c.Right != nil {
		c = c.Right
	}

	return c
}
