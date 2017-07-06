package palindrome

func palindromePermutation(input string) bool {
	m := map[rune]bool{}

	for _, r := range input {
		if _, ok := m[r]; ok {
			delete(m, r)
		} else {
			m[r] = true
		}
	}

	return len(m) <= 1
}
