package productofallints

//ExceptAtIndex returns an array with the products of all items
//except the item at index i
func ExceptAtIndex(ints []int) []int {

	if len(ints) < 2 {
		return nil
	}

	exceptAtIndex := make([]int, len(ints))

	productSoFar := 1
	for i := 0; i < len(ints); i++ {
		exceptAtIndex[i] = productSoFar
		productSoFar *= ints[i]
	}

	productSoFar = 1
	for i := len(ints) - 1; i >= 0; i-- {
		exceptAtIndex[i] *= productSoFar
		productSoFar *= ints[i]
	}

	return exceptAtIndex
}
