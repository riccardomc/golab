package inflight

import "testing"

func TestTwoMovies(t *testing.T) {
	lengths := []int{12, 20, 60, 10, 20, 15, 35}
	length := 30
	t.Run("Exist", func(t *testing.T) {
		expected := true
		actual := TwoMovies(lengths, length)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	length = 300
	t.Run("Don't Exist", func(t *testing.T) {
		expected := false
		actual := TwoMovies(lengths, length)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	length = 30
	t.Run("Exist Hash", func(t *testing.T) {
		expected := true
		actual := TwoMoviesHash(lengths, length)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	length = 300
	t.Run("Don't Exist Hash", func(t *testing.T) {
		expected := false
		actual := TwoMoviesHash(lengths, length)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})
}
