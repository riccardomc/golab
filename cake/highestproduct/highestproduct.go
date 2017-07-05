package highestproduct

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//HighestProduct returns the highest product of three integers in input
func HighestProduct(input []int) int {

	highest := maxInt(input[0], input[1])
	lowest := minInt(input[0], input[1])

	highestProductOf2 := input[0] * input[1]
	lowestProductOf2 := input[0] * input[1]

	highestProductOf3 := input[0] * input[1] * input[3]

	for i := 2; i < len(input); i++ {
		current := input[i]

		highestProductOf3 = maxInt(highestProductOf3,
			maxInt(current*highestProductOf2, current*lowestProductOf2))

		highestProductOf2 = maxInt(highestProductOf2,
			maxInt(current*highest, current*lowest))

		lowestProductOf2 = minInt(lowestProductOf2,
			minInt(current*highest, current*lowest))

		highest = maxInt(highest, current)

		lowest = minInt(lowest, current)
	}

	return highestProductOf3
}
