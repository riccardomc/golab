package stringreverse

func stringReverse(input string) string {

	bytes := []byte(input)

	for i := 0; i < len(bytes)/2; i++ {
		symmetric := len(bytes) - 1 - i
		temp := bytes[i]
		bytes[i] = bytes[symmetric]
		bytes[symmetric] = temp
	}

	return string(bytes)
}
