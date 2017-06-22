package datastructures

import (
	"math/rand"
	"testing"
	"time"
)

func TestBinarySearchTree(t *testing.T) {

	t.Run("Insert/Search", func(t *testing.T) {
		tree := NewBTNode()

		rand.Seed(time.Now().UTC().UnixNano())
		data := make([]int, 10)
		for i := 0; i < 10; i++ {
			data[i] = rand.Int() % 100
			tree.Insert(data[i])
		}

		for i := 0; i < 10; i++ {
			found := tree.BinarySearch(data[i])
			if found == nil {
				t.Error("Node not found but should be there")
				return
			}

			if found.Key != data[i] {
				t.Errorf("%d != %d", found.Key, data[i])
			}
		}

		for i := 100; i < 110; i++ {
			found := tree.BinarySearch(i)
			if found != nil {
				t.Error("Node found but should not be there")
			}

		}

	})
}
