package okeros

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	t.Run("Test 2 Rect", func(t *testing.T) {
		r1 := Rectangle{6, 1, 7, 4}
		r2 := Rectangle{9, 3, 7, 5}
		expected := Rectangle{9, 3, 4, 2}
		actual := Intersection(r1, r2)
		fmt.Println(actual, expected)
	})
}
