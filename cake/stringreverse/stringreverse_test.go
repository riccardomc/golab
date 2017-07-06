package stringreverse

import "testing"

func TestStringReverse(t *testing.T) {
	t.Run("reverse", func(t *testing.T) {
		input := "this is a test"
		expected := "tset a si siht"
		actual := stringReverse(input)
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})
}
