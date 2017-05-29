package arraysstrings

import "testing"

type sample struct {
	s1    string
	s2    string
	value bool
}

func TestZeroOrOneDistance(t *testing.T) {

	samples := []sample{
		// insert
		sample{s1: "some", s2: "somet", value: true},
		sample{s1: "some", s2: "sotme", value: true},
		sample{s1: "some", s2: "tsome", value: true},
		sample{s1: "some", s2: "tstome", value: false},
		sample{s1: "some", s2: "somett", value: false},
		// remove
		sample{s1: "some", s2: "som", value: true},
		sample{s1: "some", s2: "ome", value: true},
		sample{s1: "some", s2: "soe", value: true},
		sample{s1: "some", s2: "so", value: false},
		// replace or 0 distance
		sample{s1: "some", s2: "some", value: true},
		sample{s1: "some", s2: "smme", value: true},
		sample{s1: "some", s2: "somm", value: true},
		sample{s1: "some", s2: "kome", value: true},
		sample{s1: "some", s2: "komm", value: false},
	}

	t.Run("insert", func(t *testing.T) {
		for _, s := range samples {
			if ZeroOrOneDistance(s.s1, s.s2) != s.value {
				t.Errorf("%s %s should be %t", s.s1, s.s2, s.value)
			}
		}
	})
}
