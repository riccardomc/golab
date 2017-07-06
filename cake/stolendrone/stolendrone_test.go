package stolendrone

import "testing"

func TestFindUnique(t *testing.T) {
	input := []int{10, 100, 9, 11, 12, 10, 11, 2, 9, 12, 100}
	expected := 2
	t.Run("Hash Map", func(t *testing.T) {
		actual := FindUnique(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})
	t.Run("XOR", func(t *testing.T) {
		actual := FindUniqueXOR(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})
}
