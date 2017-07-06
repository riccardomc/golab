package permutations

import (
	"fmt"
	"testing"
)

func TestPermutations(t *testing.T) {
	t.Run("abc", func(t *testing.T) {
		actual := permutations("abcdefghil")
		fmt.Println(len(actual))
	})
}
