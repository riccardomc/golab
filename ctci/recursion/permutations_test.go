package recursion

import "testing"

type test struct {
	input  []int
	output [][]int
}

var tests []test

func compareSliceOfSlices(expected, actual [][]int) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := 0; i < len(expected); i++ {
		if len(expected[i]) != len(actual[i]) {
			return false
		}

		for j := 0; j < len(expected[i]); j++ {
			if expected[i][j] != actual[i][j] {
				return false
			}
		}
	}

	return true
}

func TestPermutations(t *testing.T) {

	tests = []test{
		test{
			input: []int{1},
			output: [][]int{
				{1},
			},
		},
		test{
			input: []int{1, 2},
			output: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		test{
			input: []int{1, 2, 3},
			output: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	myTest := tests[0]
	t.Run("1 Element", func(t *testing.T) {
		actual := Permutations(myTest.input)
		if !compareSliceOfSlices(myTest.output, actual) {
			t.Error(myTest.output, "!=", actual)
		}
	})

	myTest = tests[1]
	t.Run("2 Elements", func(t *testing.T) {
		actual := Permutations(myTest.input)
		if !compareSliceOfSlices(myTest.output, actual) {
			t.Error(myTest.output, "!=", actual)
		}
	})

	myTest = tests[2]
	t.Run("3 Elements", func(t *testing.T) {
		actual := Permutations(myTest.input)
		if !compareSliceOfSlices(myTest.output, actual) {
			t.Error(myTest.output, "!=", actual)
		}
	})
}
