package arraysstrings

func PalindromePermutation(input string) bool {

	counts := make(map[rune]int)

	for _, r := range input {
		count, found := counts[r]
		if found {
			counts[r] = count + 1
		} else {
			counts[r] = 1
		}
	}

	if len(input)%2 == 0 {
		for _, v := range counts {
			if v%2 != 0 {
				return false
			}
		}
	} else {
		odd_count_found := false
		for _, v := range counts {
			if v%2 != 0 {
				if odd_count_found {
					return false
				} else {
					odd_count_found = true
				}
			}
		}
	}

	return true
}
