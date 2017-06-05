package recursion

// TripleStepRecursive recursive
func TripleStepRecursive(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1 + TripleStepRecursive(1)
	}

	if n == 3 {
		return 1 + TripleStepRecursive(2) + TripleStepRecursive(1)
	}

	return TripleStepRecursive(n-1) + TripleStepRecursive(n-2) + TripleStepRecursive(n-3)
}

// TripleStepMemo memoization
func TripleStepMemo(n int) int {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 4

	if n <= 3 {
		return m[n]
	}

	for i := 4; i <= n; i++ {
		m[i] = m[i-1] + m[i-2] + m[i-3]
	}

	return m[n]
}

// TripleStepDynamic dynamic
func TripleStepDynamic(n int) int {
	m := []int{1, 1, 2, 4}

	if n <= 3 {
		return m[n]
	}

	for i := 4; i <= n; i++ {
		m[0] = m[1] + m[2] + m[3]
		m[1] = m[2]
		m[2] = m[3]
		m[3] = m[0]
	}

	return m[0]

}
