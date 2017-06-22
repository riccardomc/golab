package datastructures

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestBinarySearchTree(t *testing.T) {

	t.Run("Insert/Find", func(t *testing.T) {
		tree := NewBTNode()

		rand.Seed(time.Now().UTC().UnixNano())
		data := make([]int, 10)
		for i := 0; i < 10; i++ {
			data[i] = rand.Int() % 100
			tree.Insert(data[i])
		}
		actual := " "
		BTNodeBFSApply(tree.Left, func(n *BTNode) {
			actual = actual + " " + strconv.Itoa(n.Key)
		})
		t.Log(actual)

	})
}
