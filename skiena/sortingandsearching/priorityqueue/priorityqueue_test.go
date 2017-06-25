package priorityqueue

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestPriorityQueue(t *testing.T) {
	t.Run("Insert one element", func(t *testing.T) {
		q := NewPriorityQueue(10)

		q.Insert(10)
		if q.Size() != 1 {
			t.Errorf("Size %d != %d", q.Size(), 1)
		}

		expected := 10
		actual := q.ExtractMin()

		if actual != expected {
			t.Errorf("Key %d != %d", actual, expected)
		}

		if q.Size() != 0 {
			t.Errorf("Size %d != %d", q.Size(), 0)
		}
	})

	t.Run("Insert 10 random elements", func(t *testing.T) {
		q := NewPriorityQueue(10)
		data := make([]int, 10)
		rand.Seed(time.Now().Unix())

		for i := 0; i < 10; i++ {
			data[i] = rand.Int() % 50
			q.Insert(data[i])
		}

		sort.Ints(data)

		for i := 0; i < 10; i++ {
			min := q.ExtractMin()
			if min != data[i] {
				t.Errorf("%s != %s", min, data[i])
			}
		}
	})
}
