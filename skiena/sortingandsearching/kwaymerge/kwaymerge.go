package kwaymerge

import (
	"math"
)

//KWayMerge the lists in list
func KWayMerge(list [][]int) []int {
	l := make([]int, 0)

	if len(list) == 0 {
		return l
	}

	for {
		emptyLists := 0
		min := math.MaxInt64
		minIndex := -1
		for i := 0; i < len(list); i++ {
			if len(list[i]) == 0 {
				emptyLists++
				continue
			}
			if list[i][0] < min {
				min = list[i][0]
				minIndex = i
			}
		}

		l = append(l, min)
		if len(list[minIndex]) > 0 {
			list[minIndex] = list[minIndex][1:]
		}

		if emptyLists == len(list)-1 {
			break
		}
	}

	return l
}
