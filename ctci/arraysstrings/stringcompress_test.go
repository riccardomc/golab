package arraysstrings

import "testing"

func TestCompress(t *testing.T) {

	samples := map[string]string{
		"aabbbbcc": "a2b4c2",
		"aabbbbcd": "aabbbbcd",
		"aaabbbc":  "a3b3c1",
		"abcd":     "abcd",
		"aaaab":    "a4b1",
		"aa":       "aa",
		"a":        "a",
		"":         "",
	}

	t.Run("Compress", func(t *testing.T) {
		for sample, expected := range samples {
			actual := Compress(sample)
			if actual != expected {
				t.Errorf("%s != %s ", actual, expected)
			}
		}
	})
}
