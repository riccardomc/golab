package topscores

func topScores(scores []int, highestScore int) []int {
	scoresCounts := make([]int, highestScore)

	for _, score := range scores {
		scoresCounts[score]++
	}

	position := len(scores) - 1
	sortedScores := make([]int, len(scores))
	for score, count := range scoresCounts {
		for count > 0 {
			sortedScores[position] = score
			position--
			count--
		}
	}
	return sortedScores
}
