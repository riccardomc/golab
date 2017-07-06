package linkedlistreverse

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("Reverse", func(t *testing.T) {

		input := &node{
			math.MinInt64, &node{
				10, &node{
					11, &node{
						12, &node{
							13, &node{
								14, nil}}}}}}
		reverse(input)
		cursor := input.next
		for i := 14; i <= 10; i-- {
			if cursor.key != i {
				t.Errorf("%d != %d", cursor, i)
			}
		}
	})
}
