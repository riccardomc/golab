package highestproduct

import "testing"

func TestHighestProduct(t *testing.T) {
	t.Run("Postitive", func(t *testing.T) {
		input := []int{1, 7, 9, 5, 2, 3, 8}
		expected := 7 * 9 * 8
		actual := HighestProduct(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}

	})

	t.Run("Negatives", func(t *testing.T) {
		input := []int{1, 5, 9, -7, 2, 3, -8}
		expected := 9 * -7 * -8
		actual := HighestProduct(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}

	})

	t.Run("Positives and Negatives", func(t *testing.T) {
		input := []int{1, 5, -9, 7, 2, 3, 8}
		expected := 5 * 7 * 8
		actual := HighestProduct(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}

	})
}
