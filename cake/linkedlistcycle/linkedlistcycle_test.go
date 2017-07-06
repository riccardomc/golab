package linkedlistcycle

import (
	"testing"
)

func TestCycle(t *testing.T) {

	input := newLinkedList()
	cyclicNode := &node{100, nil}
	cyclicNode.next = cyclicNode
	input.next = &node{10, &node{11, &node{12, cyclicNode}}}
	t.Run("Cycle", func(t *testing.T) {
		expected := true
		actual := cycle(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	t.Run("Cycle", func(t *testing.T) {
		expected := true
		actual := cycleLap(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	input.next = &node{10, &node{11, &node{12, nil}}}
	t.Run("No cycle", func(t *testing.T) {
		expected := false
		actual := cycle(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})

	t.Run("No cycle Lap", func(t *testing.T) {
		expected := false
		actual := cycleLap(input)
		if actual != expected {
			t.Errorf("%t != %t", actual, expected)
		}
	})
}
