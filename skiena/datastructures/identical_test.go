package datastructures

import (
	"testing"

	. "github.com/riccardomc/golab/skiena/datastructures/binarytree"
)

func TestIdentical(t *testing.T) {

	tree1 := &BTNode{
		1,
		&BTNode{2,
			&BTNode{4, nil, nil},
			&BTNode{5, nil, nil},
		},
		&BTNode{3,
			&BTNode{6, nil, nil},
			&BTNode{7, nil, nil},
		},
	}

	tree2 := &BTNode{
		1,
		&BTNode{2,
			&BTNode{4, nil, nil},
			&BTNode{5, nil, nil},
		},
		&BTNode{3,
			&BTNode{6, nil, nil},
			&BTNode{7,
				&BTNode{100, nil, nil},
				nil,
			},
		},
	}

	t.Run("Same tree", func(t *testing.T) {
		expected := true
		actual := Identical(tree1, tree1)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	t.Run("Extra leaf", func(t *testing.T) {
		expected := false
		actual := Identical(tree1, tree2)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	t.Run("Add missing leaf", func(t *testing.T) {
		tree1.Right.Right.Left = &BTNode{100, nil, nil}
		expected := true
		actual := Identical(tree1, tree2)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	t.Run("Drop subtree", func(t *testing.T) {
		tree1.Right = nil
		expected := false
		actual := Identical(tree1, tree2)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})
}
