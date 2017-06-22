package datastructures

import (
	"strconv"
	"testing"
)

func TestFind(t *testing.T) {

	tree := &BTNode{
		1,
		&BTNode{
			2,
			&BTNode{
				4,
				nil,
				nil,
			},
			&BTNode{
				5,
				&BTNode{
					100,
					&BTNode{
						101,
						nil,
						nil,
					},
					nil,
				},
				nil,
			},
		},
		&BTNode{
			3,
			&BTNode{
				6,
				nil,
				nil,
			},
			&BTNode{
				7,
				nil,
				nil,
			},
		},
	}

	t.Run("Find root", func(t *testing.T) {
		expected := 1
		actual := Find(tree, 1)

		if actual == nil {
			t.Error("Got nil!")
			return
		}

		if actual.Key != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("Find mid", func(t *testing.T) {
		expected := 3
		actual := Find(tree, 3)

		if actual == nil {
			t.Error("Got nil!")
			return
		}

		if actual.Key != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("Find right leaf", func(t *testing.T) {
		expected := 5
		actual := Find(tree, 5)

		if actual == nil {
			t.Error("Got nil!")
			return
		}

		if actual.Key != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("Find deep leaf", func(t *testing.T) {
		expected := 101
		actual := Find(tree, 101)

		if actual == nil {
			t.Error("Got nil!")
			return
		}

		if actual.Key != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("Find non existing", func(t *testing.T) {
		actual := Find(tree, 111)

		if actual != nil {
			t.Error("Should be nil!")
		}
	})
}

func TestTraversal(t *testing.T) {

	tree := &BTNode{
		1,
		&BTNode{
			2,
			&BTNode{
				4,
				nil,
				nil,
			},
			&BTNode{
				5,
				nil,
				nil,
			},
		},
		&BTNode{
			3,
			&BTNode{
				6,
				nil,
				nil,
			},
			&BTNode{
				7,
				nil,
				nil,
			},
		},
	}

	t.Run("DFS", func(t *testing.T) {
		actual := ""
		expected := " 1 2 4 5 3 6 7"
		toStr := func(n *BTNode) {
			actual = actual + " " + strconv.Itoa(n.Key)
		}
		BTNodeDFSApply(tree, toStr)

		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}

	})

	t.Run("BFS", func(t *testing.T) {
		actual := ""
		expected := " 1 2 3 4 5 6 7"
		toStr := func(n *BTNode) {
			actual = actual + " " + strconv.Itoa(n.Key)
		}
		BTNodeBFSApply(tree, toStr)

		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}

	})
}
