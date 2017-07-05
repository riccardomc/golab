package twostacks

import (
	"testing"
)

func TestTwoStacks(t *testing.T) {
	t.Run("Naive", func(t *testing.T) {
		queue := NewIntQueue()
		input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := input
		for i := 0; i < len(input); i++ {
			queue.EnqueueNaive(input[i])
		}

		for i := 0; i < len(input); i++ {
			actual := queue.DequeueNaive()
			if actual != expected[i] {
				t.Errorf("%d != %d", actual, expected)
			}
		}
	})

	t.Run("Not Naive", func(t *testing.T) {
		queue := NewIntQueue()
		input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := input
		for i := 0; i < len(input); i++ {
			queue.Enqueue(input[i])
		}

		for i := 0; i < len(input); i++ {
			actual := queue.Dequeue()
			if actual != expected[i] {
				t.Errorf("%d != %d", actual, expected)
			}
		}
	})
}
