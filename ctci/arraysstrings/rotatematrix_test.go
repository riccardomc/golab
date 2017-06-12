package arraysstrings

import "testing"

func TestRotate(t *testing.T) {

	s1 := [][]int{
		[]int{0, 1, 2, 3},
		[]int{4, 5, 6, 7},
		[]int{8, 9, 10, 11},
		[]int{12, 13, 14, 15},
	}

	t.Run("Rotate 2x2", func(t *testing.T) {
		t.Log(s1)
		Rotate(s1)
		t.Log(s1)
	})
}
