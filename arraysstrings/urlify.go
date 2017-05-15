package arraysstrings

import "strings"

// Urlify yeah
func Urlify(inputString string, size int) string {

	shift := 0
	input := strings.Split(inputString, "")

	for i := 0; i < size+shift; i++ {
		if input[i] == " " {
			for j := size + shift - 1; j > i; j-- {
				input[j+2] = input[j]
			}
			input[i] = "$"
			input[i+1] = "2"
			input[i+2] = "0"
			i = i + 2
			shift = shift + 2
		}
	}

	return strings.Join(input, "")
}
