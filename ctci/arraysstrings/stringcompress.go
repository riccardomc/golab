package arraysstrings

import "strconv"

// Compress returns s "compressed"
func Compress(s string) string {
	compressed := ""
	count := 1

	if len(s) <= 1 {
		return s
	}

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count = count + 1
		} else {
			compressed += s[i-1:i] + strconv.Itoa(count)
			count = 1
		}
	}

	compressed += s[len(s)-1:len(s)] + strconv.Itoa(count)

	if len(compressed) < len(s) {
		return compressed
	}

	return s
}
