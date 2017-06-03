package arraysstrings

import "testing"

func TestPalindromePermutation(t *testing.T) {

	samples := make(map[string]bool)

	samples["tactcoa"] = true
	samples["tactcxoa"] = false
	samples["ababa"] = true
	samples["ababab"] = false
	samples["ababb"] = true
	samples["aa"] = true
	samples["ab"] = false
	samples["acxbcaddb"] = true

	t.Run("PalindromePermutation", func(t *testing.T) {
		for k, v := range samples {
			result := PalindromePermutation(k)
			if result != v {
				t.Error(k, "should be", v)
			}
		}

	})
}
