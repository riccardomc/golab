package topscores

import (
	"testing"
)

func TestTopScores(t *testing.T) {
	t.Run("no repetitions", func(t *testing.T) {
		input := []int{37, 89, 41, 65, 91, 53}
		expected := []int{91, 89, 65, 53, 41, 37}
		actual := topScores(input, 100)
		for i := 0; i < len(input); i++ {
			if actual[i] != expected[i] {
				t.Errorf("%d: %d != %d", i, actual, expected)
			}
		}
	})

	t.Run("with repetitions", func(t *testing.T) {
		input := []int{37, 89, 41, 65, 65, 65, 65, 91, 53}
		expected := []int{91, 89, 65, 65, 65, 65, 53, 41, 37}
		actual := topScores(input, 100)
		for i := 0; i < len(input); i++ {
			if actual[i] != expected[i] {
				t.Errorf("%d: %d != %d", i, actual, expected)
			}
		}
	})
}
