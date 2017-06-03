package arraysstrings

// ZeroOrOneDistance returns true if s2 is the result of an insert, remove or
// replace operation applied to s1
func ZeroOrOneDistance(s1, s2 string) bool {

	// insert
	if len(s2) == len(s1)+1 {
		return distance(s1, s2)
	}

	// remove
	if len(s2) == len(s1)-1 {
		return distance(s2, s1)
	}

	if len(s2) == len(s1) {
		difference := false
		for i := 0; i < len(s2); i++ {
			if s1[i] != s2[i] {
				if difference {
					return false
				}
				difference = true
			}
		}
		return true
	}

	return false
}

func distance(s1, s2 string) bool {
	difference := 0
	for i := 0; i < len(s1); i++ {
		if s2[i] != s1[i-difference] {
			if difference == 1 {
				return false
			}
			difference = 1
		}
	}
	return true
}
