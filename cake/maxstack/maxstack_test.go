package maxstack

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxStack(t *testing.T) {
	t.Run("Random push and pops", func(t *testing.T) {
		operations := rand.Int() % 1000
		stack := NewMaxStack()
		for i := 0; i < operations; i++ {
			if rand.Int()%2 == 0 || stack.Empty() {
				stack.Push(rand.Int() % 100)
			} else if !stack.Empty() {
				stack.Pop()
			}
			fmt.Println(stack.data, stack.Max())
		}
	})
}
