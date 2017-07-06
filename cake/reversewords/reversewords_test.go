package reversewords

import "testing"

func TestReverseWords(t *testing.T) {
	input := "this is a test"
	t.Run(input, func(t *testing.T) {
		expected := "test a is this"
		actual := reverse(input)
		if actual != expected {
			t.Errorf("'%s' != '%s'", actual, expected)
		}
	})
}
