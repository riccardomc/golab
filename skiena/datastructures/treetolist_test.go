package datastructures

import (
	"math/rand"
	"sort"
	"testing"
)

func TestTreeToList(t *testing.T) {
	t.Run("Simple tree", func(t *testing.T) {
		expected := "[ 0 1 2 3 4 5 6 7 8 9 ]"

		tree := NewBTNode()
		for i := 0; i < 10; i++ {
			tree.Insert(i)
		}

		actual := TreeToList(tree.Left).Reverse().ToString()

		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	t.Run("Random tree", func(t *testing.T) {

		size := 100
		tree := NewBTNode()
		data := make([]int, size)
		for i := 0; i < size; i++ {
			data[i] = rand.Int() % 1000
			tree.Insert(data[i])
		}

		expectedList := sliceToList(data)
		actualList := TreeToList(tree.Left)
		actual := actualList.ToString()
		expected := expectedList.ToString()

		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})
}

//sliceToList returns a singly linked list with the elements in data without
//duplicates
func sliceToList(data []int) *Node {
	list := &Node{}

	sort.Ints(data)
	seen := make(map[int]bool)
	for _, v := range data {
		if _, duplicate := seen[v]; !duplicate {
			seen[v] = true
			list.In(v)
		}
	}

	return list
}
