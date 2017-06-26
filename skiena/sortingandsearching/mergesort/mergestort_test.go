package mergesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Run("1 element", func(t *testing.T) {
		input := []int{100}
		expected := []int{100}
		MergeSort(input)
		for i := 0; i < len(input); i++ {
			if input[i] != expected[i] {
				t.Errorf("elements at %d mismatch", i)
			}
		}
	})

	t.Run("3 elements", func(t *testing.T) {
		input := []int{1000, 1, 90}
		expected := []int{1, 90, 1000}
		MergeSort(input)
		for i := 0; i < len(input); i++ {
			if input[i] != expected[i] {
				t.Errorf("elements at %d mismatch", i)
			}
		}
	})

	t.Run("1000 elements", func(t *testing.T) {
		input := []int{}
		expected := []int{}

		for i := 0; i < 1000; i++ {
			element := rand.Int() % 1000
			input = append(input, element)
			expected = append(expected, element)
		}

		sort.Ints(expected)

		MergeSort(input)
		for i := 0; i < len(input); i++ {
			if input[i] != expected[i] {
				t.Errorf("elements at %d mismatch", i)
			}
		}
	})
}
