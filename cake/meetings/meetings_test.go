package meetings

import "testing"

func TestMerge(t *testing.T) {
	t.Run("10 meetings", func(t *testing.T) {
		input := []Meeting{
			Meeting{5, 8},
			Meeting{3, 4},
			Meeting{3, 5},
			Meeting{3, 8},
			Meeting{9, 12},
		}

		expected := []Meeting{
			Meeting{3, 8},
			Meeting{9, 12},
		}

		actual := MergeMeetingsBrute(input)

		if len(actual) != len(expected) {
			t.Error(actual, expected)
			return
		}

		for i, v := range actual {
			if v != expected[i] {
				t.Error(v, expected[i])
			}
		}

	})
}
