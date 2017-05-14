package arraysstrings

import "testing"

type PermutationsSample struct {
	a            string
	b            string
	permutations bool
}

func TestPermutations(t *testing.T) {

	var samples []PermutationsSample
	samples = append(samples, PermutationsSample{
		a:            "abba",
		b:            "baab",
		permutations: true,
	})
	samples = append(samples, PermutationsSample{
		a:            "abcd",
		b:            "bdca",
		permutations: true,
	})
	samples = append(samples, PermutationsSample{
		a:            "abc",
		b:            "bdca",
		permutations: false,
	})
	samples = append(samples, PermutationsSample{
		a:            "abcc",
		b:            "bdca",
		permutations: false,
	})
	samples = append(samples, PermutationsSample{
		b:            "abcc",
		a:            "bdca",
		permutations: false,
	})

	t.Run("PermutationsBrute", func(t *testing.T) {
		for _, sample := range samples {
			if PermutationsBrute(sample.a, sample.b) != sample.permutations {
				t.Error(
					sample.a,
					" and ",
					sample.b,
					" should be ",
					sample.permutations)
			}
		}
	})

	t.Run("PermutationsSorted", func(t *testing.T) {
		for _, sample := range samples {
			if PermutationsBrute(sample.a, sample.b) != sample.permutations {
				t.Error(
					sample.a,
					" and ",
					sample.b,
					" should be ",
					sample.permutations)
			}
		}
	})
}
