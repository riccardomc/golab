package linkedlists

import (
	"strconv"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	l := New()

	intToString := func(x interface{}) string {
		return strconv.Itoa(x.(int))
	}

	t.Run("RemoveDups with nil list", func(t *testing.T) {
		expected := "[]"
		actual := RemoveDups(l).ToString(intToString)
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	for i := 0; i < 10; i++ {
		l.PushValue(i)
	}

	t.Run("RemoveDups with no dups", func(t *testing.T) {
		expected := "[9 8 7 6 5 4 3 2 1 0 ]"
		actual := RemoveDups(l).ToString(intToString)
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	l.PushValue(5)

	t.Run("RemoveDups with single dup", func(t *testing.T) {
		expected := "[5 9 8 7 6 4 3 2 1 0 ]"
		actual := RemoveDups(l).ToString(intToString)
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	l.PushValue(6)
	l.PushValue(6)
	l.PushValue(6)
	l.PushValue(6)

	t.Run("RemoveDups with multiple dups at the beginning", func(t *testing.T) {
		expected := "[6 5 9 8 7 4 3 2 1 0 ]"
		actual := RemoveDups(l).ToString(intToString)
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	l.PushValue(6)
	l.PushValue(6)
	l.PushValue(6)
	l.PushValue(6)
	l.PushValue(10)
	l.PushValue(11)
	l.PushValue(12)

	t.Run("RemoveDups with multiple dups in the middle", func(t *testing.T) {
		expected := "[12 11 10 6 5 9 8 7 4 3 2 1 0 ]"
		actual := RemoveDups(l).ToString(intToString)
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})
}
