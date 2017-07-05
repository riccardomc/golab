package validbst

import (
	"math/rand"
	"testing"
)

func generateTree() *BST {
	var tree *BST
	for i := 0; i < 10; i++ {
		if tree == nil {
			tree = NewBST(rand.Int() % 100)
			continue
		}
		Insert(tree, rand.Int()%100)
	}
	return tree
}

func TestValidBST(t *testing.T) {
	rand.Seed(101)
	input := generateTree()
	t.Run("Valid", func(t *testing.T) {
		expected := true
		actual := ValidDFS(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
			t.Log("\n" + Dump(input))
		}
	})

	tree := input
	for ; tree.Right != nil; tree = tree.Right {
	}
	tree.Right = NewBST(-1)

	t.Run("Not Valid", func(t *testing.T) {
		expected := false
		actual := ValidDFS(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
			t.Log("\n" + Dump(input))
		}
	})

	input =
		&BST{50,
			&BST{30,
				&BST{20, nil, nil},
				&BST{60, nil, nil},
			},
			&BST{80,
				&BST{70, nil, nil},
				&BST{90, nil, nil},
			},
		}

	t.Run("Not Valid Sneaky", func(t *testing.T) {
		expected := false
		actual := ValidDFS(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
			t.Log("\n" + Dump(input))
		}
	})

	input =
		&BST{50,
			&BST{30,
				&BST{20, nil, nil},
				&BST{35, nil, nil},
			},
			&BST{80,
				&BST{70, nil, nil},
				&BST{90, nil, nil},
			},
		}

	t.Run("Valid Sneaky", func(t *testing.T) {
		expected := true
		actual := ValidDFS(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
			t.Log("\n" + Dump(input))
		}
	})

	input =
		&BST{50,
			&BST{30,
				&BST{20, nil, nil},
				&BST{35, nil, nil},
			},
			&BST{80,
				&BST{40, nil, nil},
				&BST{90, nil, nil},
			},
		}

	t.Run("Invalid Sneaky Lowerbound", func(t *testing.T) {
		expected := false
		actual := ValidDFS(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
			t.Log("\n" + Dump(input))
		}
	})
}
