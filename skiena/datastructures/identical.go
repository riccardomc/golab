package datastructures

// Identical returns true if the t1, t2 are identical binary trees
func Identical(t1, t2 *BTNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}

	if t1.Key != t2.Key {
		return false
	}

	return Identical(t1.Left, t2.Left) && Identical(t1.Right, t2.Right)
}
