package productofallints

import "testing"

func TestExceptAtIndex(t *testing.T) {
	t.Run("ExceptAtIndex", func(t *testing.T) {
		input := []int{1, 7, 3, 4}
		expected := []int{84, 12, 28, 21}
		actual := ExceptAtIndex(input)

		for i := 0; i < len(input); i++ {
			if actual[i] != expected[i] {
				t.Errorf("%d: %d != %d", i, actual[i], expected[i])
			}
		}
	})
}
