package datastructures

import "testing"

func TestReverseWords(t *testing.T) {
	input := "I feel pretty"
	expected := "pretty feel I "
	t.Run(input, func(t *testing.T) {
		actual := ReverseWords(input)
		if actual != expected {
			t.Errorf("'%s' != '%s'", actual, expected)
		}
	})

	input = "Rising up back on the street Did my time took my chances"
	expected = "chances my took time my Did street the on back up Rising "
	t.Run(input, func(t *testing.T) {
		actual := ReverseWords(input)
		if actual != expected {
			t.Errorf("'%s' != '%s'", actual, expected)
		}
	})
}
