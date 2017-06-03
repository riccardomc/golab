package arraysstrings

func lcgMaker(seed int) func() int {
	next := seed
	m := 9
	a := 4
	c := 1

	return func() int {
		next = (a*next + c) % m
		return next
	}
}
