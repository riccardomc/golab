package datastructures

import "testing"

func TestMiddle(t *testing.T) {

	t.Run("0 elements", func(t *testing.T) {
		l := &Node{-1, nil}

		actual := Middle(l)

		if actual != nil {
			t.Errorf("should be null")
		}
	})

	t.Run("1 element", func(t *testing.T) {
		l := &Node{-1, nil}

		l.In(1)

		expected := 1
		actual := Middle(l).Key

		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("2 elements", func(t *testing.T) {
		l := &Node{-1, nil}

		l.In(1)
		l.In(2)

		expected := 1
		actual := Middle(l).Key

		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("3 elements", func(t *testing.T) {
		l := &Node{-1, nil}
		l.In(1)
		l.In(2)
		l.In(3)

		expected := 2
		actual := Middle(l).Key

		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("4 elements", func(t *testing.T) {
		l := &Node{-1, nil}
		l.In(1)
		l.In(2)
		l.In(3)
		l.In(4)

		expected := 2
		actual := Middle(l).Key

		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

	t.Run("100 elements", func(t *testing.T) {
		l := &Node{-1, nil}
		for i := 1; i <= 100; i++ {
			l.In(i)
		}

		expected := 50
		actual := Middle(l).Key

		if actual != expected {
			t.Errorf("%d != %d", actual, expected)
		}
	})

}
