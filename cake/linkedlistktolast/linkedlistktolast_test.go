package linkedlistktolast

import (
	"math"
	"testing"
)

func TestKtolast(t *testing.T) {
	t.Run("3toLast", func(t *testing.T) {
		input := &node{math.MaxInt64, &node{1, &node{2, &node{3, &node{4, &node{5, &node{6, nil}}}}}}}

		expected := 4
		actual := ktolast(input, 3)
		if actual == nil || actual.key != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})
}
