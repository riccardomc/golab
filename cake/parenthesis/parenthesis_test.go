package parenthesis

import "testing"

func TestParenthesis(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		input := "Sometimes (when I nest them (my parentheticals) too much (like this (and this))) they get confusing."
		expected := 79
		actual := parenthesis(input, 10)

		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})
}
