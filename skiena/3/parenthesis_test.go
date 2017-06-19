package parenthesis

import "testing"

func TestParenthesis(t *testing.T) {

	input := "()"
	expected := true
	t.Run(input, func(t *testing.T) {
		actual := Parenthesis(input)

		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	input = "(()"
	expected = false
	t.Run(input, func(t *testing.T) {
		actual := Parenthesis(input)

		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	input = "(())"
	expected = true
	t.Run(input, func(t *testing.T) {
		actual := Parenthesis(input)

		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	input = "((())())()"
	expected = true
	t.Run(input, func(t *testing.T) {
		actual := Parenthesis(input)

		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	input = "(((())())())"
	expected = true
	t.Run(input, func(t *testing.T) {
		actual := Parenthesis(input)

		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	input = "(((())())()))"
	expected = false
	t.Run(input, func(t *testing.T) {
		actual := Parenthesis(input)

		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})
}
