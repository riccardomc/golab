package reversewords

func reverseBytes(input []byte, begin, end int) {
	for begin < end {
		temp := input[begin]
		input[begin] = input[end]
		input[end] = temp
		begin++
		end--
	}
}

func reverse(message string) string {
	bytes := []byte(message)

	reverseBytes(bytes, 0, len(bytes)-1)

	previousI := 0
	for i := 0; i <= len(bytes); i++ {
		if i == len(bytes) || bytes[i] == ' ' {
			reverseBytes(bytes, previousI, i-1)
			previousI = i + 1
		}
	}

	return string(bytes)
}
