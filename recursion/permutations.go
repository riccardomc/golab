package recursion

//Permutations does what its name implies
func Permutations(input []int) [][]int {

	permutations := make([][]int, 0)
	visited := make([]bool, len(input))
	var p func(int, []int)

	p = func(level int, permutation []int) {
		if level >= len(input) {
			perm := make([]int, len(input))
			copy(perm, permutation)
			permutations = append(permutations, perm)
			return
		}

		for i := 0; i < len(input); i++ {
			if !visited[i] {
				permutation[level] = input[i]
				visited[i] = true
				p(level+1, permutation)
				visited[i] = false
			}
		}
	}

	p(0, make([]int, len(input)))

	return permutations
}
