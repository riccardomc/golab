package cakethief

import "testing"

func TestMaxDuffleBagValue(t *testing.T) {
	input := []CakeType{
		CakeType{2, 15},
		CakeType{3, 90},
		CakeType{7, 160},
	}

	t.Run("Capacity 20", func(t *testing.T) {
		expected := 555
		actual := MaxDuffleBagValue(input, 20)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})
}
