package milliongazillion

func FindRotationPoint(words []string) int {
	start := 0
	end := len(words) - 1

	for start < end {
		mid := (start + end) / 2

		if words[start] >= words[mid] {
			end = mid
		} else {
			start = mid
		}
	}

	return end + 1
}
