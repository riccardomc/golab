package linkedlists

import (
	"strconv"
	"testing"
)

func TestSum(t *testing.T) {

	l1 := New()
	l2 := New()

	l1.PushValue(6)
	l1.PushValue(1)
	l1.PushValue(7)

	l2.PushValue(2)
	l2.PushValue(9)
	l2.PushValue(5)

	t.Run("Sum 1", func(t *testing.T) {
		expected := "[2 1 9 ]"
		actual := Sum(l1, l2).ToString(func(x interface{}) string { return strconv.Itoa(x.(int)) })
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	l1 = New()
	l2 = New()

	l1.PushValue(9)
	l1.PushValue(9)
	l1.PushValue(9)

	l2.PushValue(2)
	l2.PushValue(2)

	t.Run("Sum 2", func(t *testing.T) {
		expected := "[1 2 0 1 ]"
		actual := Sum(l1, l2).ToString(func(x interface{}) string { return strconv.Itoa(x.(int)) })
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})

	l1 = New()
	l2 = New()

	l1.PushValue(9)
	l2.PushValue(9)

	t.Run("Sum 3", func(t *testing.T) {
		expected := "[8 1 ]"
		actual := Sum(l1, l2).ToString(func(x interface{}) string { return strconv.Itoa(x.(int)) })
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	})
}
