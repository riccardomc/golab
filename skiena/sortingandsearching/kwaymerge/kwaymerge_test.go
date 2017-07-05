package kwaymerge

import (
	"math/rand"
	"sort"
	"testing"
)

func generateInput(lists int) [][]int {
	list := make([][]int, lists)
	for i := 0; i < lists; i++ {
		size := rand.Int() % 50
		list[i] = make([]int, size)
		for j := 0; j < size; j++ {
			list[i][j] = rand.Int() % 1000
		}
		sort.Ints(list[i])
	}
	return list
}

func TestKWayMerge(t *testing.T) {

	t.Run("10 lists", func(t *testing.T) {

		lists := generateInput(10)
		sorted := KWayMerge(lists)

		prev := sorted[0]

		for i := 1; i < len(sorted); i++ {
			if prev > sorted[i] {
				t.Errorf("element %d at %d is out of order", i, sorted[i])
				return
			}
		}
	})
}
