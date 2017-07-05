package apple

import "testing"

func TestMaxProfit(t *testing.T) {
	input := []int{10, 7, 5, 8, 11, 8}
	expected := 6
	t.Run("MaxProfitNaive", func(t *testing.T) {
		actual := MaxProfitNaive(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("MaxProfitGreedy", func(t *testing.T) {
		actual := MaxProfitGreedy(input)
		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})
}
