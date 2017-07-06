package permutations

func permutations(input string) []string {

	solutions := []string{}
	var perms func(string, map[int]bool)
	perms = func(permutation string, visited map[int]bool) {

		if len(permutation) > len(input) {
			return
		}

		if len(permutation) == len(input) {
			solutions = append(solutions, permutation)
			return
		}

		for i, r := range input {
			if _, found := visited[i]; !found {
				permutation += string(r)
				visited[i] = true
				perms(permutation, visited)
				delete(visited, i)
				permutation = permutation[:len(permutation)-1]
			}
		}
	}

	perms("", map[int]bool{})

	return solutions
}
