package arraysstrings

import (
	"testing"
)

func TestUnique(t *testing.T) {

	samples := make(map[string]bool)
	samples["acdc"] = false
	samples["abba"] = false
	samples["abcd"] = true
	samples["love"] = true

	t.Run("Unique", func(t *testing.T) {
		for k, v := range samples {
			if samples[k] != Unique(k) {
				t.Error(k, "should be", v)
			}
		}
	})

	t.Run("Unique2", func(t *testing.T) {
		for k, v := range samples {
			if samples[k] != Unique2(k) {
				t.Error(k, "should be", v)
			}
		}
	})
}
