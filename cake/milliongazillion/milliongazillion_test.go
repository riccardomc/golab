package milliongazillion

import (
	"sort"
	"testing"
)

func TestFindRotationPoint(t *testing.T) {
	input := []string{
		"ptolemaic",
		"retrograde",
		"supplant",
		"undulate",
		"xenoepist",
		"asymptote", // <-- rotates here!
		"babka",
		"banoffee",
		"engender",
		"karpatka",
		"othellolagkage",
	}

	t.Run("words", func(t *testing.T) {
		expected := 5
		actual := FindRotationPoint(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	sort.Strings(input)

	t.Run("words without rotation", func(t *testing.T) {
		expected := len(input) - 1
		actual := FindRotationPoint(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	input = []string{
		"c",
		"d",
		"e",
		"a",
		"b",
	}

	t.Run("letters", func(t *testing.T) {
		expected := 3
		actual := FindRotationPoint(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})
}
