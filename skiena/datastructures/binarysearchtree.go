package datastructures

import "math"

//BinarySearch a key in the tree
func (t *BTNode) BinarySearch(k int) *BTNode {
	if k > t.Key {
		if t.Right != nil {
			return t.Right.BinarySearch(k)
		}
		return nil
	}

	if k < t.Key {
		if t.Left != nil {
			return t.Left.BinarySearch(k)
		}
		return nil
	}

	return t
}

//Insert a key in the tree
func (t *BTNode) Insert(k int) {
	if k > t.Key {
		if t.Right == nil {
			t.Right = &BTNode{k, nil, nil}
		} else {
			t.Right.Insert(k)
		}
	}

	if k < t.Key {
		if t.Left == nil {
			t.Left = &BTNode{k, nil, nil}
		} else {
			t.Left.Insert(k)
		}
	}
}

//NewBTNode get a new binary tree
func NewBTNode() *BTNode {
	return &BTNode{math.MaxInt64, nil, nil}
}
