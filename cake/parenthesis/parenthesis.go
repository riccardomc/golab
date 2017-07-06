package parenthesis

func parenthesis(message string, open int) int {
	bytes := []byte(message)

	level := 1
	for i := open + 1; i < len(bytes); i++ {
		if bytes[i] == '(' {
			level++
		} else if bytes[i] == ')' {
			level--
		}

		if level <= 0 {
			return i
		}
	}

	return -1
}
