package duplicatenumber

import "testing"

func TestDuplicateNumber(t *testing.T) {
	t.Run("yi", func(t *testing.T) {
		input := []int{4, 2, 5, 3, 3, 1}
		expected := 3
		actual := duplicatenumber(input)

		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})
}
