package linkedlists

import (
	"strconv"
	"testing"
)

func TestCompare(t *testing.T) {
	l1 := New()
	l2 := New()

	t.Run("Compare with nil", func(t *testing.T) {
		expected := false
		actual := l1.Compare(nil)
		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	t.Run("Compare empty lists", func(t *testing.T) {
		expected := true
		actual := l1.Compare(l2)
		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	for i := 0; i < 10; i++ {
		l1.PushValue(i)
		l2.PushValue(i)
	}

	t.Run("Compare lists with 10 nodes", func(t *testing.T) {
		expected := true
		actual := l1.Compare(l2)
		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	l1.Pop()

	t.Run("Compare lists with different sizes", func(t *testing.T) {
		expected := false
		actual := l1.Compare(l2)
		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})

	l1.PushValue(100)

	t.Run("Compare lists with different elements", func(t *testing.T) {
		expected := false
		actual := l1.Compare(l2)
		if actual != expected {
			t.Errorf("%b != %b", actual, expected)
		}
	})
}

func TestToString(t *testing.T) {
	l := New()

	intToString := func(x interface{}) string {
		return strconv.Itoa(x.(int))
	}

	t.Run("ToString with empty list", func(t *testing.T) {
		expected := "[]"
		actual := l.ToString(intToString)

		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	l.PushValue(0)

	t.Run("ToString with one element", func(t *testing.T) {
		expected := "[0 ]"
		actual := l.ToString(intToString)

		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	for i := 1; i < 10; i++ {
		l.PushValue(i)
	}

	t.Run("ToString with 10 elements", func(t *testing.T) {
		expected := "[9 8 7 6 5 4 3 2 1 0 ]"
		actual := l.ToString(intToString)

		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})
}
