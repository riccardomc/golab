package inflight

import (
	. "github.com/riccardomc/golab/cake/validbst"
)

// movieLengths 10 30 20 15 35

// flightLength 30

//TwoMovies 14
func TwoMovies(movieLengths []int, flightLength int) bool {

	var tree *BST

	if len(movieLengths) < 2 {
		return false
	}

	tree = NewBST(movieLengths[0])
	for i := 1; i < len(movieLengths); i++ {
		Insert(tree, movieLengths[i])
	}

	for i := 0; i < len(movieLengths); i++ {
		if movieLengths[i] > flightLength {
			continue
		}
		neededLength := flightLength - movieLengths[i]
		if binarySearch(tree, neededLength) != nil {
			return true
		}
	}

	return false
}

//TwoMoviesHash 14 with HashMap
func TwoMoviesHash(movieLengths []int, fligthLength int) bool {
	lengthsMap := make(map[int]bool, 0)

	for _, v := range movieLengths {
		lengthsMap[v] = true
	}

	for i := 0; i < len(movieLengths); i++ {
		if movieLengths[i] > fligthLength {
			continue
		}
		neededLength := fligthLength - movieLengths[i]
		_, ok := lengthsMap[neededLength]
		if ok {
			return true
		}
	}
	return false
}

func binarySearch(tree *BST, key int) *BST {
	if tree == nil {
		return nil
	}

	c := tree
	for c != nil {

		if key == c.Key {
			return c
		}

		if key > c.Key {
			c = c.Right
			continue
		}

		if key < c.Key {
			c = c.Left
		}
	}

	return c
}
