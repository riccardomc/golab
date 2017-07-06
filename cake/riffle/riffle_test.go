package riffle

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsRiffle(t *testing.T) {
	rand.Seed(time.Now().Unix())
	cards := make([]int, 52)
	for i := 0; i < 52; i++ {
		cards[i] = i
	}
	shuffle(cards)

	t.Run("Riffle by construction", func(t *testing.T) {
		return
		half1, half2 := riffle(cards)
		expected := true
		actual := isRiffle(cards, half1, half2)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	t.Run("Very unlikely to be Riffle", func(t *testing.T) {
		half1, half2 := unlikelyToBeRiffle(cards)
		expected := false
		actual := isRiffle(cards, half1, half2)

		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})
}
