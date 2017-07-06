package riffle

import (
	"math/rand"
)

func isRiffle(shuffled, half1, half2 []int) bool {

	for i := len(shuffled) - 1; i >= 0; i-- {
		if len(half1) > 0 {
			last := len(half1) - 1
			if shuffled[i] == half1[last] {
				half1 = half1[:last]
				continue
			}
		}
		if len(half2) > 0 {
			last := len(half2) - 1
			if shuffled[i] == half2[last] {
				half2 = half2[:last]
				continue
			}
		}
		return false
	}

	return true
}

func riffle(cards []int) ([]int, []int) {
	half1 := make([]int, 0)
	half2 := make([]int, 0)
	for i := 0; i < len(cards); i++ {
		if rand.Int()%2 == 0 {
			half1 = append(half1, cards[i])
		} else {
			half2 = append(half2, cards[i])
		}
	}

	return half1, half2
}

func unlikelyToBeRiffle(cards []int) ([]int, []int) {
	half1 := make([]int, 0)
	half2 := make([]int, 0)
	for i := 0; i < len(cards); i++ {
		if rand.Int()%2 == 0 {
			half1 = append(half1, cards[i])
		} else {
			half2 = append(half2, cards[i])
		}
	}
	shuffle(half1)
	shuffle(half2)
	return half1, half2
}

func shuffle(cards []int) []int {
	for i := 0; i < len(cards); i++ {
		r := randomIndex(i, len(cards))
		temp := cards[i]
		cards[i] = cards[r]
		cards[r] = temp
	}
	return cards
}

func randomIndex(floor, ceil int) int {
	if ceil == floor {
		return ceil
	}
	return rand.Intn(ceil-floor) + floor
}
