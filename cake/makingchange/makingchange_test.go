package makingchange

import (
	"fmt"
	"testing"
)

func TestMakeChange(t *testing.T) {
	t.Run("Yo", func(t *testing.T) {
		MakeChange([]int{1, 2, 3}, 4)
		fmt.Println(MakeChangeBottomUp([]int{1, 2, 3}, 4))
	})
}
