package datastructures

//ReverseWords like "I feel pretty" -> "pretty feel I"
func ReverseWords(s string) string {
	reverse := ""
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			reverse = word + " " + reverse
			word = ""
		} else {
			word += string(s[i])
		}
	}

	return word + " " + reverse
}
