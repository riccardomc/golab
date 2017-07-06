package palindrome

import "testing"

func TestPalindromePermutation(t *testing.T) {

	inputs := map[string]bool{
		"civic":         true,
		"anab":          false,
		"aaann":         true,
		"abvc":          false,
		"aaabbbcccc":    false,
		"aaaabbbbccccc": true,
	}

	for input, expected := range inputs {

		t.Run(input, func(t *testing.T) {
			actual := palindromePermutation(input)
			if actual != expected {
				t.Errorf("%t != %t", actual, expected)
			}
		})
	}
}
