package arraysstrings

import (
	"sort"
	"strings"
)

// Unique returns true if all runes in string are unique, false otherwise
func Unique(input string) bool {
	seen := make(map[rune]bool)

	for _, r := range input {
		_, found := seen[r]
		if found {
			return false
		}
		seen[r] = true
	}

	return true
}

// Unique2 same as Unique, implemented with no data structures
func Unique2(input string) bool {
	s := strings.Split(input, "")
	sort.Strings(s)
	return uniqueAux(1, s)
}

func uniqueAux(index int, input []string) bool {

	if index >= len(input) {
		return true
	}

	if input[index-1] == input[index] {
		return false
	}

	return uniqueAux(index+1, input)
}
