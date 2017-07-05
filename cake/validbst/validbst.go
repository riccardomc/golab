package validbst

import (
	"math"
	"strconv"
)

//BST isa tree
type BST struct {
	Key   int
	Left  *BST
	Right *BST
}

//BSTBounds tracks lower and upper bounds of node
type BSTBounds struct {
	node       *BST
	lowerBound int
	upperBound int
}

//NewBST gives you a new BST
func NewBST(key int) *BST {
	return &BST{key, nil, nil}
}

//Insert key in BST t
func Insert(t *BST, key int) {
	if key >= t.Key {
		if t.Right == nil {
			t.Right = NewBST(key)
		} else {
			Insert(t.Right, key)
		}
	}

	if key < t.Key {
		if t.Left == nil {
			t.Left = NewBST(key)
		} else {
			Insert(t.Left, key)
		}
	}
}

//Dump a string representation of BST t
func Dump(t *BST) string {

	repr := ""

	var dump func(*BST, int)

	dump = func(t *BST, level int) {
		if t == nil {
			return
		}

		dump(t.Right, level+1)
		for i := 0; i < level; i++ {
			repr += "    "
		}
		repr += strconv.Itoa(t.Key) + "\n"
		dump(t.Left, level+1)
	}

	dump(t, 0)

	return repr
}

//Valid this function is actually wrong, since checking the parent is not
//enough. We should check all the ancestors.
func Valid(tree *BST) bool {

	if tree.Left == nil && tree.Right == nil {
		return true
	}

	left := true
	if tree.Left != nil {
		if tree.Key <= tree.Left.Key {
			return false
		}
		left = Valid(tree.Left)
	}

	if !left {
		return false
	}

	right := true
	if tree.Right != nil {
		if tree.Key > tree.Right.Key {
			return false
		}
		right = Valid(tree.Right)
	}

	if !right {
		return false
	}

	return true
}

//ValidDFS check if tree is balanced using DFS
func ValidDFS(t *BST) bool {

	stack := make([]*BSTBounds, 0)

	push := func(b *BSTBounds) {
		stack = append(stack, b)
	}

	pop := func() *BSTBounds {
		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return b
	}

	push(&BSTBounds{t, math.MinInt64, math.MaxInt64})
	for len(stack) > 0 {
		bounds := pop()
		node := bounds.node
		if node.Key < bounds.lowerBound || node.Key > bounds.upperBound {
			return false
		}

		if node.Left != nil {
			push(&BSTBounds{node.Left, bounds.lowerBound, node.Key})
		}

		if node.Right != nil {
			push(&BSTBounds{node.Right, node.Key, bounds.upperBound})
		}
	}

	return true
}
