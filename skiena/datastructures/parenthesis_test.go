package datastructures

import "testing"

func TestParenthesis(t *testing.T) {

	input := "()"
	expected := true
	expectedI := -1
	t.Run(input, func(t *testing.T) {
		actual, actualI := Parenthesis(input)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}

		if actualI != expectedI {
			t.Errorf("%d != %d", actualI, expectedI)
		}
	})

	input = "(()"
	expected = false
	expectedI = 0
	t.Run(input, func(t *testing.T) {
		actual, actualI := Parenthesis(input)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}

		if actualI != expectedI {
			t.Errorf("%d != %d", actualI, expectedI)
		}
	})

	input = "(()))()()()()"
	expected = false
	expectedI = 4
	t.Run(input, func(t *testing.T) {
		actual, actualI := Parenthesis(input)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}

		if actualI != expectedI {
			t.Errorf("%d != %d", actualI, expectedI)
		}
	})

	input = "(())"
	expected = true
	t.Run(input, func(t *testing.T) {
		actual, _ := Parenthesis(input)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	input = "((())())()"
	expected = true
	t.Run(input, func(t *testing.T) {
		actual, _ := Parenthesis(input)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	input = "(((())())())"
	expected = true
	t.Run(input, func(t *testing.T) {
		actual, _ := Parenthesis(input)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	input = "(((())())()))"
	expected = false
	expectedI = 12
	t.Run(input, func(t *testing.T) {
		actual, actualI := Parenthesis(input)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}

		if actualI != expectedI {
			t.Errorf("%d != %d", actualI, expectedI)
		}
	})
}
