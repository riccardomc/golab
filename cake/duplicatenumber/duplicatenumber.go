package duplicatenumber

func duplicatenumber(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	sumOf1toN := (len(numbers) * (len(numbers) - 1)) / 2
	return sum - sumOf1toN
}
