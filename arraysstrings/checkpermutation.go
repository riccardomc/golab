package arraysstrings

import (
	"sort"
	"strings"
)

// PermutationsBrute check if a is a permutation of b brute force
func PermutationsBrute(a, b string) bool {

	if len(a) != len(b) {
		return false
	}

	mapA := make(map[rune]int)
	mapB := make(map[rune]int)

	for _, runeA := range a {
		count, found := mapA[runeA]
		if found {
			mapA[runeA] = count + 1
		} else {
			mapA[runeA] = 0
		}
	}

	for _, runeB := range b {
		count, found := mapB[runeB]
		if found {
			mapB[runeB] = count + 1
		} else {
			mapB[runeB] = 0
		}
	}

	if len(mapA) != len(mapB) {
		return false
	}

	for k, v := range mapA {
		if v != mapB[k] {
			return false
		}
	}

	return true
}

// PermutationSorted sort the strings and compare
func PermutationSorted(a, b string) bool {

	if len(a) != len(b) {
		return false
	}

	sliceA := strings.Split(a, "")
	sliceB := strings.Split(b, "")

	sort.Strings(sliceA)
	sort.Strings(sliceB)

	for i := 0; i < len(sliceA); i++ {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}
