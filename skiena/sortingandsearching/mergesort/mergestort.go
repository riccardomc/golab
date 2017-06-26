package mergesort

//MergeSort is a classic MergeSort implementation that uses a queue to merge
func MergeSort(s []int) {
	mergesort(s, 0, len(s)-1)
}

func mergesort(s []int, low, high int) {
	if low < high {
		middle := (low + high) / 2
		mergesort(s, low, middle)
		mergesort(s, middle+1, high)
		merge(s, low, middle, high)
	}
}

func merge(s []int, low, middle, high int) {
	bufferLow := make([]int, 0)
	bufferHigh := make([]int, 0)

	for i := low; i <= middle; i++ {
		bufferLow = append(bufferLow, s[i])
	}

	for i := middle + 1; i <= high; i++ {
		bufferHigh = append(bufferHigh, s[i])
	}

	i := low
	for len(bufferLow) > 0 && len(bufferHigh) > 0 {
		if bufferLow[0] <= bufferHigh[0] {
			s[i] = bufferLow[0]
			bufferLow = bufferLow[1:]
		} else {
			s[i] = bufferHigh[0]
			bufferHigh = bufferHigh[1:]
		}
		i++
	}

	for len(bufferLow) > 0 {
		s[i] = bufferLow[0]
		bufferLow = bufferLow[1:]
		i++
	}

	for len(bufferHigh) > 0 {
		s[i] = bufferHigh[0]
		bufferHigh = bufferHigh[1:]
		i++
	}
}
