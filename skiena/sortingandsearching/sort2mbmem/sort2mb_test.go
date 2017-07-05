package sort2mbmem

import (
	"testing"
)

func TestGenerateInput(t *testing.T) {
	t.Run("generateInput", func(t *testing.T) {
		generateInput("/tmp/bla", 1)
		readBytes("/tmp/bla")
	})
}

func TestMaxHeap(t *testing.T) {
	t.Run("test MaxHeap", func(t *testing.T) {
		input := []int32{2, 0, 1, 3, 9, 7, 5, 4, 6, 8}
		heap := NewMaxHeap(10)

		for _, v := range input {
			heap.Insert(v)
		}

		var element int32 = 9
		for heap.size > 0 {
			max := heap.Max()
			if max != element {
				t.Errorf("%d != %d", max, element)
			}
			element--
		}

	})
}
