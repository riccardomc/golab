package recursion

import (
	"fmt"
	"testing"
)

func TestTripleStepRecursiveRecursive(t *testing.T) {

	samples := map[int]int{
		1: 1,
		2: 2,
		3: 4,
		4: 7,
		5: 13,
		6: 24,
		7: 44,
	}

	for k, expected := range samples {
		t.Run(fmt.Sprintf("TripleStepRecursive %d", k), func(t *testing.T) {
			actual := TripleStepRecursive(k)
			if actual != expected {
				t.Errorf("%d: %d != %d", k, actual, expected)
			}
		})
	}
}

func BenchmarkTripleStepRecursiveRecursive(b *testing.B) {

	samples := []int{10, 20, 30, 35}

	for _, sample := range samples {
		b.Run(fmt.Sprintf("TripleStepRecursive(%d)", sample), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TripleStepRecursive(sample)
			}
		})
	}
}

func TestTripleStepMemo(t *testing.T) {

	samples := map[int]int{
		1: 1,
		2: 2,
		3: 4,
		4: 7,
		5: 13,
		6: 24,
		7: 44,
	}

	for k, expected := range samples {
		t.Run(fmt.Sprintf("TripleStepMemo %d", k), func(t *testing.T) {
			actual := TripleStepMemo(k)
			if actual != expected {
				t.Errorf("%d: %d != %d", k, actual, expected)
			}
		})
	}
}

func BenchmarkTripleStepMemo(b *testing.B) {

	samples := []int{10, 20, 30, 35}

	for _, sample := range samples {
		b.Run(fmt.Sprintf("TripleStepMemo(%d)", sample), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TripleStepMemo(sample)
			}
		})
	}
}

func TestTripleStepDynamic(t *testing.T) {

	samples := map[int]int{
		1: 1,
		2: 2,
		3: 4,
		4: 7,
		5: 13,
		6: 24,
		7: 44,
	}

	for k, expected := range samples {
		t.Run(fmt.Sprintf("TripleStepDynamic %d", k), func(t *testing.T) {
			actual := TripleStepDynamic(k)
			if actual != expected {
				t.Errorf("%d: %d != %d", k, actual, expected)
			}
		})
	}
}

func BenchmarkTripleStepDynamic(b *testing.B) {

	samples := []int{10, 20, 30, 35}

	for _, sample := range samples {
		b.Run(fmt.Sprintf("TripleStepDynamic(%d)", sample), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TripleStepDynamic(sample)
			}
		})
	}
}
