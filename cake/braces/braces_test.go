package braces

import "testing"

func TestValidBraces(t *testing.T) {

	input := "{[]()}"
	expected := true
	t.Run("input", func(t *testing.T) {
		actual := validBraces(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	input = "{[(])}"
	expected = false
	t.Run("input", func(t *testing.T) {
		actual := validBraces(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	input = "{[}"
	expected = false
	t.Run("input", func(t *testing.T) {
		actual := validBraces(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})
}
