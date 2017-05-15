package arraysstrings

import (
	"testing"
)

type Test struct {
	input    string
	size     int
	expected string
}

func TestUrlify(t *testing.T) {

	var samples []Test

	samples = append(samples, Test{
		input:    "a test  ",
		size:     6,
		expected: "a$20test",
	})

	samples = append(samples, Test{
		input:    "this is a test      ",
		size:     14,
		expected: "this$20is$20a$20test",
	})

	t.Run("Urlify", func(t *testing.T) {
		for _, sample := range samples {
			urlified := Urlify(sample.input, sample.size)
			if urlified != sample.expected {
				t.Errorf("'%s' != '%s'", urlified, sample.expected)
			}
		}
	})
}
