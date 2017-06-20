package datastructures

import "testing"

func gimmieAList(n int) *Node {

	l := &Node{-1, nil}

	for i := 0; i < n; i++ {
		l.In(i)
	}

	return l
}

func TestReverse(t *testing.T) {
	t.Run("Reverse 10 elements", func(t *testing.T) {
		l := gimmieAList(10)
		expected := "[ 0 1 2 3 4 5 6 7 8 9 ]"
		actual := l.Reverse().ToString()
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	t.Run("Reverse empty list", func(t *testing.T) {
		l := gimmieAList(0)
		expected := "[ ]"
		actual := l.Reverse().ToString()
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	t.Run("Reverse 1 element list", func(t *testing.T) {
		l := gimmieAList(1)
		expected := "[ 0 ]"
		actual := l.Reverse().ToString()
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	t.Run("Reverse 2 elements list", func(t *testing.T) {
		l := gimmieAList(2)
		expected := "[ 0 1 ]"
		actual := l.Reverse().ToString()
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})
}
