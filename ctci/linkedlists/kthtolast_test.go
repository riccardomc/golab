package linkedlists

import (
	"strconv"
	"testing"
)

func TestKth(t *testing.T) {

	intToString := func(x interface{}) string {
		return strconv.Itoa(x.(int))
	}

	l := New()

	for i := 0; i < 10; i++ {
		l.PushValue(i)
	}

	t.Run("3th to last on 10 elements", func(t *testing.T) {
		expected := 3
		actual := Kth(l, 3).Value.(int)
		if actual != expected {
			t.Errorf("%d != %d %s", actual, expected, l.ToString(intToString))
		}
	})

	t.Run("0th to last on 10 elements", func(t *testing.T) {
		expected := 0
		actual := Kth(l, 0).Value.(int)
		if actual != expected {
			t.Errorf("%d != %d %s", actual, expected, l.ToString(intToString))
		}
	})

	t.Run("100th to last on 10 elements", func(t *testing.T) {
		actual := Kth(l, 100)
		if actual != nil {
			t.Errorf("%d != %d %s", actual, nil, l.ToString(intToString))
		}
	})

	l = New()

	l.PushValue(0)
	l.PushValue(1)

	t.Run("1th to last on 2 elements", func(t *testing.T) {
		expected := 1
		actual := Kth(l, 1).Value.(int)
		if actual != expected {
			t.Errorf("%d != %d %s", actual, expected, l.ToString(intToString))
		}
	})

	l = New()

	l.PushValue(0)

	t.Run("0th to last on 1 elements", func(t *testing.T) {
		expected := 1
		actual := Kth(l, 1).Value.(int)
		if actual != expected {
			t.Errorf("%d != %d %s", actual, expected, l.ToString(intToString))
		}
	})
}
